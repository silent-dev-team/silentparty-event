
<template>
  <NuxtLayout name="baseline">
    <h1>Tickets Silent Party</h1>
    <p>Bitte gib alle Daten so an wie auf deinem Personalausweis. Ansonsten ist dein Ticket ungültig.</p>
    <br>
    <form @submit.prevent="submit(form)">
  
      <div class="input-wrapper">
        <label for="firstName">Vorname:</label>
        <input required id="firstName" type="text" name="firstName" v-model="form.firstName" />
        <div v-if="errors && !namePattern.test(form.firstName)">Bitte gib einen gültigen Vornamen ein.</div>
      </div>

      <div class="input-wrapper">
        <label for="lastName">Nachname:</label>
        <input required id="lastName" type="text" name="lastName" v-model="form.lastName" />
        <div v-if="errors && !namePattern.test(form.lastName)">Bitte gib einen gültigen Nachnamen ein.</div>
      </div>

      <div class="input-wrapper">
        <label for="street">Straße:</label>
        <input required id="street" type="text" name="street" v-model="form.street" />
      </div>

      <div class="input-wrapper">
        <label for="housenumber">Hausnummer:</label>
        <input required id="housenumber" type="text" name="housenumber" v-model="form.housenumber" />
      </div>

      <div class="input-wrapper">
        <label for="zipCode">Postleitzahl:</label>
        <input required id="zipCode" type="text" name="zipCode" v-model="form.zipCode" />
        <div v-if="errors && !zipCodePattern.test(form.zipCode)">Bitte gib eine gültige Postleitzahl ein.</div>
      </div>

      <div class="input-wrapper">
        <label for="place">Ort:</label>
        <input required id="place" type="text" name="place" v-model="form.place" />
        <div v-if="errors && !namePattern.test(form.place)">Bitte gib eine gültigen Ort ein.</div>
      </div>

      <div class="input-wrapper">
        <label for="email">E-Mail:</label>
        <input required id="email" type="email" name="email" v-model="form.email" />
        <div v-if="errors && !namePattern.test(form.email)">Bitte gib eine gültigen E-Mail ein.</div>
      </div>

      <button type="submit">
        <template v-if="waiting">Laden</template>
        <template v-if="!waiting">Abschicken</template>
      </button>
      <p v-if="errors">Etwas ist schiefgelaufen.</p>
      <p v-if="success">Abgeschickt.</p>     
    </form>
  </NuxtLayout>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import axios from 'axios';
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
  email: ''
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
    const response = await axios.post('/api/form-submit', form);
    console.log(response.data);
    success = true;

     // QR-Code erstellen und anzeigen
     const qrCodeData = JSON.stringify(form);
    const qrCodeImageUrl = await generateQRCode(qrCodeData);
    showQRCode(qrCodeImageUrl);
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

<style>

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
