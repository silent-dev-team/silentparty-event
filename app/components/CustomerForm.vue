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

const emit = defineEmits(['update:modelValue', 'submit', 'cancel', 'focus', 'blur'])

const form = ref<ICustomerData>(props.modelValue);

// watch(() => form, (value) => {
//   form = value;
//   emit('update:modelValue', form);
// })

// Regex-Patterns 
const zipCodePattern = /^\d{5}$/; 
const namePattern = /^[A-Za-zäöüÄÖÜß\-~'_ !?]+$/;
const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/; //passt noch nicht
const datePattern = /^\d{4}-\d{2}-\d{2}$/;

let errors = ref(false);
let success = ref(false);

const birthdate = ref<string|undefined>( form.value.birthdate ? new Date(form.value.birthdate).toISOString().substring(0,10) : undefined );

async function submit(form: ICustomerData) {
  errors.value = false;
  success.value = false;


  // Validierung des Geburtsdatums
  if (!datePattern.test(birthdate.value)) {
    errors.value = true;
    return;
  }

  // trimmen
  form.email = (form.email || '').toLowerCase().trim();
  form.firstName = form.firstName.trim();
  form.lastName = form.lastName.trim();
  form.place = form.place.trim();
  form.street = form.street.trim();
  form.housenumber = form.housenumber.trim();
  form.zipCode = form.zipCode.trim();
  form.birthdate = birthdate.value.trim();

  // Validierung der Postleitzahl
  if (!zipCodePattern.test(form.zipCode)) {
    errors.value = true;
    return;
  }

  // Validierung des Vornamens
  if (!namePattern.test(form.firstName)) {
    errors.value = true;
    return;
  }

  // Validierung des Nachnamens
  if (!namePattern.test(form.lastName)) {
    errors.value = true;
    return;
  }

  // Validierung des Ortes
  if (!namePattern.test(form.place)) {
    errors.value = true;
    return;
  }

  // Validierung der E-Mail
  if (!emailPattern.test(form.email)) {
    errors.value = true;
    return;
  }

  // Umwandlung des Geburtsdatums in ISO-Format
  const [day, month, year] = form.birthdate.split('.');
  const isoDate = new Date(`${year}-${month}-${day}`);
  if (isNaN(isoDate.getTime())) {
    errors.value = true;
    return;
  }


  try {
    emit('submit')
  } catch (e) {
    errors.value = true;
  }
}

const formRef = ref();

const {pause, resume} = useKeyLogger((logs) => {
  if (logs === enter) {
    formRef.value?.validate().then((isValid: boolean) => {
      if (isValid && !props.readonly) {
        submit(form.value);
      }
    });
  }
}, {timeout: 500, immediate: true});

const focus = ref(false);
watch(focus, (newVal) => {
  console.log('focus changed:', newVal);
  if (newVal) {
    pause();
    emit('focus');
  } else {
    resume();
    emit('blur');
  }
});
</script>

<template>
  <v-form @submit.prevent="submit(form)" ref="formRef" autocomplete="on">
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
              @focus="focus = true"
              @blur="focus = false"
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
              @focus="focus = true"
              @blur="focus = false"
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
            @focus="focus = true"
            @blur="focus = false"
          ></v-text-field>
          </v-col>

          <v-col cols="12" sm="4" class="my-0 py-0">
          <v-text-field
            v-model="form.housenumber"
            label="Hausnummer"
            required
            id="housenumber"
            :readonly="readonly"
            @focus="focus = true"
            @blur="focus = false"
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
              @focus="focus = true"
              @blur="focus = false"
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
              @focus="focus = true"
              @blur="focus = false"
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
              @focus="focus = true"
              @blur="focus = false"
            ></v-text-field>
          </v-col>
        </v-row>
        <v-row no-gutters >
          <v-col cols="12" sm="12" class="my-0 py-0">
            <v-text-field
              v-model="birthdate"
              label="Geburtsdatum (TT.MM.JJJJ)"
              required
              autocomplete="birthdate"
              type="date"
              :readonly="readonly"
              :rules="[datePattern.test(birthdate) || 'Bitte gib ein gültiges Geburtsdatum ein.']"
              @focus="focus = true"
              @blur="focus = false"
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
