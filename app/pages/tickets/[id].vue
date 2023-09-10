<script setup lang="ts">
import { RecordModel, RecordOptions } from 'pocketbase';

const route = useRoute();
const pb = usePocketbase();
const id = route.params.id as string;

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
console.log(id);

async function refreshTicket() {
  ticket = await pb.collection('tickets').getOne<RecordModel & Ticket>(id);
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
  const rec = await pb.collection('tickets').update<RecordModel & Ticket>(id, form);
  console.log(rec);
  await refreshTicket();
}

await refreshTicket();
</script>

<template>
  <v-card class="ma-5 pa-5">
    <v-card-title>
      <h1>Ticket bearbeiten</h1>
    </v-card-title>
    <v-divider></v-divider>
    <v-card-text>
      Hier kannst du dein Ticket bearbeiten.
    </v-card-text>
    <customer-form v-model="form" :disabled="form == ticket" @submit="updateTicket()" submitText="Ticket ändern"/>
  </v-card>
</template>