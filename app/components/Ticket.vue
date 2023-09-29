<script setup lang="ts">
import { RecordModel, RecordOptions } from 'pocketbase';

const props = defineProps({
  id: {
    type: String,
    required: true
  },
  readonly: {
    type: Boolean,
    default: false
  }
});

const pb = usePocketbase();

let ticket = {} as RecordModel & Ticket;

let form = $ref<CustomerData>({
  firstName: '',
  lastName: '',
  street: '',
  housenumber: '',
  zipCode: '',
  place: '',
  email: '',
});
console.log(props.id);

async function refreshTicket() {
  ticket = await pb.collection('tickets').getOne<RecordModel & Ticket>(props.id);
  form = {
    firstName: ticket.firstName,
    lastName: ticket.lastName,
    street: ticket.street,
    housenumber: ticket.housenumber,
    zipCode: ticket.zipCode,
    place: ticket.place,
    email: ticket.email,
  };
}

async function updateTicket() {
  const rec = await pb.collection('tickets').update<RecordModel & Ticket>(props.id, form);
  console.log(rec);
  await refreshTicket();
}

await refreshTicket();
</script>

<template>
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
    <v-card-text>
      <customer-form v-model="form" :readonly="readonly" :disabled="form == ticket" @submit="updateTicket()" submitText="Daten speichern"/>
    </v-card-text>
  </v-card>
</template>