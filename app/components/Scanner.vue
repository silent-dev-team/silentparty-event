<script setup lang="ts">
import QrScanner from 'qr-scanner';

const props = defineProps({
  reset: {
    type: Boolean,
    default: false
  },
  overlaypath: {
    type: String,
    default: ''
  }
});

const emit = defineEmits<{
  onScan: [value: string]
  results: [value: string[]]
  "update:reset": [value: boolean]
}>()

let cams = $ref<QrScanner.Camera[]>([]);
let qrScanner = $ref<QrScanner>();

let showSettings = $ref(false);

let blink = $ref(false);
let flip = $ref(false);

let results = $ref<string[]>([]);

function setCamera(label: string) {
  qrScanner!.setCamera(cams.find(el => el.label == label)!.id);
  localStorage.setItem('cam', label);
}

onMounted(() => { 
  if (process.client) {
    QrScanner.listCameras(true).then(devices => {
      cams = devices;
      console.log(devices);
    });


    const vid = window.document.getElementById('qr-video');
    qrScanner = new QrScanner(vid as HTMLVideoElement, (result:any) => {
      if (result instanceof Error) return;
      if (results.includes(result)) return;
      results.unshift(result);
      emit('onScan', result);
      emit('results', results);
      blink = true;
      setTimeout(() => blink = false, 300);
    });
    qrScanner.start();
  }
});

function clearHistory() {
  results = [];
}

onUnmounted(() => {
  qrScanner?.stop();
});

watch(() => props.reset, () => {
  if (!props.reset) return;
  clearHistory();
  emit('update:reset', false);
});

localStorage.setItem('cam', JSON.stringify(results));
</script>

<template>
  <div>
    <div class="frame">
      <div style="position: relative;">
        <div :class="`frame ${flip?'flip':''}`">
          <video id="qr-video" :class="`video ${blink?'blink':''}`"></video>
        </div>
        <div style="width: 100%; height: 80%; position: absolute; top:0;left:0;" class="d-flex justify-center align-center">
          <img :src="`${props.overlaypath}`" style="width:60vw; max-width: 50%; aspect-ratio: 1/1; max-height: min(50%,50vh);">
        </div>
      </div>
      <v-btn 
          style="position: absolute; right: 1rem; top: 1rem;z-index: 100;"
          color="transparent-white"
          icon="mdi-cog"
          size="large"
          @click="showSettings = true"
          >
        </v-btn>
    </div>
    <!-- <div class="history">
      <div v-for="item, i in results">
        <v-card :class="`mx-auto my-1 text-center toast ${i==0?'big':'0'}`" :style="`background-color: rgba(255,255,255,${1 - i*0.2}); color: rgba(1,1,1,${1 - i*0.15});`">
          <p>{{ item }}</p>
        </v-card>
      </div>
    </div> -->
    <v-dialog
        v-model="showSettings"
        width="auto"
      >
        <v-card width="800px" max-width="90%">
          <v-card-title>Settings</v-card-title>
          <v-card-text class="pa-3">
            <v-select
              :items="cams.map(cam => cam.label)"
              @update:model-value="setCamera($event)"
              label="Select a camera"
              variant="outlined"
            />
            <v-checkbox label="Flip" v-model="flip"></v-checkbox>
          </v-card-text>
          <v-card-actions class="d-flex justify-content-right">
            <v-btn @click="showSettings = false">OK</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
  </div>
</template>

<style scoped>

.frame{
  overflow: hidden;
  display: flex;
  justify-content: center;
  align-items: center;
  height:100vh;
}
.video {
  height:100vh;
  transition: all 0.15s ease-in-out; 
  filter:brightness(1);

}

.flip {
  -moz-transform: scale(-1, 1);
  -webkit-transform: scale(-1, 1);
  -o-transform: scale(-1, 1);
  -ms-transform: scale(-1, 1);
  transform: scale(-1, 1);
}

.blink {
  filter:brightness(5);
}

.history {
  position: fixed;
  top: 0;
  right: 0;
  width: 100px;
  height: 30vh;
  background-color: transparent;
  overflow: hidden;
}

.mode {
  position: fixed;
  top: 5px;
  left: 0;
  width: 100px;
  height: 30vh;
  background-color: transparent;
  overflow: hidden;
}
.big{
  animation: scaleDown 0.5s ease-in-out;
}

@keyframes scaleDown {
  0% {
    transform: scale(1.4);
  }
  100% {
    transform: scale(1);
  }
}
.card {
  background-color: rgba(255,255,255,1);
  color: rgba(1,1,1,1);
  left: 2.5%;
  width: 95%;
  position: absolute;
  top:95%;
}
</style>