<script lang="ts" setup>
definePageMeta({
  layout: 'admin',
});

const pb = usePocketbase();
const notifyer = useNotifyer();

let scannerReset = $ref(false);

let dialog = $ref(false);
let hp = $ref<HeadPhoneRecord>();
let ticket = $ref<TicketRecord>();
let shure = $ref(false);

watch(() => dialog, (v) => {
  if (!v) resetScanner();
})

function unlink() {
  if (!hp) return;
  pb.unlink(hp.qr)
    .then(() => {
      notifyer.notify('Unlinked', 'success')
      resetScanner();
    })
    .catch(() => notifyer.notify('Unlink failed', 'error'))
}

async function onScan(s:string) {
  if (!re.hp.test(s)) return;
  hp = await pb.collection('hp').getFirstListItem<HeadPhoneRecord>(`qr = "${s}"`);
  if (!hp) {
    notifyer.notify('HP nicht gefunden', 'error')
    return
  }
  console.log(hp)
  let ticket_hp
  try {
    ticket_hp = await pb.collection('ticket_hp').getFirstListItem(`hp = "${hp.id}" && end = null`)
  } catch (e) {
    dialog = true;
    return
  }
  if (ticket_hp) {
    ticket = await pb.collection('tickets').getOne<TicketRecord>(ticket_hp.ticket)
  } else {
    hp = undefined;
    resetScanner();
  }
  dialog = true;
  resetScanner();
}

let status = $computed(() => {
  if (!hp) return 'nicht gefunden';
  if (!hp.lent && !hp.defect) return 'bereit';
  if (hp.lent && !hp.defect) return 'verliehen';
  if (!hp.lent && hp.defect) return 'defekt';
  if (hp.lent && hp.defect) return 'verliehen und defekt';
})

function resetScanner() {
  if (!scannerReset) return;
  setTimeout(() => scannerReset = true, 800);
}
</script>

<template>
    <Scanner class="full-screen" :overlaypath="Overlay.HP" @onScan="onScan($event)" :reset="scannerReset" @update:reset="scannerReset = $event"/>
    <v-dialog fullscreen v-model="dialog">
        <v-card minWidth="250px" maxWidth="420px" class="mx-auto">
          <v-card-title>
            <v-icon icon="mdi-headphones" color="black"></v-icon>
            Kopfhörer 
          </v-card-title>
          <v-card-text>
            <div class="row">
            <a style="font-size: larger;">QR-Code:</a>
            <a style="font-size: larger; font-weight: bolder;">{{ hp?.qr }}</a>
            
            </div>
            <div class="row">
                <a style="font-size: larger;">Status</a>
                <a style="font-size: larger; font-weight: bolder;">{{ status }}</a>
            </div>
            <div v-if="ticket" class="row">
                <a style="font-size: larger;">Besitzer:</a>
                <a style="font-size: larger; font-weight: bolder;">{{ ticket.firstName }} {{ ticket.lastName }}</a>
            </div>
          </v-card-text>
          <v-card-actions class="justify-space-between">
            <v-btn variant="flat" color="primary" class="mt-6" @click="dialog = false">schließen</v-btn>
            <v-btn variant="outlined" color="red" class="mt-6" v-if="hp?.lent" @click="shure = true">Zurückgeben</v-btn>

         </v-card-actions>
        </v-card>
    </v-dialog>
    <v-dialog v-model="shure">
        <v-card maxWidth="340px" class="mx-auto">
            <v-card-title style="font-size: 16px;">
              Kopfhörer Wirklich zurückgeben?
            </v-card-title>
            <v-card-actions class="justify-space-between" >
                <v-btn variant="flat" color="primary" size="small" class="mt-6" @click="shure = false">Nein, abbrechen</v-btn>
                <v-btn variant="outlined" color="red" size="small" class="mt-6" @click="shure = false;dialog=false;unlink()">Ja, Zurückgeben</v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<style scoped>

.row {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
}

</style>
