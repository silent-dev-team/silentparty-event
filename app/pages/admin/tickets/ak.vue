<script setup lang="ts">
definePageMeta({
  layout: 'admin',
});

function onScan(s:string) {
  console.log(s)
  alert(s)
}

</script>

<template>
<Scanner class="full-screen" @on-scan="onScan($event)"/>
</template>