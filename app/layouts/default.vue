<script setup lang="ts">
const pb = usePocketbase();
const rt = useRuntimeConfig();

function goToApp(){
  window.location.href = rt.public.pocketbase;
}

function logout() {
  pb.authStore.clear();
  window.location.reload();
}
</script>


<template>
  <div v-if="pb.authStore.isAdmin" style="position: absolute; left: 1rem; top: 1rem;">
    <v-btn @click="goToApp()">zur App</v-btn>
  </div>
  <div v-if="pb.authStore.isAdmin" style="position: absolute; right: 1rem; top: 1rem;">
    <v-btn @click="logout()">Logout</v-btn>
  </div>
  <div class="bg-blue-grey-lighten-5 pa-3 pa-sm-8" style="min-height: 100vh; width: 100%;">
    <div v-if="pb.authStore.isAdmin" class="mt-15"></div>
    <slot></slot>
  </div>
</template>

