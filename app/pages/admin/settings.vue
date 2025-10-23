<script lang="ts" setup>

definePageMeta({
  layout: 'admin',
});

const sonstiges = 'sonstiges'
const stationen = [
  'Eingang',
  'Ausgang',
  'KH-Ausgabe',
  sonstiges
]

const settings = useSettingsStore()
const selectedStation = ref(stationen.includes(settings.stationName) ? settings.stationName : sonstiges)
watch(selectedStation, (newVal) => {
  if (newVal === sonstiges) {
    settings.stationName = ''
  } else {
    settings.stationName = newVal
  }
})

const {notify} = useNotifyer()
const limit = ref(5)
const clickCounter = ref(0)
watch(clickCounter, (newVal) => {
  if (newVal >= limit.value) {
    settings.admin = !settings.admin
    clickCounter.value = 0
    notify(`Admin-Modus ${settings.admin ? 'aktiviert' : 'deaktiviert'}`, 'success')
  }
})
</script>

<template>
  <v-container style="max-width: 600px;">
    <v-row>
      <v-col>
        <h1>Settings</h1>
        <v-form class="w-100">
          <v-select
            v-model="selectedStation"
            label="Station Name"
            hint="Der Name der Station, an der das Gerät eingesetzt wird."
            persistent-hint
            :items="stationen"
          ></v-select>
          <v-text-field
            v-model="settings.stationName"
            v-if="selectedStation === sonstiges"
            label="Stationsname (Sonstiges)"
            hint="Die eindeutige ID der Station. Sollte für jede Station unterschiedlich sein."
            persistent-hint
          ></v-text-field>
          <v-switch color="green" v-model="settings.showIDCardPreview" label="Perso-Kontrollansicht"></v-switch>
          <v-switch color="green" v-model="settings.noInteraction" label="Confirm-Dialoge überspringen"></v-switch>
          <v-switch color="green" v-model="settings.showDashboard" label="Dashboard-Overlay anzeigen"></v-switch>
          <v-btn class="mx-auto w-100 mt-5" @click="clickCounter++" style="opacity: 10%;">Admin-Modus {{settings.admin ? 'deaktivieren' : 'aktivieren'}} {{clickCounter ? `${clickCounter}/${limit}` : ''}}</v-btn>
        </v-form>
      </v-col>
    </v-row>
  </v-container>
</template>
