<script setup lang="ts">
definePageMeta({
  layout: 'admin',
});
const {data} = useEventData();
</script>

<template>
  <div class="py-3 px-2" style="max-width: 1800px; margin: auto;">
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
  grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
}
</style>
