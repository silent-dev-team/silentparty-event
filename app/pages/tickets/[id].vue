<script setup lang="ts">
const route = useRoute();
const id = route.params.id as string;
const pb = usePocketbase();
const router = useRouter();

function update() {
  setTimeout(() => {
    router.push('/tickets/filled');
  }, 1000);
}
</script>

<template>
  <div class="mx-auto" style="max-width: 800px;">
    <Ticket :id="id" @update="update"/>
  </div>
</template>
