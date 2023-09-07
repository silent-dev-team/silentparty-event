
<template>
    <h1>Tickets Silent Party</h1>
    <p>Bitte gib alle Daten so an wie auf deinem Personalausweis. Ansonsten ist dein Ticket ungültig.</p>
    <br>
    <form @submit.prevent="submit(form)">

      <div class="d-flex flex-column justify-center" style="width: 100%;">
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

        <v-btn class="mx-3 h-100" type="submit" color="primary">Abschicken</v-btn>

        <p v-if="errors">Etwas ist schiefgelaufen.</p>
        <p v-if="success">Abgeschickt.</p>     
      </div>
    </form>
</template>

<script setup lang="ts">
//import axios from 'axios';
import QRCode from 'qrcode';

interface Form {
  firstName: string;
  lastName: string;
  street: string;
  housenumber: string;
  zipCode: string;
  place: string;
  email: string;
}

let form = $ref<Form>({
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

// Regex-Patterns 
const zipCodePattern = /^\d{5}$/; 
const namePattern = /^[A-Za-zäöüÄÖÜß\-~']+$/;
const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/; //passt noch nicht

async function submit(form: Form) {
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
    const qrCodeImageUrl = await generateQRCode(qrCodeData);
    showQRCode(qrCodeImageUrl);
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
    return qrCodeImageUrl;
  } catch (error) {
    console.error('Fehler beim Erstellen des QR-Codes:', error);
    return '';
  }
}
// Funktion zum Anzeigen des QR-Codes in einem Dialogfenster
function showQRCode(qrCodeImageUrl: string) {
  const qrCodeImage = new Image();
  qrCodeImage.src = qrCodeImageUrl;

  const qrCodeDialog = window.open('', '_blank', 'height=400,width=400');
  if (qrCodeDialog) {
    qrCodeDialog.document.write('<html><body style="text-align: center;">');
    qrCodeDialog.document.write('<h2>Dein QR-Code:</h2>');
    qrCodeDialog.document.write('<img src="' + qrCodeImageUrl + '">');
    qrCodeDialog.document.write('</body></html>');
  } else {
    console.error('Fehler beim Öffnen des QR-Code-Dialogs. Stelle sicher, dass der Popup-Blocker deaktiviert ist.');
  }
}
</script>

<style scoped>

form{
  width: 30rem;
  display: block;
  flex-direction: column;
  gap: 1rem;
}

.input-wrapper {
  display: flex;
  flex-direction: column;
  width: 100%;
  label{
    font-size: 1rem;
  }

  input, textarea{
    border: black;
    border-radius: 0.5rem;
    margin: 0.2rem;
    background-color: white;
    padding: 1rem;
    color: black;
  }

}

button{
  border: none;
  padding: 1rem;
  border-radius: 0.5rem;
  background-color: red;
}
</style>
