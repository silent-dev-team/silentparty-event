<script setup lang="ts">
import { RecordModel } from 'pocketbase';

definePageMeta({
  layout: 'admin',
});

const pb = usePocketbase();
// const checkout = useCheckout();
const notifyer = useNotifyer();

let vvkItem: RecordModel
try {
  vvkItem = await pb.collection('shop_items').getFirstListItem<ShopItemRecord>('title = "VVK Ticket"')
} catch (e) {
  notifyer.notify('VVK Ticket nicht in Datenbank gefunden', 'error')
  throw e
}

let dialog = $ref(false)
let id = $ref('')
let ticket:TicketRecord
let scannerReset = $ref(false)

async function onScan(url: string) {
  console.log(url)
  const ticketId = url.split('/').pop()
  if (!ticketId) return;
  id = ticketId
  ticket = await pb.collection('tickets').getOne<TicketRecord>(id);
  dialog = true
}

async function sell() {
  const updatedTicket = await pb.collection('tickets').update<TicketRecord>(id, {sold: true});
  if (!updatedTicket.sold) {
    notifyer.notify('Verkauf fehlgeschlagen: Ticket konnte nicht als verkauft markiert werden', 'error')
    return
  }
  pb.checkout([{
    amount: 1,
    itemId: vvkItem.id,
  }]).then(() => {
    dialog = false
    reset()
  }).catch(() => {
    notifyer.notify('Verkauf fehlgeschlagen: Transaktion konnte nicht erstellt werden', 'error')
  })
}

function reset() {
  if (!scannerReset) return;
  setTimeout(() => scannerReset = true, 800)
}

</script>

<template>
  <Scanner class="full-screen" :overlaypath="Overlay.Ticket" :reset="scannerReset" @update:reset="scannerReset = $event" @on-scan="onScan($event)"/>
  <v-dialog v-model="dialog" :close-on-back="true" :persistent="true">
    <v-card class="pa-3 mx-auto" width="300px">
      <v-btn style="position: absolute;" icon="mdi-close" size="sm" @click="dialog = false; reset()" /> 
      <v-card-title class="mx-auto">{{ id }}</v-card-title>
      <v-card-text v-if="ticket.sold" class="mx-auto">
        Ticket bereits verkauft!
      </v-card-text>
      <v-btn class="mx-auto" width="128px" color="success" :disabled="ticket.sold" @click="sell()">Verkaufen</v-btn>
    </v-card>
  </v-dialog>
</template>
