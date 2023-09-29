<script setup lang="ts">
import { RecordModel } from 'pocketbase';

const props = defineProps({
  id: {
    type: String,
    required: true
  }
});

const appConfig = useAppConfig();

const pb = usePocketbase();

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
  showPinDialog = false;
}

async function refreshTicket() {
  let ticket:RecordModel & Ticket;
  if (pb.authStore.isAdmin) {
    ticket = await pb.collection('tickets').getOne<RecordModel & Ticket>(props.id);
  } else {
    ticket = await fetch(`${appConfig.pocketbase.url}/api/collections/tickets/records/${props.id}?pin=${ticketPin}`)
      .then(res => res.json())
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
}

if (pb.authStore.isAdmin) await refreshTicket();
</script>

<template>
  <v-dialog v-model="showPinDialog">
    <v-card class="mx-auto pa-10">
      <PinField @update="handlePin($event)"/>
    </v-card>
  </v-dialog>
  <v-card class="ma-5 pa-5 mx-auto" maxWidth="800px">
    <v-card-title>
      <h1>Ticket {{ id }}</h1>
    </v-card-title>
    <v-divider></v-divider>
    <!-- <div class="mb-3">
      <v-card-text class="">
        <h3><v-icon>mdi-check</v-icon>Bezahlt</h3>
        <h3><v-icon>mdi-check</v-icon>Kontolliert</h3>
        <h3><v-icon>mdi-check</v-icon>Mit Kopfhörer verknüpft</h3>
      </v-card-text>
      <v-btn color="primary" @click="ticket.used = true">Ticket entwerten</v-btn>
    </div> -->
    <v-divider></v-divider>
    <v-card-text v-if="renderComponent">
      <CustomerForm v-model="form" @submit="updateTicket()" submitText="Daten speichern"/>
    </v-card-text>
  </v-card>
</template>