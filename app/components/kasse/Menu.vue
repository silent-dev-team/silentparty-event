<script setup lang="ts">
const { storeId } = $defineProps<{
  storeId: string;
}>();

const kasseStore = useKasseStore(storeId);
let sortedShop = $ref(kasseStore.sortedShop)

kasseStore.$subscribe((state) => {
  console.log("kasseStore.$subscribe", state);
  sortedShop = kasseStore.sortedShop
});

</script>

<template>
  <div class="grid mx-3">
  <div v-for="item in sortedShop" :key="item.id">
    <KasseShopItem :storeId="storeId" :itemId="item.id"/>
  </div>
</div>
</template>

<style scoped>
.grid {
  max-height: calc(100vh - 120px);
  overflow-y: auto;
  padding-bottom: 20px;
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(210px, auto));
  grid-template-rows: repeat(auto-fill, minmax(200px, auto));
  gap: 40px;
}
@media (max-width: 1228px) {
  .grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(180px, auto));
    grid-template-rows: repeat(auto-fill, minmax(200px, auto));
    gap: 40px;
  }
}

@media (max-width: 1028px) {
  .grid {
    grid-template-columns: repeat(auto-fill, minmax(150px, auto));
    grid-template-rows: repeat(auto-fill, minmax(150px, auto));
    gap: 20px;
  }
}
@media (max-width: 800px) {
  .grid {
    grid-template-columns: repeat(auto-fill, minmax(120px, auto));
    grid-template-rows: repeat(auto-fill, 120px);
    gap: 20px;
  }
}

</style>
