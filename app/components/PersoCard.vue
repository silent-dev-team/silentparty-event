<script setup lang="ts">
import { computed, ref } from 'vue'

const props = defineProps<{
  surname: string
  givenNames: string
  dob: string // TT.MM.JJJJ
  photoUrl?: string
  initialZoom?: boolean
  uppercase?: boolean
}>()

const isZoomed = ref(!!props.initialZoom)

const surnameText = computed(() => props.uppercase === false ? (props.surname ?? '') : (props.surname ?? '').toUpperCase())
const givenText   = computed(() => props.uppercase === false ? (props.givenNames ?? '') : (props.givenNames ?? '').toUpperCase())
const dobText     = computed(() => props.dob ?? '')

function toggleZoom(){
  isZoomed.value = !isZoomed.value
}
</script>

<template>
  <div class="inline-flex flex-col items-center gap-3">
    <!-- Card -->
    <div
      class="relative rounded-[18px] shadow-[0_0_0_1px_rgba(0,0,0,.08),0_10px_25px_rgba(0,0,0,.15)] overflow-hidden id-texture holo transition-transform duration-200"
      :class="[isZoomed ? 'scale-[1.15]' : 'scale-100']"
      style="width:520px; aspect-ratio: 85.6 / 53.98;"
      aria-label="Beispiel eines Personalausweis-Frontlayouts"
    >
      <!-- Foto -->
      <div class="absolute left-3 top-10 w-[34%] aspect-[3/4] rounded-md bg-neutral-300/50 border border-black/10 overflow-hidden flex items-center justify-center">
        <img v-if="photoUrl" :src="photoUrl" alt="Ausweisfoto" class="w-full h-full object-cover" />
        <span v-else class="text-[10px] text-neutral-700">Foto</span>
      </div>

      <!-- Felder (nur Vorderseite / Name, Vornamen, DOB) -->
      <div class="absolute right-3 left-[40%] top-10 text-[10px] text-neutral-900/90 space-y-[7px]">
        <div class="grid grid-cols-[90px_1fr] items-baseline gap-2">
          <div class="text-neutral-600">Name / Surname</div>
          <div class="font-semibold truncate" :title="surnameText">{{ surnameText || '—' }}</div>
        </div>
        <div class="grid grid-cols-[90px_1fr] items-baseline gap-2">
          <div class="text-neutral-600">Vornamen / Given Names</div>
          <div class="font-semibold truncate" :title="givenText">{{ givenText || '—' }}</div>
        </div>
        <div class="grid grid-cols-[90px_1fr] items-baseline gap-2">
          <div class="text-neutral-600">Geburtsdatum / Date of Birth</div>
          <div class="font-semibold">{{ dobText || '—' }}</div>
        </div>
      </div>

      <!-- Hinweis unten -->
      <div class="absolute inset-x-0 bottom-0 text-[8.5px] text-neutral-600/90 px-3 py-1 flex items-center justify-between">
        <span>Beispiel-UI · kein amtliches Dokument</span>
        <span class="font-mono">85.6×53.98 mm</span>
      </div>
    </div>

    <!-- Zoom-Button (gewünscht) -->
    <button
      type="button"
      class="px-3 py-1 rounded-lg bg-neutral-800 text-neutral-100 hover:bg-neutral-700"
      @click="toggleZoom"
    >Zoom</button>
  </div>
</template>

<style scoped>
.id-texture {
  background: repeating-linear-gradient(135deg, rgba(255,255,255,.35) 0 6px, rgba(0,0,0,.03) 6px 12px),
              radial-gradient(1000px 600px at 10% 20%, #e7f5ea, transparent),
              radial-gradient(800px 400px at 100% 0%, #f7f6e9, transparent),
              linear-gradient(160deg, #f3f7f2 0%, #ecf5ef 45%, #f6f5ea 100%);
}
.holo::after {
  content: "";
  position: absolute; inset: 0;
  background: conic-gradient(from 180deg at 70% 30%, rgba(255,255,255,.0), rgba(255,255,255,.25), rgba(255,255,255,.0));
  mix-blend-mode: overlay; pointer-events: none;
}
</style>
