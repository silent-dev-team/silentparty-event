import type { Position } from './../composables/pocketbase';
import makeSeparatedStore from "@/composables/separatedStore"
import { usePocketbase } from "#imports"

export const useShopStore = defineStore('shop', {
  state: () => ({
    _shop: [] as IShopItem[],
    _pb: usePocketbase(),
  }),
  getters: {
    shop: (state) => state._shop,
    getShopItem: (state) => (id: string) => state._shop.find(item => item.id === id),
    getShopTag: (state) => (tag: string) => state._shop.filter(item => item.tags?.includes(tag)),
  },
  actions: {
    async loadShop() {
      console.log('load shop')
      const items = await this._pb.collection('shop_items').getFullList<ShopItemRecord>()
      console.log(items)
      this._shop = items
    },
  },
})

export const useKasseStore = makeSeparatedStore((id:string) => defineStore(`kasse/${id}`, {
  state: () => ({
    _count: 0,
    
    _shop: useShopStore().getShopTag(id),
    _cart: [] as CartItem[],

    _pb: usePocketbase(),
  }),
  getters: {
    count: (state) => state._count,
    
    shop: (state) => state._shop,
    sortedShop: (state) => {
      const shop = [...state._shop]
      shop.sort(
        (a, b) => {
          if (a.pfand && !b.pfand) return -1;
          if (!a.pfand && b.pfand) return 1;
          return a.price - b.price;
        }
      )
      return shop
    },
    cart: (state) => state._cart,

    getShopItem: (state) => (id: string) => state._shop.find(item => item.id === id),
    getCartItem: (state) => (id: string) => state._cart.find(item => item.id === id),

    getCartCount: (state) => (id: string) => state._cart.find(item => item.id === id)?.quantity ?? 0,

    cartSum: (state) => {
      let sum = 0
      state._cart.forEach(item => {
        const shopItem = state._shop.find(shopItem => shopItem.id === item.id)
        if (shopItem) {
          sum += shopItem.price * item.quantity
        }
      })
      return sum
    },
  },
  actions: {

    // Cart
    addOneCardItem(id: string) {
      this.addCartItem(id, 1)
    },
    addCartItem(id: string, quantity: number) {
      const item = this.getCartItem(id);

      if (item) {
        item.quantity += quantity
      } else {
        this._cart.push({ id, quantity })
      }

      if (item?.quantity == 0) {
        this._cart = this._cart.filter(item => item.id !== id)
        return
      }

      const shopItem = this.getShopItem(id)
      if (!shopItem) return;
      if (!shopItem.pfand && shopItem.pfand_item) {
        this.addCartItem(shopItem.pfand_item, quantity)
      }
    },

    removeOneCartItem(id: string) {
      this.removeCartItem(id, 1)
    },
    removeCartItem(id: string, quantity: number) {
      const item = this.getCartItem(id);
      const shopItem = this.getShopItem(id)

      if (!item && shopItem?.pfand) {
        this._cart.push({ id: id, quantity: -1})
        return
      } else {
        if (!item) return;
      }

      item!.quantity -= quantity

      if (item!.quantity == 0) {
        this._cart = this._cart.filter(item => item.id !== id)
      }

      if (!shopItem || shopItem!.pfand) return;
      if (!shopItem.pfand && shopItem.pfand_item) {
        this.removeCartItem(shopItem.pfand_item, quantity)
      }
    },

    async checkout() {
      const positions: Position[] = this._cart.map(item => {
        return {
          itemId: item.id,
          amount: item.quantity
        }
      })

      await this._pb.checkout(positions).catch(err => {
        console.error(err)
        throw err
      })
    },

    async loadShop() {
      console.log('load item from shop to kasse')
      const shopStore = useShopStore()
      this._shop = shopStore.getShopTag(id)
    },
  },
}))


type CartItem = {
  id: string
  quantity: number
}
