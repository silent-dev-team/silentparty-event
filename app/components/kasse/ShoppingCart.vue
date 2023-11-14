<script setup lang="ts">
const { storeId } = $defineProps<{
  storeId: string;
}>();

const kasseStore = useKasseStore(storeId);
const notifyer = useNotifyer()
const fromShop = kasseStore.getShopItem

let showCheckout = $ref(false)

function checkout() {
  try {
    kasseStore.checkout()
    notifyer.notify("Erfolgreich bezahlt", "success")
    kasseStore.$reset()
  } catch (e) {
    console.error(e)
    notifyer.notify("Fehler beim Bezahlen", "error")
  } finally {
    showCheckout = false
  }
}

</script>

<template>
  <v-card class="d-flex flex-column justify-start pt-2" >
    <KasseCartItem v-for="item in kasseStore.cart" :key="item.id" :title="fromShop(item.id)!.title" :amount="item.quantity" :price="fromShop(item.id)!.price * item.quantity" />
    <div style="flex-grow: 1;"></div>
    <v-divider ></v-divider>
    <div class="d-flex justify-space-between pa-3 my-3" style="flex-shrink: 1;">
      <h2>Summe:</h2>
      <h2>{{ parseFloat(""+kasseStore.cartSum).toFixed(2) }} €</h2>
      
    </div>
    <v-btn :disabled="kasseStore.cartSum == 0" color="red" variant="outlined" class="ma-3" @click="kasseStore.$reset()"> Abbrechen </v-btn>
    <v-btn :disabled="kasseStore.cartSum == 0" color="green" class="ma-3" @click="showCheckout = true"> Bezahlen </v-btn>
    <Checkout :requested="kasseStore.cartSum" :shown="showCheckout" @update:shown="showCheckout = $event" @paied="checkout" @cancled="showCheckout = false" />
  </v-card>
</template>

<style scoped>
  h2{
    font-size: 1.5rem;
  }

  @media (max-width: 800px) {
    h2{
      font-size: 1.2rem;
    }
  }
</style>
```