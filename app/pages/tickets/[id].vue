<script setup lang="ts">
const route = useRoute();
const id = route.params.id as string;

</script>

<template>
  <div class="bg-blue-grey-lighten-5" style="min-height: 100vh; width: 100%; padding: 3rem;">
    <Ticket :id="id" :readonly="true"/>
  </div>
</template>