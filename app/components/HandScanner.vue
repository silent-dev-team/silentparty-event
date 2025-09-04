<script setup lang="ts">
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

function focusInput() {
  const input = document.getElementById('hand-scan-input') as HTMLInputElement;
  input.value = '';
  input.focus();
}
</script>

<template>
  <div class="w-full">
    <div class="mx-auto w-75 pa-2" style="max-width: 800px;">
      <v-text-field
        id="hand-scan-input"
        label="Scan"
        dense
        outlined
        class="ma-2"
      />
    </div>
  </div>
</template>
