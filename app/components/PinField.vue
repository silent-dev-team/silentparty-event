<script setup lang="ts">

const props = defineProps({
  reset: Boolean
});
const emit = defineEmits(['update']);


let pins = $ref(['', '', '', '']);
let inputRefs = $ref<any[]>([]);

function handleInput(index:number) {
  const value = pins[index];
  if (value && /^[0-9]{1}$/.test(value)) {
    const nextIndex = index + 1;
    if (inputRefs[nextIndex]) {
      inputRefs[nextIndex]?.focus();
    }
  } else {
    pins[index] = "";
  }
  if (pins.every(pin => pin.length === 1)) {
    emit('update', pins.join(''));
  }
}

function resetPins() {
  pins = ['', '', '', ''];
  inputRefs[0]?.focus();
}

onMounted(() => {
  inputRefs[0]?.focus();
});

watch(() => props.reset, () => {
  resetPins();
});

</script>

<template>
  <div class="pin-inputs">
    <v-text-field
      width="30px"
      v-for="(pin, index) in pins"
      :key="index"
      v-model="pins[index]"
      @input="() => handleInput(index)"
      type="number"
      pattern="\d"
      :ref="el => (inputRefs[index] = el)"
      variant="outlined"
    ></v-text-field>
  </div>
</template>

<style scoped>
.pin-inputs {
  display: flex;
  gap: 10px;
}

input {
  width: 40px;
  height: 40px;
  text-align: center;
  font-size: 20px;
  border: 1px solid #cccccc;
  border-radius: 4px;
}
</style>