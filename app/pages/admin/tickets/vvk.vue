<script setup lang="ts">
import type { RecordModel } from 'pocketbase';

definePageMeta({
  layout: 'admin',
});

const pb = usePocketbase();
// const checkout = useCheckout();
const notifyer = useNotifyer();
const settingsStore = useSettingsStore();

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

async function onScan(scan: string) {
  console.log(scan)
  if (scan == enter) {
    if (ticket && id) sell();
    return
  }
  const ticketId = scan.split('/').pop()
  if (!ticketId) return;
  try {
    ticket = await pb.collection('tickets').getOne<TicketRecord>(ticketId);
  } catch (e) {
    notifyer.notify('Ticket nicht gefunden', 'error')
    reset()
    return
  }
  if (!ticket) {
    notifyer.notify('Ticket nicht gefunden', 'error')
    reset()
    return
  }
  id = ticket.id
  if (settingsStore.noInteraction) {
    sell();
  } else {
    dialog = true
  }
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

  <v-btn 
    style="position: absolute; right: 1rem; top: 5rem;z-index: 100;"
    :color="settingsStore.noInteraction ? 'red lighten-2' : 'transparent-white'"
    icon="mdi-forum-remove"
    size="large"
    @click="settingsStore.toggleNoInteraction()"
    >
  </v-btn>
  <Scanner class="full-screen" :overlaypath="Overlay.Ticket" :reset="scannerReset" @update:reset="scannerReset = $event" @scan="onScan($event)"/>
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
