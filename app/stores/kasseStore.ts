import makeSeparatedStore from "@/composables/separatedStore"

export const useKasseStore = makeSeparatedStore((id:string) => defineStore(`kasse/${id}`, {
  state: () => ({
    _count: 0,
    
    _shop: beispiel, // as IShopItem[],
    _cart: [] as CartItem[],

    //TODO: lastPriceChangeTime
    //prices should be updated every 10 secs or subscribe to changes
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
    }
  },
}))


type CartItem = {
  id: string
  quantity: number
}


const beispiel:IShopItem[] = [
  {
    id : "1",
    title : "Bier",
    price : 1.50,
    description : "Ein Bier",
    pfand : false,
    pfand_item : "2",
    tags : ['bar'],
  },
  {
    id : "2",
    title : "Becher",
    price : 1.0,
    description : "Ein Becher",
    pfand : true,
    pfand_item : "",
    tags : ['bar'],
  },
  {
    id: "3",
    title: "Cola",
    price: 2.0,
    description: "Cola",
    pfand: false,
    pfand_item: "2",
    tags: ['bar'],
  },
  {
    id: "4",
    title: "Wasser",
    price: 1.0,
    description: "Wasser",
    pfand: false,
    pfand_item: "2",
    tags: ['bar'],
  }
]