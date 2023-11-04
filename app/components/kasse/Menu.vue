import { re } from '../../composables/regex';
<script setup lang="ts">
const { storeId } = $defineProps<{
  storeId: string;
}>();

const kasseStore = useKasseStore(storeId);

</script>

<template>
  <div class="grid mr-3">
  <div v-for="item in kasseStore.sortedShop" :key="item.id">
    <KasseShopItem :storeId="storeId" :itemId="item.id"/>
  </div>
</div>
</template>

<style scoped>
.grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, 230px);
  grid-template-rows: repeat(auto-fill, 230px);
  gap: 40px;
}
</style>
