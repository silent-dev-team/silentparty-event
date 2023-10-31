<script setup lang="ts">
const { storeId } = $defineProps<{
  storeId: string;
}>();

const kasseStore = useKasseStore(storeId);

</script>

<template>
  <div class="mr-3" style="display: grid; grid-template-columns: repeat(auto-fill, 230px);gap: 40px;">
  <div v-for="item in kasseStore.sortedShop" :key="item.id">
    <KasseShopItem :storeId="storeId" :itemId="item.id"/>
  </div>
</div>
</template>