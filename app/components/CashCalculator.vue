<script setup lang="ts">

const probs = defineProps({
  requested: {
    type: Number,
    required: true
  },
  shown: {
    type: Boolean,
    default: true
  }
});

const emit = defineEmits(['paied', 'cancled']);

let requested = $computed(() => `${ probs.requested }`)
let show = $toRef(probs, 'shown');

let sum = $ref("");
let page = $ref(0);
sum = "0";

function add(val: string) {
  sum += val;
}
function reset() {
  sum = "0";
}
</script>

<template>
  <v-dialog fullscreen transition="dialog-bottom-transition" v-model="show">
    <transition name="fade" mode="out-in">

      <v-card class="wrapper" height="100vh" v-if="page == 0" :key="1">
        <div class="display">
          <h2>Zurück</h2>
          <h1 style="font-size: 50px; align-self: flex-end">
            {{ (parseInt(sum) / 100).toFixed(2) }} €
          </h1>
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
        <v-btn @click="page = 1" flat color="green" :disabled="parseInt(sum) / 100 < parseInt(requested)">Checkout
          <v-icon style="margin-left: 5px" icon="mdi-arrow-right-bold-circle-outline">
          </v-icon></v-btn>
      </v-card>


      <v-card v-else height="100vh" class="wrapper" :key="2">
        <div class="display">
          <div class="table-row">
            <div class="name">Kosten</div>
            <div class="value"> {{ parseInt(requested).toFixed(2) }} €</div>


          </div>
          <div class="table-row">
            <div class="name">gegeben</div>
            <div class="value"> {{ (parseInt(sum) / 100).toFixed(2) }} €</div>

          </div>
          <div class="table-row">
            <div class="name">rückgabe </div>
            <div class="value"> {{ (parseInt(requested) - parseInt(sum) / 100).toFixed(2) }} € </div>


          </div>

        </div>
        <v-btn @click="page = 0" flat color="orange">ändern
          <v-icon style="margin-left: 5px" icon="mdi-alpha-x-circle-outline ">
          </v-icon></v-btn>
        <v-btn @click="page = 0; emit('paied')" flat color="green">Speichern
          <v-icon style="margin-left: 5px" icon="mdi-alpha-x-circle-outline ">
          </v-icon></v-btn>
        wohoo
      </v-card>
    </transition>
  </v-dialog>
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
  min-height: 80px;
  width: 40%;
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
  margin: 20px;
  width: 100px;
  height: 100px;
  border-radius: 100px;
  background-color: #fff;
  align-items: center;
  font-size: 30px;
  font-weight: bold;
  justify-content: center;
  padding: 10px;
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
</style>
