<script setup lang="ts">
definePageMeta({
  layout: 'admin',
});

const pb = usePocketbase()

let showScanner = $ref(false)

let form = $ref<CustomerData>({
  firstName: '',
  lastName: '',
  street: '',
  housenumber: '',
  zipCode: '',
  place: '',
  email: '',
});

function loadDataFromQRString(s:string){
  const content = JSON.parse(s)
  console.log(content)
  form.firstName = content.firstName
  form.lastName = content.lastName
  form.street = content.street
  form.housenumber = content.housenumber
  form.zipCode = content.zipCode
  form.place = content.place
  form.email = content.email
  showScanner = false
}

async function createTicket(){
  const rec = await pb.collection('tickets').create<TicketRecord>(form)
  console.log(rec)
}

</script>

<template>
  <v-dialog width="90%" v-model="showScanner">
    <v-card class="pa-3">
      <scanner v-if="showScanner" @onScan="loadDataFromQRString($event)"></scanner>
    </v-card>
  </v-dialog>
    <v-card class="pa-5 mx-auto" maxWidth="800px">
      <v-card-title>
        Ticket erstellen
      </v-card-title>
      <v-divider></v-divider>
      <v-card-text>
        Hier kannst du dein Ticket erstellen.
      </v-card-text>
      <customer-form v-model="form" @submit="createTicket" submitText="Ticket erstelen"/>
    </v-card>
  <v-btn 
    style="position: fixed; right: 1rem; bottom: 4rem;"
    color="primary"
    icon="mdi-qrcode-scan"
    size="large"
    @click="showScanner = true"
  ></v-btn>
</template>