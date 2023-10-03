<script setup lang="ts">
definePageMeta({
  layout: 'admin',
});

//TODO: load pocketbase backend route

const pb = usePocketbase()

const vvkItem = await pb.collection('shop_items').getFirstListItem('title = "VVK Ticket"')

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
  const oldRecord = await pb.collection('tickets').getOne<TicketRecord>(id);
  if (oldRecord.sold) {
    alert('Ticket bereits verkauft')
    return
  }
  const newRecord = await pb.collection('tickets').update<TicketRecord>(id, {sold: true});
  if (!newRecord.sold) {
    alert('Verkauf fehlgeschlagen')
    return
  }

  const payload = {
    positions: [{
      amount: 1,
      itemId: vvkItem.id,
    }]
  }
  console.log(payload)
  fetch('/api/v1/new-transactions', {
    method: 'POST',
    body: JSON.stringify(payload),
    headers: {
      'Content-Type': 'application/json',
    },
  })
  dialog = false
}

</script>

<template>
  <Scanner class="full-screen" @on-scan="onScan($event)"/>
  <v-btn @click="onScan('.../oe1q37fqumvu7n4')">TEST</v-btn>
  <v-dialog v-model="dialog">
    <v-card class="pa-3 mx-auto" width="300px">
      <v-card-title class="mx-auto">{{ id }}</v-card-title>
      <v-btn class="mx-auto" width="128px" color="primary" @click="sell()">Verkaufen</v-btn>
    </v-card>
  </v-dialog>
</template>
