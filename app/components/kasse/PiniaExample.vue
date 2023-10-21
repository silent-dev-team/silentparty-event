<script setup lang="ts">
const store1 = useKasseStore1();
const store2 = useKasseStore2();

let { count } = $(storeToRefs(store2));
// das $() macht aus ref() ein $ref() für die Konsistenz

function increment() {
  count++;
}

// ich finde store1 eigentlich ganz cool, weil er etwas übersichtlicher ist
// und weil er build in func wie $reset(), $patch() und $subscribe() hat
// würde den state dann privat lassen und nur über getter und actions zugreifen
</script>

<template>
  <div>
    <h1>Pinia Example</h1>
    <v-btn @click="store1.$reset()">Reset store1</v-btn>
    <h2 class="mt-3">Store 1 getter/actions</h2>
    <h3>store1.count: {{ store1.count }}</h3>
    <v-btn @click="store1.increment()">Increment</v-btn>

    <h2 class="mt-3">Store 1 direct</h2>
    <h3>store1._count: {{ store1._count }}</h3>
    <v-btn @click="store1._count++">Increment</v-btn>

    <h2 class="mt-3">Store 2</h2>
    <v-btn @click="store2.$reset()">Reset store2</v-btn>
    <h3>store2.count: {{ store2.count }}</h3>
    <h3>count: {{ count }}</h3>
    <div class="d-flex justify-space-around">
      <v-btn @click="store2.increment()">Increment intern</v-btn>
      <v-btn @click="store2.count++">Increment local</v-btn>
      <v-btn @click="increment()">Increment external func</v-btn>
    </div>
  </div>
</template>