<script setup lang="ts">
definePageMeta({
  layout: 'admin',
});

const router = useRouter();

let showScanner = $ref(false);

function onScan(code: string) {
  router.push({path: `/admin/tickets/${code}`});
}

function push(p: string) {
  router.push({path: p});
}

</script>

<template>

  <v-dialog width="90%" v-model="showScanner">
    <v-card class="pa-3">
      <scanner v-if="showScanner" @onScan="onScan($event)"></scanner>
    </v-card>
  </v-dialog>
<div class="pa-3">
  <!-- <v-card class="my-3" @click="push('/admin/tickets/create')">
    <v-card-title>
      Neues Ticket
    </v-card-title>
  </v-card> -->

  <v-card class="my-3" @click="push('/admin/tickets/register')">
    <v-card-title>
      Papierticket verkaufen
    </v-card-title>
  </v-card>

  <v-card class="my-3" @click="showScanner = true">
    <v-card-title>
      Ticket kontollieren
    </v-card-title>
  </v-card>
  <!-- <div style="position: absolute; height: 100vh;" class="mr-8">
    <Scanner @on-scan="onScan($event)"/>
  </div> -->

</div>
</template>