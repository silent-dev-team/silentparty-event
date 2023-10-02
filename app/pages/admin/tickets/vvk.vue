<script setup lang="ts">
definePageMeta({
  layout: 'admin',
});

const pb = usePocketbase()

let dialog = $ref(false)
let id = $ref('')

function onScan(url: string) {
  console.log(url)
  const ticketId = url.split('/').pop()
  if (!ticketId) return;
  id = ticketId
  dialog = true
}

async function sell() {
  const record = await pb.collection('tickets').update<TicketRecord>(id, {sold: true});
  if (!record.sold) {
    alert('Verkauf fehlgeschlagen')
  }
  dialog = false
}

</script>

<template>
  <Scanner class="full-screen" @on-scan="onScan($event)"/>
  <v-dialog v-model="dialog">
    <v-card class="pa-3 mx-auto" width="300px">
      <v-card-title class="mx-auto">{{ id }}</v-card-title>
      <v-btn class="mx-auto" width="128px" color="primary" @click="sell()">Verkaufen</v-btn>
    </v-card>
  </v-dialog>
</template>
