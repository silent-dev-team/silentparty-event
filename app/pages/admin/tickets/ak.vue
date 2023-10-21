<script setup lang="ts">
import { Overlay } from '@/types/enums';
import { RecordModel } from 'pocketbase';

definePageMeta({
  layout: 'admin',
});

const pb = usePocketbase();

let akItem: RecordModel
try {
  akItem = await pb.collection('shop_items').getFirstListItem<ShopItemRecord>('title = "AK Ticket"')
} catch (e) {
  alert('AK Ticket nicht in Datenbank gefunden')
  throw e
}

let pfandItem: RecordModel
try {
  pfandItem = await pb.collection('shop_items').getFirstListItem<ShopItemRecord>('title = "HP Pfand"')
} catch (e) {
  alert('Pfand nicht in Datenbank gefunden')
  throw e
}

let realPfandPrice = $computed(() => -pfandItem.price)

let scannerReset = $ref(false);

let dialog = $ref(false);
let cashCalc = $ref(false);
let showError = $ref(false);

let ticket = $ref<TicketRecord>();
let hp = $ref<HeadPhoneRecord>()
let price = $ref(0.0);

let overlay = $computed<Overlay>(() => 
  ticket ? Overlay.HP : Overlay.Ticket
);

let customerData =$ref<ICustomerData>(initCustomerData());

async function registerCustomer(){
  createTicket();
  dialog = false;
  cashCalc = true;
}

async function onScan(s:string) {
  if (dialog || cashCalc || showError) return;

  if (re.url.test(s) && !ticket) {
    loadCustomerDataFromTicket(s);
  }

  if (s.startsWith('{') && s.endsWith('}') && !ticket) {
    loadCustomerDataFromForm(s);
  }

  //TODO: dialog abbrechen können

  if (re.hp.test(s) && !hp) {
    if (!ticket) {
      alert('Bitte zuerst Ticket scannen')
      reset()
      return
    }

    hp = await getHeadPhoneFromQRString(s);
    if(!hp){
      alert('Kopfhörer nicht in Datenbank gefunden')
      resetScanner();
      return
    }
    if (hp.lent) {
      alert('Kopfhörer ist bereits verliehen')
      hp = undefined;
      resetScanner();
      return
    }
    alert(`Kopfhörer erfolgreich gescannt ${(hp as HeadPhoneRecord).qr}`)
    await linkTicketToHP();
    reset();
  }
}

//side effects: customerData, dialog, showError
function loadCustomerDataFromForm(s:string){
  try{
    customerData = JSON.parse(s);
    dialog = true;
    console.log(customerData)
  }catch(e){
    if(!showError){
      showError = true;
      setTimeout(() => {
        showError = false;
        resetScanner();
      }, 3000);
    }
  }
}

//side effects: ticket, customerData, dialog, showError
async function loadCustomerDataFromTicket(s:string){
  const id = s.split('/').pop()!;
  ticket = await pb.collection('tickets').getOne<TicketRecord>(id);
  if (!ticket) {
    alert('Ticket nicht in Datenbank gefunden')
    resetScanner();
    return
  }
  if (ticket.used) {
    alert('Ticket wurde bereits benutzt')
    reset()
    return
  }
  dialog = true;
}

//side effects: hp
async function getHeadPhoneFromQRString(s:string){
  return pb.collection('hp').getFirstListItem<HeadPhoneRecord>(`qr = "${s}"`);
}

async function createTicket(){
  let data = $$(customerData).value as any;
  data.filled = true;
  const record = await pb.collection('tickets').create<TicketRecord>(data);
  console.log(record)
  ticket = record;
}

async function sell(){
  if (!ticket) {
    reset();
    alert('Verkauf fehlgeschlagen: Kein Ticket ausgewählt')
    return
  };
  ticket = await pb.collection('tickets').update<TicketRecord>(ticket!.id, {sold: true});
  if (!ticket.sold) {
    reset();
    alert('Verkauf fehlgeschlagen: Ticket konnte nicht als verkauft markiert werden')
    return
  }
  pb.checkout([{
    amount: 1,
    itemId: akItem.id,
  }]).then(() => {
    cashCalc = false
  }).catch(() => {
    reset();
    alert('Verkauf fehlgeschlagen: Transaktion konnte nicht erstellt werden')
  })
}

async function linkTicketToHP() {
  const link = await pb.collection('ticket_hp').create<TicketHeadPhoneRecord>({
    ticket: ticket!.id,
    hp: hp!.id,
  });
  console.log(link)
  if (link.ticket !== ticket!.id || link.hp !== hp!.id) {
    alert('Kopfhörer konnte nicht mit Ticket verknüpft werden')
    resetScanner();
    return
  }
  reset()
}

function resetScanner(){
  // ...
  if (scannerReset) return;
  scannerReset = true;
  setTimeout(() => scannerReset = false, 100);
}

function reset(){
  ticket = undefined;
  hp  = undefined;
  customerData = initCustomerData();
  dialog = false;
  cashCalc = false;
  showError = false;
  price = 0.0;
  resetScanner();
}

</script>

<template>
  <Scanner class="full-screen" :overlaypath="overlay" @onScan="onScan($event)" :reset="scannerReset"/>
  <v-dialog v-model="dialog" :persistent="true">
    <v-card v-if="!ticket" class="pa-4">
      <CustomerForm v-model="customerData" @submit="registerCustomer" submitText="Bestätigen"></CustomerForm>
    </v-card>
    <Ticket v-if="ticket" :id="ticket?.id" submitText="Bestätigen" @update="dialog = false"></Ticket>
  </v-dialog>
  <v-dialog v-model="showError" :persistent="true">
    <v-card 
      class="text-red-darken-4 text-center" 
      color="red-lighten-3" 
      text="Der Code ist nicht Valide" 
    />
  </v-dialog>
  <CashCalculator :requested="akItem.price + realPfandPrice" :shown="cashCalc" @paied="sell" @cancled="reset()"/>
</template>