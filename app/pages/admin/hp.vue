<script lang="ts" setup>
import type { RecordModel } from 'pocketbase';

definePageMeta({
  layout: 'admin',
});

const pb = usePocketbase();
const notifyer = useNotifyer();
const router = useRouter();

let pfandItem: RecordModel
try {
  pfandItem = await pb.collection('shop_items').getFirstListItem<ShopItemRecord>('title = "Kopfhoerer"')
} catch (e) {
  notifyer.notify('Pfand nicht in Datenbank gefunden', 'error')
  throw e
}

let scannerReset = $ref(false);

let dialog = $ref(false);
let hp = $ref<HeadPhoneRecord>();
let ticket = $ref<TicketRecord>();
let shure = $ref(false);
let label = $ref(false);
let newLabel = $ref('');

watch(() => dialog, (v) => {
  if (!v) reset();
})

function unlink() {
  if (!hp) return;
  pb.unlink(hp.qr)
    .then(() => {
      notifyer.notify('Unlinked', 'success')
      pb.checkout([{amount: -1, itemId: pfandItem.id}])
      reset();
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
  newLabel = hp.label;
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

async function changeLabel(){
  if (!hp) return;
  hp.label = newLabel;
  try{
    hp = await pb.collection('hp').update<HeadPhoneRecord>(hp.id, hp);
    notifyer.notify('Label geändert', 'success')
  } catch (e) {
    notifyer.notify('Label ändern fehlgeschlagen', 'error')
    console.log(e)
  }
  label = false;
}

async function toggleDefect() {
  if (!hp) return;
  hp.defect = !hp.defect;
  try {
    hp =  await pb.collection('hp').update<HeadPhoneRecord>(hp.id, hp)
    if (hp.defect) {
      notifyer.notify('Defekt gesetzt', 'success')
    } else {
      notifyer.notify('wieder heile', 'success')
    }
  } catch (e) {
    notifyer.notify('Defekt setzen fehlgeschlagen', 'error')
    console.log(e)
  }
}

let status = $computed(() => {
  if (!hp) return 'nicht gefunden';
  if (!hp.lent && !hp.defect) return 'bereit';
  if (hp.lent && !hp.defect) return 'verliehen';
  if (!hp.lent && hp.defect) return 'defekt';
  if (hp.lent && hp.defect) return 'verliehen und defekt';
})

function reset() {
  hp = undefined;
  ticket = undefined;
  dialog = false;
  shure = false;
  label = false;
  newLabel = '';
  resetScanner();
}

function resetScanner() {
  if (scannerReset) return;
  setTimeout(() => scannerReset = true, 300);
}
</script>

<template>
    <div class="subappmarker">
      <v-icon class="mt-1">mdi-headphones</v-icon>
      <span class="ml-2" style="position: relative; top: 4px;">HP</span>
    </div>
    <Scanner     
      v-touch="{
        left: () => router.push('/admin/tickets/ak'),
      }"
      class="full-screen" 
      :overlaypath="Overlay.HP" 
      @scan="onScan($event)" 
      :reset="scannerReset" 
      @update:reset="scannerReset = $event"
    />
    <v-dialog fullscreen v-model="dialog" :persistent="true">
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
            <div class="row">
                <a style="font-size: larger; cursor: pointer;" @click="label = true">Label</a>
                <a style="font-size: larger; font-weight: bolder; cursor: pointer;" @click="label = true">{{ hp?.label }}</a>
            </div>
            <div v-if="ticket" class="row">
                <a style="font-size: larger;">Besitzer:</a>
                <a style="font-size: larger; font-weight: bolder;">{{ ticket.firstName }} {{ ticket.lastName }}</a>
            </div>
          </v-card-text>
          <v-card-actions class="justify-space-between">
            <v-btn variant="flat" color="primary" class="mt-6" @click="dialog = false">schließen</v-btn>
            <v-btn variant="outlined" color="red" class="mt-6" v-if="hp?.lent" @click="shure = true">Zurückgeben</v-btn>
            <v-btn variant="flat" color="red"     class="mt-6" :disabled="!hp" v-if="!hp?.lent && !hp?.defect" @click="toggleDefect()">Defekt</v-btn>
            <v-btn variant="flat" color="success" class="mt-6" :disabled="!hp" v-if="!hp?.lent && hp?.defect" @click="toggleDefect()">Heile</v-btn>
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
    <v-dialog v-model="label">
        <v-card maxWidth="340px" class="mx-auto" >
            <v-card-title style="font-size: 16px;">
              Kopfhörer labeln
            </v-card-title>
            <v-card-text>
                <v-text-field v-model="newLabel" style="width: 300px;" label="Label" variant="outlined"></v-text-field>
            </v-card-text>
            <v-card-actions class="justify-end" >
                <v-btn variant="flat" color="success" size="small" class="mt-6" @click="changeLabel()">Speichern</v-btn>
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
