<script setup lang="ts">
import { RecordModel } from 'pocketbase';
import { useNotifyer } from '../stores/notifyStore';

const props = defineProps({
  id: {
    type: String,
    required: true
  },
  submitText: {
    type: String,
    default: 'Speichern'
  }
});

const emit = defineEmits<{
  update: [value: void]
}>()

const runtimeConfig = useRuntimeConfig();
const pb_url = runtimeConfig.public.pocketbase;

const pb = usePocketbase();
const notifyer = useNotifyer();

let hideAll = $ref(false);
let renderComponent = $ref(true);

let form = $ref<ICustomerData>(initCustomerData());

let ticketPin = $ref('');

let showPinDialog = $ref(!pb.authStore.isAdmin);

function handlePin(pin: string) {
  ticketPin = pin;
  refreshTicket();
}

function wrongPin() {
  notifyer.notify('Pin falsch', 'error');
  ticketPin = '';
  showPinDialog = true;
}

async function refreshTicket() {
  let ticket:RecordModel & ITicket;
  if (pb.authStore.isAdmin) {
    ticket = await pb.collection('tickets').getOne<RecordModel & ITicket>(props.id);
  } else {
    ticket = await fetch(`${pb_url}/api/collections/tickets/records/${props.id}?pin=${ticketPin}`)
      .then(res => {
        if (res.status === 401) {
          showPinDialog = true;
          throw new Error('Unauthorized');
        }
        if (res.status === 404) {
          wrongPin();
          throw new Error('Ticket not found, pin wrong or ticket not sold');
        }
        if (res.status !== 200) throw new Error('Unknown error');
        showPinDialog = false;
        return res.json();
      })
      .catch(err => {
        console.error(err);
        ticketPin = '';
      });
  }
  form = transformTicketToCustomerData(ticket);
  renderComponent = false;
  nextTick(() => {
    renderComponent = true;
  });
}

async function updateTicket() {
  let payload = form as Partial<ITicket>;
  payload.filled = true;
  const rec = await pb.collection('tickets').update<RecordModel & ITicket>(props.id, payload);
  console.log(rec);
  await refreshTicket();
  notifyer.notify('Daten gespeichert. Wir sehen und auf der Party :)', 'success');
  emit('update');
}

const resp = await fetch(`${pb_url}/api/collections/tickets/exists/${props.id}`)
if (resp.status === 404) {
  hideAll = true;
} 

if (pb.authStore.isAdmin) await refreshTicket();

</script>

<template>
  <v-card v-if="hideAll" color="grey" class="mx-auto text-center">
    <v-card-title >
      Kein Ticket gefunden.
    </v-card-title>
    <v-card-text>
      Versuche deine Ticket erneut zu scannen oder wende dich an das Team unter <br> info@silentparty-hannover.de
    </v-card-text>
  </v-card>
  <v-dialog v-if="!hideAll" v-model="showPinDialog" :persistent="true">
    <v-card class="mx-auto pa-10" max-width="380px">
      <PinField @update="handlePin($event)" :reset="ticketPin === ''"/>
    </v-card>
  </v-dialog>
  <v-card class="pa-2" v-if="!hideAll" maxWidth="800px">
    <v-card-title class="pb-0">
      <h1 class="text-h4 pb-0">Dein Ticket </h1>
    </v-card-title>
    <v-card-subtitle>
      <h2 class="mb-1 text-h6">{{ id }}</h2>
    </v-card-subtitle>
    <v-divider></v-divider>
    <v-card-text v-if="renderComponent">
      <CustomerForm class="mt-5" v-model="form" @submit="updateTicket()" :submitText="submitText"/>
    </v-card-text>
  </v-card>
</template>