<script setup lang="ts">

const { storeId } = $defineProps<{
  storeId: string;
}>();

const kasseStore = useKasseStore(storeId);
const shopStore = useShopStore();
const notifyer = useNotifyer();

shopStore.$subscribe((state) => {
  console.log("shopStore.$subscribe", state);
  notifyer.notify("Shop aktualisiert", 'info')
  kasseStore.loadShop();
});

</script>

<template>
  <div class="d-flex justify-space-between ma-2" style="height: 100%;">
    <KasseMenu class="menu" :storeId="storeId" style="flex: 3 0;"/>
    <KasseShoppingCart class="cart" :storeId="storeId" style="flex: 1 0; align-self: stretch;"/>
  </div>
</template>

<style scoped>

.menu{
  flex: 3 0;
}
.cart{
    min-width: 220px;
  }
@media (max-width: 800px) {
  .cart{
    min-width: 190px;
  }
  .menu{
    flex: 4 0;
  }
}

</style>