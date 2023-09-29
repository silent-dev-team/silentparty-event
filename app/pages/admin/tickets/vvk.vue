import { Scanner } from '../../../.nuxt/components';
<script setup lang="ts">
definePageMeta({
  layout: 'admin',

});

function onScan(url: string) {
  console.log(url)
  const code = url.split('/').pop()
  console.log(code)
}

</script>

<template>
  <Scanner class="full-screen" @on-scan="onScan($event)"/>
</template>
