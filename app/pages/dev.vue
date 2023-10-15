<script setup lang="ts">

  let val = $ref(0);
  let _shown = $ref(false);
</script>

<template>
  <cash-calculator :requested="val" :shown="_shown" @paied="_shown=false; val=0" @cancled="_shown=false"> </cash-calculator>
  <input type="number" v-model="val">
  <button @click="_shown=!_shown">+</button>
</template>
