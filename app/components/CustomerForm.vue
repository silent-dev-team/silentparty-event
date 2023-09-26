<script setup lang="ts">
const props = defineProps({
  modelValue: {
    type: Object as PropType<CustomerData>,
    required: true
  },
  submitText: {
    type: String,
    default: 'Abschicken'
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

const emit = defineEmits(['update:modelValue', 'submit'])

let form = $ref<CustomerData>(props.modelValue);

// watch(() => form, (value) => {
//   form = value;
//   emit('update:modelValue', form);
// })

// Regex-Patterns 
const zipCodePattern = /^\d{5}$/; 
const namePattern = /^[A-Za-zäöüÄÖÜß\-~']+$/;
const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/; //passt noch nicht

let errors = $ref(false);
let success = $ref(false);

async function submit(form: CustomerData) {
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
    await emit('submit')
  } catch (e) {
    errors = true;
  }
}
</script>

<template>
  <form @submit.prevent="submit(form)">
        <v-row no-gutters class="mb-3">
          <v-col cols="12" sm="6">
            <v-text-field
              v-model="form.firstName"
              label="Vorname"
              required
              :readonly="readonly"
              :rules="[namePattern.test(form.firstName) || 'Bitte gib einen gültigen Vornamen ein.']"
            ></v-text-field>
          </v-col>

          <v-col cols="12" sm="6">
            <v-text-field
              v-model="form.lastName"
              label="Nachname"
              required
              :readonly="readonly"
              :rules="[namePattern.test(form.lastName) || 'Bitte gib einen gültigen Nachnamen ein.']"
            ></v-text-field>
          </v-col>
        </v-row>

        <v-row no-gutters class="mb-3">
          <v-col cols="12" sm="8">
          <v-text-field
            v-model="form.street"
            label="Straße"
            required
            :readonly="readonly"
          ></v-text-field>
          </v-col>

          <v-col cols="12" sm="4">
          <v-text-field
            v-model="form.housenumber"
            label="Hausnummer"
            required
            :readonly="readonly"
          ></v-text-field>
          </v-col>
        </v-row>

        <v-row no-gutters class="mb-3">
          <v-col cols="12" sm="4">
            <v-text-field
              v-model="form.zipCode"
              label="Postleitzahl"
              required
              :readonly="readonly"
              :rules="[zipCodePattern.test(form.zipCode) || 'Bitte gib eine gültige Postleitzahl ein.']"
            ></v-text-field>
          </v-col>

          <v-col cols="12" sm="8">
            <v-text-field
              v-model="form.place"
              label="Ort"
              required
              :readonly="readonly"
              :rules="[namePattern.test(form.place) || 'Bitte gib einen gültigen Ort ein.']"
            ></v-text-field>
          </v-col>
        </v-row>

        <v-row no-gutters class="mb-3">
          <v-col cols="12" sm="12">
            <v-text-field
              v-model="form.email"
              label="E-Mail"
              required
              :readonly="readonly"
              :rules="[emailPattern.test(form.email) || 'Bitte gib eine gültige E-Mail ein.']"
            ></v-text-field>
          </v-col>
        </v-row>
        <div v-if="!readonly" class="w-100 d-flex justify-content-center">
            <v-btn :disabled="disabled || form.firstName == '' || form.lastName == ''" class="mx-auto my-3 h-100 pa-3" type="submit" color="primary">{{ props.submitText }}</v-btn>
        </div>
        <p v-if="errors">Etwas ist schiefgelaufen.</p>
        <p v-if="success">Abgeschickt.</p>     
      </form>
</template>