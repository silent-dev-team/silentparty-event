<script setup lang="ts">

const { storeId } = $defineProps<{
  storeId: string;
}>();

const kasseStore = useKasseStore(storeId);

</script>

<template>
  <div class="d-flex justify-space-between ma-2" style="height: 100%;">
    <KasseMenu :storeId="storeId" style="flex: 3 0;"/>
    <KasseShoppingCart :storeId="storeId" style="flex: 1 0; align-self: stretch"/>
  </div>
</template>