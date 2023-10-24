<script lang="ts" setup>
definePageMeta({
  layout: 'admin',
});

const pb = usePocketbase();
const notifyer = useNotifyer();

let seller = $ref('');
let startScanner = $ref(false);
let scannerReset = $ref(false);

async function onScan(s:string) {
  const id = s.split('/').pop();
  if (!id) return;
  const resp = await pb.collection('tickets').update<TicketRecord>(id, {_seller: $$(seller).value});
}

</script>

<template>
  <Scanner v-if="startScanner" class="full-screen" :overlaypath="Overlay.Ticket" @onScan="onScan($event)" :reset="scannerReset" @update:reset="scannerReset = $event"/>
  <v-card v-else>
    <v-card-title>
      <v-text-field v-model="seller" label="Verkäufer" outlined></v-text-field>
      <v-btn @click="startScanner = true">Scan</v-btn>
    </v-card-title>
  </v-card>
</template>