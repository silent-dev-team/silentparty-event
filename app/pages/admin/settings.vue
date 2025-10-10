<script lang="ts" setup>

definePageMeta({
  layout: 'admin',
});

const settings = useSettingsStore()
const stationen = [
  'Eingang',
  'Ausgang',
  'KH-Ausgabe'
]
</script>

<template>
  <v-container style="max-width: 600px;">
    <v-row>
      <v-col>
        <h1>Settings</h1>
        <v-form>
          <v-autocomplete
            v-model="settings.stationName"
            label="Station Name"
            hint="Der Name der Station, an der das Gerät eingesetzt wird."
            persistent-hint
            :items="stationen"
          ></v-autocomplete>
          <v-switch color="green" v-model="settings.showIDCardPreview" label="Perso-Kontrollansicht"></v-switch>
          <v-switch color="green" v-model="settings.noInteraction" label="Confirm-Dialoge überspringen"></v-switch>
        </v-form>
      </v-col>
    </v-row>
  </v-container>
</template>
