<script setup lang="ts">
import Dashboard from '~/components/Dashboard.vue';

definePageMeta({
  layout: 'admin',
});

const settingsStore = useSettingsStore();

const overbooks = ref<number>(0);
const {setOverbooks} = useEventData();
</script>

<template>
  <div v-if="settingsStore.admin" class="py-3 px-2" style="max-width: 80%; width: 1500px; margin: auto; z-index: 50;">
    <h2 class="text-center text-2xl font-bold mb-4">Overbooks Verwaltung</h2>
    <div class="flex items-center space-x-4 mx-auto" style="width: 400px;">
      <v-text-field
        v-model="overbooks"
        label="Anzahl der Overbooks"
        type="number"
      >
        <template #append>
          <v-btn color="primary" @click="setOverbooks(overbooks)">Speichern</v-btn>
        </template>
      </v-text-field>
    </div>
  </div>
  <Dashboard />
</template>

