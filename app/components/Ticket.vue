<script setup lang="ts">
import { RecordModel } from 'pocketbase';

const props = defineProps({
  id: {
    type: String,
    required: true
  }
});

const runtimeConfig = useRuntimeConfig();
const pb_url = runtimeConfig.public.pocketbase;

const pb = usePocketbase();

let wrongPinError = $ref(false);
let dataChanged = $ref(false);

let renderComponent = $ref(true);

let form = $ref<CustomerData>({
  firstName: '',
  lastName: '',
  street: '',
  housenumber: '',
  zipCode: '',
  place: '',
  email: '',
});

let ticketPin = $ref('');

let showPinDialog = $ref(!pb.authStore.isAdmin);

function handlePin(pin: string) {
  ticketPin = pin;
  refreshTicket();
}

function wrongPin() {
  wrongPinError = true;
  ticketPin = '';
  showPinDialog = true;
}

async function refreshTicket() {
  let ticket:RecordModel & Ticket;
  if (pb.authStore.isAdmin) {
    ticket = await pb.collection('tickets').getOne<RecordModel & Ticket>(props.id);
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
  form = {
    firstName: ticket.firstName,
    lastName: ticket.lastName,
    street: ticket.street,
    housenumber: ticket.housenumber,
    zipCode: ticket.zipCode,
    place: ticket.place,
    email: ticket.email,
  };
  renderComponent = false;
  nextTick(() => {
    renderComponent = true;
  });
}

async function updateTicket() {
  let payload = form as Partial<Ticket>;
  payload.filled = true;
  const rec = await pb.collection('tickets').update<RecordModel & Ticket>(props.id, payload);
  console.log(rec);
  await refreshTicket();
  dataChanged = true;
}

if (pb.authStore.isAdmin) await refreshTicket();
</script>

<template>
  <v-snackbar v-model="wrongPinError" color="error" location="top">
    <span>Pin falsch</span>
  </v-snackbar>
  <v-snackbar v-model="dataChanged" color="success" location="top">
    <span>Daten gespeichert</span>
  </v-snackbar>
  <v-dialog v-model="showPinDialog" :persistent="true">
    <v-card class="mx-auto pa-10">
      <PinField @update="handlePin($event)" :reset="ticketPin === ''"/>
    </v-card>
  </v-dialog>
  <v-card class="ma-5 pa-5 mx-auto" maxWidth="800px">
    <v-card-title>
      <h1>Ticket {{ id }}</h1>
    </v-card-title>
    <v-divider></v-divider>
    <v-card-text v-if="renderComponent">
      <CustomerForm v-model="form" @submit="updateTicket()" submitText="Daten speichern"/>
    </v-card-text>
  </v-card>
</template>