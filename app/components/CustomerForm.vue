<script setup lang="ts">
type CancleButtonType = null | String;
const props = defineProps({
  modelValue: {
    type: Object as PropType<ICustomerData>,
    required: true
  },
  submitText: {
    type: String,
    default: 'Abschicken'
  },
  cancelText: {
    type: [null, String] as PropType<CancleButtonType>,
    default: null
  },
  disabled: {
    type: Boolean,
    default: false
  },
  readonly: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['update:modelValue', 'submit', 'cancel'])

let form = $ref<ICustomerData>(props.modelValue);

// watch(() => form, (value) => {
//   form = value;
//   emit('update:modelValue', form);
// })

// Regex-Patterns 
const zipCodePattern = /^\d{5}$/; 
const namePattern = /^[A-Za-zäöüÄÖÜß\-~'_ !?]+$/;
const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/; //passt noch nicht

let errors = $ref(false);
let success = $ref(false);

async function submit(form: ICustomerData) {
  errors = false;
  success = false;

  // trimmen
  form.email = form.email.toLowerCase().trim();
  form.firstName = form.firstName.trim();
  form.lastName = form.lastName.trim();
  form.place = form.place.trim();
  form.street = form.street.trim();
  form.housenumber = form.housenumber.trim();

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
    await emit('submit')
  } catch (e) {
    errors = true;
  }
}
</script>

<template>
  <v-form @submit.prevent="submit(form)" autocomplete="on">
    <v-container class="ma-0 pa-0">
        <v-row no-gutters >
          <v-col cols="12" sm="6" class="my-0 py-0">
            <v-text-field
              v-model="form.firstName"
              label="Vorname"
              required
              autocomplete="given-name"
              :readonly="readonly"
              :rules="[namePattern.test(form.firstName) || 'Bitte gib einen gültigen Vornamen ein.']"
            ></v-text-field>
          </v-col>

          <v-col cols="12" sm="6" class="my-0 py-0">
            <v-text-field
              v-model="form.lastName"
              label="Nachname"
              required
              autocomplete="family-name"
              :readonly="readonly"
              :rules="[namePattern.test(form.lastName) || 'Bitte gib einen gültigen Nachnamen ein.']"
            ></v-text-field>
          </v-col>
        </v-row>

        <v-row no-gutters >
          <v-col cols="12" sm="8" class="my-0 py-0">
          <v-text-field
            v-model="form.street"
            label="Straße"
            required
            id="street"
            autocomplete="street-address"
            :readonly="readonly"
          ></v-text-field>
          </v-col>

          <v-col cols="12" sm="4" class="my-0 py-0">
          <v-text-field
            v-model="form.housenumber"
            label="Hausnummer"
            required
            id="housenumber"
            :readonly="readonly"
          ></v-text-field>
          </v-col>
        </v-row>

        <v-row no-gutters >
          <v-col cols="12" sm="4" class="my-0 py-0">
            <v-text-field
              v-model="form.zipCode"
              label="Postleitzahl"
              required
              autocomplete="postal-code"
              :readonly="readonly"
              :rules="[zipCodePattern.test(form.zipCode) || 'Bitte gib eine gültige Postleitzahl ein.']"
            ></v-text-field>
          </v-col>

          <v-col cols="12" sm="8" class="my-0 py-0">
            <v-text-field
              v-model="form.place"
              label="Ort"
              required
              autocomplete="address-level2"
              :readonly="readonly"
              :rules="[namePattern.test(form.place) || 'Bitte gib einen gültigen Ort ein.']"
            ></v-text-field>
          </v-col>
        </v-row>

        <v-row no-gutters >
          <v-col cols="12" sm="12" class="my-0 py-0">
            <v-text-field
              v-model="form.email"
              label="E-Mail"
              required
              autocomplete="email"
              :readonly="readonly"
              :rules="[emailPattern.test(form.email) || 'Bitte gib eine gültige E-Mail ein.']"
            ></v-text-field>
          </v-col>
        </v-row>
        <div v-if="!readonly" class="w-100 d-flex justify-center">
            <v-btn v-if="cancelText" class="ma-3 h-100 pa-3" color="red" @click="emit('cancel')">{{ props.cancelText }}</v-btn>
            <v-btn :disabled="disabled || form.firstName == '' || form.lastName == ''" class="ma-3 h-100 pa-3" type="submit" color="primary">{{ props.submitText }}</v-btn>
        </div>
        <p v-if="errors">Etwas ist schiefgelaufen.</p>
        <p v-if="success">Abgeschickt.</p>     
    </v-container>
  </v-form>
</template>

<style scoped>
</style>