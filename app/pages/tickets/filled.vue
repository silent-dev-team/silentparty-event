<script setup lang="ts">
const pb = usePocketbase();
</script>

<template>
  <v-card class="card pa-5 text-center">
    <h4>Dein Ticket wurde erfolgreich personalisiert.</h4>
  </v-card>
</template>

<style scoped>
.card {
  max-width: 500px;
  margin: 0 auto;
  animation: popup 0.5s ease-in-out;
}

@keyframes popup {
  0% {
    transform: scale(0.5);
    opacity: 0;
  }
  100% {
    transform: scale(1);
    opacity: 1;
  }
}
</style>
