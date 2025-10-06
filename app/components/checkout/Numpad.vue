<script setup lang="ts">

const props = defineProps({
  requested: {
    type: Number,
    required: true
  },
  preset: {
    type: Number,
    default: 0
  },
});

const emit = defineEmits(['paied', 'cancled']);

let requested = $computed(() => `${ props.requested }`)

let sum = $ref( props.preset ? (props.preset * 100).toFixed(0) : "0" );

function add(val: string) {
  sum += val+"";
}
function reset() {
  sum = "0";
}
</script>

<template>
  <v-card class="wrapper" height="100vh"  :key="1">
    <div class="display">
      <div class="w-100 d-flex justify-space-between text-green-darken-3 overflow-hidden text-no-wrap">
        <h2>
          Preis
        </h2>
        <h2 >
          {{ (parseFloat(requested)).toFixed(2) }} €
        </h2>
      </div>
      <div class="w-100 d-flex justify-space-between overflow-hidden text-no-wrap">
        <h1>
          Gegeben
        </h1>
        <h1 style=" align-self: flex-end">
          {{ (parseInt(sum) / 100).toFixed(2) }} €
        </h1>
      </div>
      <div class="w-100 d-flex justify-space-between overflow-hidden text-red-darken-3 text-no-wrap">
        <h2>
          Rückgeld
        </h2>
        <h2 style=" align-self: flex-end">
          {{ (( Math.min(parseFloat(requested) - parseInt(sum)/ 100,0)) ).toFixed(2) }} €
        </h2>
      </div>
    </div>
    <div class="keypad_wrapper">
      <div class="row" v-for="row in [
        [1, 2, 3],
        [4, 5, 6],
        [7, 8, 9],
      ]" :key="row[0]">
        <div class="key" v-for="item in row" :key="item" @click="add(item + '')">
          {{ item }}
        </div>
      </div>

      <div class="row">
        <div class="key" style="color: red" @click="reset">R</div>
        <div class="key" @click="add('0')">0</div>
        <div class="key" @click="add('00')">00</div>
      </div>
    </div>
    <div class="d-flex justify-space-between w-100 mt-3 px-2" style="max-width: 340px;">
    <v-btn @click="reset();emit('cancled')" flat color="red" variant="outlined">cancle
      <v-icon style="margin-left: 5px" icon="mdi-close-circle-outline">
      </v-icon>
    </v-btn>
    <v-btn @click="reset();emit('paied')" flat color="green" :disabled="(parseInt(sum) / 100) < parseFloat(requested)">Checkout
      <v-icon style="margin-left: 5px" icon="mdi-arrow-right-bold-circle-outline">
      </v-icon>
    </v-btn>
  </div>
  </v-card>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity .05s ease-in, transform .05s ease-in;
}

.fade-enter-from {
  opacity: 0;
  transform: translateX(-3%);
}

.fade-leave-to {
  opacity: 0;
  transform: translateX(3%);

}

.display {
  min-height: 120px;
  width: 40%;
  min-width: 300px;
  background-color: #ffffff;
  margin: 16px;
  border-radius: 5px;
  display: flex;
  flex-direction: column;
}

.wrapper {
  display: flex;
  align-items: center;
  background-color: #f7f7f7;
}

.row {
  display: flex;
  flex-direction: row;
  justify-content: center;
}

.key {
  margin: 15px;
  width: 90px;
  height: 90px;
  border-radius: 100px;
  background-color: #fff;
  align-items: center;
  font-size: 30px;
  line-height: 10px;
  font-weight: bold;
  justify-content: center;
  display: flex;
  box-shadow: rgb(255, 254, 254) -10px -10px 40px -12px,
    rgba(0, 0, 0, 0.048) 5px 5px 20px 5px;
  transition: background 0.6s ease-out, color 0.6s ease-out;
}



.key:hover {
  cursor: pointer;
  transform: scale(0.99);
  box-shadow: rgb(255, 254, 254) -10px -10px 40px -12px,
    rgba(0, 0, 0, 0.048) 2px 2px 10px 2px;
}

.key:active {
  transition: background 0s, color 0s;

  cursor: pointer;
  transform: scale(0.99);
  color: #fff;
  background-color: #000000;
  box-shadow: rgb(255, 254, 254) -10px -10px 40px -12px,
    rgba(0, 0, 0, 0.048) 2px 2px 10px 2px;
}

.keypad_wrapper {}

h1{
  font-size: 35px;
}

@media only screen and (max-width: 470px) {
  h1{
    font-size: min(9vw, 35px);
  }

  .key{
    margin: 2%;
    flex-grow:1;
    width: 25vw;
    height: auto;
    padding: 0;
    aspect-ratio: 1 / 1;
    font-size: 20px;

  }
}
@media only screen and (max-height: 722px) {

  .key{
    margin: 2%;
    flex-grow:1;
    width: 15vh;
    height: auto;
    padding: 0;
    aspect-ratio: 1 / 1;
    font-size: 20px;

  }
}
</style>
