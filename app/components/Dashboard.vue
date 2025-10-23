<script setup lang="ts">
const {data} = useEventData();
const settingsStore = useSettingsStore();
</script>

<template>
  <div v-if="settingsStore.admin" class="py-3 px-2" style="max-width: 80%; width: 1500px; margin: auto; z-index: 50;">
    <div class="stat-row">
      <Stat label="Kopfhörer im Umlauf" :value="data?.lent_hp ?? NaN" size="xl" />
      <Stat label="Kopfhörer verfügbar" :value="data?.unused_hps ?? NaN" size="xl" />
    </div>
    <div class="stat-row">
      <Stat label="Nicht benutzte VVK" :value="data?.unused_vvk ?? NaN" size="xl" />
      <Stat label="Benutzte VVK" :value="data?.used_vvk ?? NaN" size="xl" />
    </div>
    <div class="stat-row">
      <Stat label="AK Kontingent" :value="data?.open_ak ?? NaN" size="xl" />
      <Stat label="Benutzte Abendkassen Tickets" :value="data?.used_ak ?? NaN" size="xl" />
    </div>
    <div class="stat-row">
      <Stat label="Ausgegebene Abendkassen Tickets (nicht eingelöst)" :value="data?.active_ak ?? NaN" size="xl" />
      <Stat label="Abendkassen Ticktes in Reserve" :value="data?.reserved_ak ?? NaN" size="xl" />
    </div>
  </div>
  <div v-else class="py-3 px-2" style="max-width: 80%; width: 1500px; margin: auto; z-index: 50;">
    <div class="stat-row">
      <Stat label="Abendkassenkontingent" :value="data?.open_ak ?? NaN" size="xl" />
      <Stat label="Gesamt Kopfhörer im Umlauf" :value="data?.lent_hp ?? NaN" size="xl" />
      <Stat label="Ausgegebene Abendkassen Tickets (nicht eingelöst)" :value="data?.active_ak ?? NaN" size="xl" />
    </div>
  </div>

</template>

<style scoped>
.stat-row {
  display: grid;
  gap: 1rem;
  grid-template-columns: repeat(auto-fit, minmax(448px, 1fr));
  margin-bottom: 1rem;
}
</style>
