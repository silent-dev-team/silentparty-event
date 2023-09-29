<script setup lang="ts">
//import axios from 'axios';
import { RefType } from '@vue-macros/reactivity-transform/macros';
import QRCode from 'qrcode';

let form = $ref<CustomerData>({
  firstName: '',
  lastName: '',
  street: '',
  housenumber: '',
  zipCode: '',
  place: '',
  email: '',
});

localStorage.getItem('customerData') && (form = JSON.parse(localStorage.getItem('customerData') as string));

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
    localStorage.setItem('customerData', JSON.stringify($$(form).value));
    return qrCodeImageUrl;
  } catch (error) {
    console.error('Fehler beim Erstellen des QR-Codes:', error);
    return '';
  }
}
</script>


<template>
  <div class="bg-blue-grey-lighten-5" style="min-height: 100vh; width: 100%; padding: 3rem;">
    <v-card class="pa-5 mx-auto" rounded="8px" max-width="800px">
      <v-card-title>
        <h1>Anmeldung</h1>
      </v-card-title>
      <v-card-text>
        Dieses Formular dient dazu deine Daten zu erfassen und einen QR-Code zu erstellen. Dieser QR-Code wird an der
        Kasse gescannt und du erhälst dein Ticket.
        <br>
        Bitte gib alle Daten so an wie auf deinem Personalausweis. Ansonsten ist dein Ticket ungültig.
      </v-card-text>
      <br>
      <customer-form v-model="form" @submit="submit(form)" submitText="QR-Code erstellen" />
    </v-card>
  </div>
  <v-dialog v-model="success" width="auto">
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