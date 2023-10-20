<script setup lang="ts">
const store1 = useKasseStore1();
const store2 = useKasseStore2();

let { count } = $(storeToRefs(store2));

function increment() {
  count++;
}
</script>

<template>
  <div>
    <h1>Pinia Example</h1>
    <h2 class="mt-3">Store 1 getter/actions</h2>
    {{ store1.count }}
    <v-btn @click="store1.increment()">Increment</v-btn>

    <h2 class="mt-3">Store 1 direct</h2>
    {{ store1._count }}
    <v-btn @click="store1._count++">Increment</v-btn>
    
    <h2 class="mt-3">Store 2</h2>
    {{ store2.count }}
    {{ count }}
    <v-btn @click="store2.increment()">Increment intern</v-btn>
    <v-btn @click="store2.count++">Increment local</v-btn>
    <v-btn @click="increment()">Increment external func</v-btn>
  </div>
</template>