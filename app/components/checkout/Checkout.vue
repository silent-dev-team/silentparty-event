<script setup lang="ts">

const props = defineProps({
  requested: {
    type: Number,
    required: true
  },
  shown: {
    type: Boolean,
    default: true
  }
});
const emit = defineEmits(['paied', 'cancled', 'update:shown']);
let show = $toRef(props, 'shown');

</script>

<template>
  <v-dialog :persistent="true" fullscreen transition="dialog-bottom-transition" v-model="show">
    <transition name="fade">
      <CheckoutNumpad :requested="requested" @paied="emit('paied')" @cancled="emit('cancled')"/>
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