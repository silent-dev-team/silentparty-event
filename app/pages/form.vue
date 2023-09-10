
<template>
  <v-container class="bg-blue-grey-lighten-5">
  <v-card class="pa-5 mx-auto my-3" rounded="8px" max-width="600px">
    <v-card-title><h1>Anmeldung</h1></v-card-title>
    <v-card-text>Bitte gib alle Daten so an wie auf deinem Personalausweis. Ansonsten ist dein Ticket ungültig.</v-card-text>
    <br>
    <customer-form v-model="form" @submit="submit(form)" submitText="QR-Code erstellen"/>
  </v-card>
</v-container>
<v-dialog
      v-model="success"
      width="auto"
    >
      <v-card>
        <v-card-title>Deine Daten</v-card-title>
        <v-card-item width="100%">
          <v-img width="800px" :src="qrCodeImageUrl" alt="QR-Code" />
        </v-card-item>
        <v-card-text>
          Bitte zeige diesen QR-Code an der Kasse vor.
        </v-card-text>
        <v-card-actions>

        </v-card-actions>
      </v-card>
    </v-dialog>
</template>

<script setup lang="ts">
//import axios from 'axios';
import QRCode from 'qrcode';
import { CustomerData } from '@/types/general';

let form = $ref<CustomerData>({
  firstName: '',
  lastName: '',
  street: '',
  housenumber: '',
  zipCode: '',
  place: '',
  email: '',
});

let success = $ref(false);

let qrCodeImageUrl = $ref("");

async function submit(form: CustomerData) {
  try {
    const qrCodeData = JSON.stringify(form);
    qrCodeImageUrl = await generateQRCode(qrCodeData);
    success = true;
  } catch (error) {
    console.error(error);
  }
}

  // Funktion zur Erstellung des QR-Codes
async function generateQRCode(data: string): Promise<string> {
  try {
    const qrCodeImageUrl = await QRCode.toDataURL(data);
    console.log('QR-Code erstellt:', qrCodeImageUrl);
    return qrCodeImageUrl;
  } catch (error) {
    console.error('Fehler beim Erstellen des QR-Codes:', error);
    return '';
  }
}
</script>~/types