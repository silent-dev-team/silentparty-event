<script lang="ts" setup>
definePageMeta({
  layout: 'admin',
});

const pb = usePocketbase();
const notifyer = useNotifyer();

let label = $ref('');
let startScanner = $ref(false);
let scannerReset = $ref(false);

async function onScan(s:string) {
  const id = s.split('/').pop();
  if (!id) return;
  const resp = await pb.collection('tickets').update<TicketRecord>(id, {_label: $$(label).value});
  notifyer.notify(`${id} ${label}`, 'success');
}

</script>

<template>
  <Scanner v-if="startScanner" class="full-screen" :overlaypath="Overlay.Ticket" @onScan="onScan($event)" :reset="scannerReset" @update:reset="scannerReset = $event"/>
  <div v-if="startScanner" class="overlay">
    <p>{{ label }}</p>
  </div>
  <v-card v-if="!startScanner" class="mx-auto my-5" maxWidth="400px">
    <v-card-title>
      <v-form @submit.prevent="startScanner = true">
        <v-text-field v-model="label" label="Label" outlined></v-text-field>
        <v-btn type="submit">Scan</v-btn>
      </v-form>
    </v-card-title>
  </v-card>
</template>

<style scoped>

.overlay {
  position: absolute;
  top: 1rem;
  left: 1rem;
  width: auto;
  background-color: #ffffff;
  padding: 6px 12px;
  border-radius: 6px;
  z-index: 100;
  box-shadow: 0px 3px 1px -2px var(--v-shadow-key-umbra-opacity, rgba(0, 0, 0, 0.2)), 0px 2px 2px 0px var(--v-shadow-key-penumbra-opacity, rgba(0, 0, 0, 0.14)), 0px 1px 5px 0px var(--v-shadow-key-penumbra-opacity, rgba(0, 0, 0, 0.12))
}

</style>