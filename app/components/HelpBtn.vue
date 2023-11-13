<script setup lang="ts">
import type { UnsubscribeFunc } from 'pocketbase';

const pb = usePocketbase()

let { from, msg, icon, sync } = $defineProps<{
  from: string
  msg: string
  icon: string
  sync: boolean
}>()

let alert = $ref<AlertRecord>()
let color = $computed(() => {
  if (alert?.seen) {
    return 'green'
  }
  if (alert?.active) {
    return 'red'
  }
  return undefined

})
let pulse = $computed(() => {
  if (alert?.seen) {
    return 'pulse-reverse'
  }
  if (alert?.active) {
    return 'pulse'
  }
  return undefined
})


if (sync) {
  try {
    alert = await pb.collection('alerts').getFirstListItem<AlertRecord>(`from = "${from}" && msg = "${msg}"`)
  } catch (e) {
    console.error(e)
    alert = await pb.collection('alerts').create<AlertRecord>({ msg: msg, from: from, active: false})
  }
}

const unsubscripe = await pb.collection('alerts').subscribe<AlertRecord>('*', function (e) {
  if (e.record.id === alert?.id) {
    if (!sync) {
      alert = e.record.active ? e.record as AlertRecord : undefined
    } else {
      alert = e.record as AlertRecord
    }
  }
});

async function onClick() {
  if (alert) {
    pb.collection('alerts').update<AlertRecord>(alert.id, { active: !alert.active })
  } else {
    alert = await pb.collection('alerts').create<AlertRecord>({ msg: msg, from: from, active: true})
  }
}


onUnmounted(() => {
    unsubscripe();
});

</script>

<template>
  <v-btn :class="pulse" :color="color" :icon="icon" @click="onClick"/>
</template>

<style scoped>
.pulse {
  display: block;
  cursor: pointer;
  animation: pulse 1s infinite;
}

.pulse:hover {
  animation: none;
}

.pulse-reverse {
  display: block;
  cursor: pointer;
  animation: pulse 2s reverse infinite;
}



@-webkit-keyframes pulse {
  0% {
    -webkit-box-shadow: 0 0 0 0 rgba(204, 44, 44, 0.4);
  }
  70% {
    -webkit-box-shadow: 0 0 0 20px rgba(204, 44, 44, 0);
  }
  100% {
    -webkit-box-shadow: 0 0 0 0 rgba(204, 44, 44, 0);
  }
}
@keyframes pulse {
  0% {
    -moz-box-shadow: 0 0 0 0 rgba(204, 44, 44, 0.4);
    box-shadow: 0 0 0 0 rgba(204, 44, 44, 0.4);
  }
  70% {
    -moz-box-shadow: 0 0 0 20px rgba(204, 44, 44, 0);
    box-shadow: 0 0 0 20px rgba(204, 44, 44, 0);
  }
  100% {
    -moz-box-shadow: 0 0 0 0 rgba(204, 44, 44, 0);
    box-shadow: 0 0 0 0 rgba(204, 44, 44, 0);
  }
}
</style>