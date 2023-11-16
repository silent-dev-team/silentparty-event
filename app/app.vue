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

.subappmarker {
  position: absolute;
  left: 1rem;
  top: 1rem;
  background: rgb(1, 18, 39);
  background: linear-gradient(0deg, rgb(0, 0, 0) 0%, #010b1886 70%, #0a4fa4a1 100%);
  backdrop-filter: blur(6px);
  color: white;
  text-align: center;
  border-radius: 2rem;
  width: 6rem;
  height: 2rem;
  z-index: 5000;
  animation: open 800ms ease;
}

@keyframes open {
  0% {
    color: transparent;
    width: 2rem;
    transform: translateX(-180%);
  }
  30% {
    color: transparent;
  }
  100% {
    color: white;
    width: 6rem;
    transform: translateX(0);
  }
}
</style>