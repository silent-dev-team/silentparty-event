<script setup lang="ts">
const route = useRoute();
const id = route.params.id as string;

let showPinDialog = $ref(true);

function handlePin(pin: string) {
  console.log(pin);
  showPinDialog = false;
}
</script>

<template>
  <v-dialog v-model="showPinDialog">
    <v-card class="mx-auto pa-10">
      <PinField @update="handlePin($event)"/>
    </v-card>
  </v-dialog>
  <div class="bg-blue-grey-lighten-5" style="min-height: 100vh; width: 100%; padding: 3rem;">
    <Ticket :id="id"/>
  </div>
</template>