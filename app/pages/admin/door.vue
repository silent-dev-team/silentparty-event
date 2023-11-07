<script setup lang="ts">
definePageMeta({
  layout: 'admin',
});

const pb = usePocketbase()

let alert = $ref<AlertRecord>()
let color = $computed(() => {
  if (alert?.active) {
    return 'red'
  }
  return 'green'

})

const unsubscripe = await pb.collection('alerts').subscribe<AlertRecord>('*', function (e) {
  if (e.record.msg === 'hodor' ) {
    console.log(e.record);
    alert = e.record.active ? e.record as AlertRecord : undefined
  }
});

onUnmounted(() => {
    unsubscripe();
});

</script>

<template>
  <v-sheet height="100%" width="100%" :color="color">
  </v-sheet>
</template>