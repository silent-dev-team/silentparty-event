<script setup lang="ts">
const route = useRoute();
const id = route.params.id as string;
const pb = usePocketbase();

definePageMeta({
  layout: 'admin',
});

</script>

<template>
  <Ticket :id="id" :readonly="!pb.authStore.isAdmin"/>
</template>