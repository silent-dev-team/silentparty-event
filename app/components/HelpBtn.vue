<script setup lang="ts">

const pb = usePocketbase()

const { from, msg, icon, channel, global } = defineProps<{
  from: string
  msg: string
  channel: string
  icon: string
  global?: boolean
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


onMounted(async () => {
  console.log('mounted', {from, msg, channel, global})
  if (global) {
    try {
      alert = await pb.collection('alerts').getFirstListItem<AlertRecord>(`msg = "${msg}" && channel = "${channel}" && active = true` + (global ? '' : ` && from = "${from}"`), {sort: '-created'})
    } catch (e) {
      alert = undefined
    }
  } 
});

const unsubscripe = await pb.collection('alerts').subscribe<AlertRecord>('*', function (e) {
  if (e.record.msg === msg && e.record.channel === channel && (global || e.record.from === from)) {
    alert = e.record.active ? e.record as AlertRecord : undefined
  }
});

async function onClick() {
  if (alert) {
    pb.collection('alerts').update<AlertRecord>(alert.id, { active: !alert.active })
  } else {
    alert = await pb.collection('alerts').create<AlertRecord>({ msg: msg, from: from, active: true, channel: channel})
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
