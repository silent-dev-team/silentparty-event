import { KassePiniaEnxample } from '../../.nuxt/components';
<script setup lang="ts">
definePageMeta({
  layout: 'admin',
});

</script>

<template>
  <div style="height: calc(100% - 70px);">
    <Kasse storeId="bar" />
  </div>
</template>
