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
      <div v-if="ticket.used">
        <v-list lines="one">
          <div class="d-flex">
            <v-list-item
              title="Vorname"
              :subtitle="ticket.firstName"
            ></v-list-item>
            <v-list-item
              title="Nachname"
              :subtitle="ticket.lastName"
            ></v-list-item>
          </div>
          <div class="d-flex">
            <v-list-item
              title="Straße"
              :subtitle="ticket.street"
            ></v-list-item>
          </div>
        </v-list>
      </div>
    </v-card-text>
    <div v-if="!ticket.used">
      <customer-form v-model="form" :disabled="form == ticket" @submit="updateTicket()" submitText="Ticket ändern"/>
    </div>
  </v-card>
</template>