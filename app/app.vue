<script setup lang="ts">
const notifyer = useNotifyer();
</script>

<template>
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>
  <v-snackbar v-model="notifyer.show" :color="notifyer.color" :timeout="notifyer.duration" location="top">
    <span class="w-100 mx-auto">{{ notifyer.message }}</span>
  </v-snackbar>
</template>

<style>
.full-screen {
  position: fixed;
  left: 0;
  top: 50px;
  height: 100vh;
  width: 100%;
  max-width: 100%;
}
</style>