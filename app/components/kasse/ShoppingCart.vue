<script setup lang="ts">
const { storeId } = $defineProps<{
  storeId: string;
}>();

const kasseStore = useKasseStore(storeId);
const fromShop = kasseStore.getShopItem

let checkout = $ref(false)

</script>

<template>
  <v-card class="d-flex flex-column justify-start" >
    <KasseCartItem v-for="item in kasseStore.cart" :key="item.id" :title="fromShop(item.id)!.title" :amount="item.quantity" :price="fromShop(item.id)!.price * item.quantity" />
    <div style="flex-grow: 1;"></div>
    <v-divider ></v-divider>
    <div class="d-flex justify-space-between pa-3 my-3" style="flex-shrink: 1;">
      <h2>Summe:</h2>
      <h2>{{ parseFloat(""+kasseStore.cartSum).toFixed(2) }} €</h2>
      
    </div>
    <v-btn :disabled="kasseStore.cartSum == 0" color="red" variant="outlined" class="ma-3" @click="kasseStore.$reset()"> Abbrechen </v-btn>
    <v-btn :disabled="kasseStore.cartSum == 0" color="green" class="ma-3" @click="checkout = true"> Bezahlen </v-btn>
    <Checkout :requested="kasseStore.cartSum" :shown="checkout" @update:shown="checkout = $event" @paied="kasseStore.$reset(); checkout = false" @cancled="checkout = false" />
  </v-card>
</template>