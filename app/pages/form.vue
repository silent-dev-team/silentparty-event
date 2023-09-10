
<template>
  <v-container class="bg-blue-grey-lighten-5">
  <v-card class="pa-5 mx-auto my-3" rounded="8px" max-width="600px" :loading="waiting">
    <v-card-title><h1>Anmeldung</h1></v-card-title>
    <v-card-text>Bitte gib alle Daten so an wie auf deinem Personalausweis. Ansonsten ist dein Ticket ungültig.</v-card-text>
    <br>
    <form class="px-3" @submit.prevent="submit(form)">

      <div class="d-flex flex-column justify-center">
        <v-text-field
          v-model="form.firstName"
          label="Vorname"
          required
          :rules="[namePattern.test(form.firstName) || 'Bitte gib einen gültigen Vornamen ein.']"
        ></v-text-field>
    
        <v-text-field
          v-model="form.lastName"
          label="Nachname"
          required
          :rules="[namePattern.test(form.lastName) || 'Bitte gib einen gültigen Nachnamen ein.']"
        ></v-text-field>

        <v-text-field
          v-model="form.street"
          label="Straße"
          required
        ></v-text-field>

        <v-text-field
          v-model="form.housenumber"
          label="Hausnummer"
          required
        ></v-text-field>

        <v-text-field
          v-model="form.zipCode"
          label="Postleitzahl"
          required
          :rules="[zipCodePattern.test(form.zipCode) || 'Bitte gib eine gültige Postleitzahl ein.']"
        ></v-text-field>

        <v-text-field
          v-model="form.place"
          label="Ort"
          required
          :rules="[namePattern.test(form.place) || 'Bitte gib einen gültigen Ort ein.']"
        ></v-text-field>

        <v-text-field
          v-model="form.email"
          label="E-Mail"
          required
          :rules="[emailPattern.test(form.email) || 'Bitte gib eine gültige E-Mail ein.']"
        ></v-text-field>

        <v-btn class="mx-5 my-3 h-100 pa-3" type="submit" color="primary">Abschicken</v-btn>

        <p v-if="errors">Etwas ist schiefgelaufen.</p>
        <p v-if="success">Abgeschickt.</p>     
      </div>
    </form>
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
import { CustomerData } from 'types/form';

let form = $ref<CustomerData>({
  firstName: '',
  lastName: '',
  street: '',
  housenumber: '',
  zipCode: '',
  place: '',
  email: '',
});

let waiting = $ref(false);
let errors = $ref(false);
let success = $ref(false);

let qrCodeImageUrl = $ref("");

// Regex-Patterns 
const zipCodePattern = /^\d{5}$/; 
const namePattern = /^[A-Za-zäöüÄÖÜß\-~']+$/;
const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/; //passt noch nicht

async function submit(form: CustomerData) {
  waiting = true;
  errors = false;
  success = false;

  // Validierung der Postleitzahl
  if (!zipCodePattern.test(form.zipCode)) {
    errors = true;
    return;
  }

  // Validierung des Vornamens
  if (!namePattern.test(form.firstName)) {
    errors = true;
    return;
  }

  // Validierung des Nachnamens
  if (!namePattern.test(form.lastName)) {
    errors = true;
    return;
  }

  // Validierung des Ortes
  if (!namePattern.test(form.place)) {
    errors = true;
    return;
  }


  try {
    //const response = await axios.post('/api/form-submit', form);
    //console.log(response.data);
    
    // QR-Code erstellen und anzeigen
    const qrCodeData = JSON.stringify(form);
    qrCodeImageUrl = await generateQRCode(qrCodeData);
    success = true;
  } catch (error) {
    console.error(error);
    errors = true;
  } finally {
    waiting = false;
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
</script>
../types/form