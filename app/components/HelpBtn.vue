<script setup lang="ts">
import type { UnsubscribeFunc } from 'pocketbase';

const pb = usePocketbase()

let { from, msg } = $defineProps<{
  from: string
  msg: string
  icon: string
}>()

let alert = $ref<AlertRecord>()
let color = $computed(() => {
  console.log("color");
  if (alert?.seen) {
    return 'blue'
  }
  if (alert?.active) {
    return 'red'
  }
  return undefined

})

const unsubscripe = await pb.collection('alerts').subscribe<AlertRecord>('*', function (e) {
  if (e.record.id === alert?.id) {
    console.log(e.record);
    alert = e.record.active ? e.record as AlertRecord : undefined
  }
});

async function onClick() {
  if (alert) {
    pb.collection('alerts').update<AlertRecord>(alert.id, { active: false })
  } else {
    alert = await pb.collection('alerts').create<AlertRecord>({ msg: msg, from: from, active: true})
  }
}


onUnmounted(() => {
    unsubscripe();
});

</script>

<template>
  <v-btn :color="color" :icon="icon" @click="onClick"/>
</template>