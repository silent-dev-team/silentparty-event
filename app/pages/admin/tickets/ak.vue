<script setup lang="ts">
import { Overlay } from '@/types/enums';
import type { RecordModel } from 'pocketbase';

definePageMeta({
  layout: 'admin',
});

const pb = usePocketbase();
const notifyer = useNotifyer();

// ---- Reactive Variables ----
let scannerReset = $ref(false);
let dialog = $ref(false);
let checkoutDialog = $ref(false);
let showError = $ref(false);

let ticket = $ref<TicketRecord>();
let hp = $ref<HeadPhoneRecord>()
let mode = $ref<'ak' | 'vvk' | undefined>();
let price = $computed(() => {
  if (!mode) return 0;
  return (mode === 'ak' ? akItem.price : 0) + pfandItem.price;
});
let overlay = $computed<Overlay>(() => ticket ? Overlay.HP : Overlay.Ticket);
let customerData = $ref<ICustomerData>(initCustomerData());


// ---- Data Fetching ----
let akItem: RecordModel
let pfandItem: RecordModel

try {
  akItem = await pb.collection('shop_items').getFirstListItem<ShopItemRecord>('title = "AK Ticket"')
} catch (e) {
  notifyer.notify('AK Ticket nicht in Datenbank gefunden', 'error')
  throw e
}
try {
  pfandItem = await pb.collection('shop_items').getFirstListItem<ShopItemRecord>('title = "Kopfhoerer"')
} catch (e) {
  notifyer.notify('Pfand nicht in Datenbank gefunden', 'error')
  throw e
}


// ---- Utility Functions ----
function resetScanner() {
  if (scannerReset) return;
  setTimeout(() => {
    scannerReset = true;
  }, 800);
}

watch(() => scannerReset, () => {
  console.log('scanner reset', scannerReset)
});

function reset() {
  ticket = undefined;
  hp = undefined;
  customerData = initCustomerData();
  dialog = false;
  checkoutDialog = false;
  showError = false;
  resetScanner();
}

async function delayedReset(t: number = 1000) {
  await new Promise(resolve => setTimeout(resolve, t));
  reset();
}

// ---- Main Functions ----

//side effects: ticket, customerData, dialog, showError
async function onScan(s: string) {
  if (dialog || checkoutDialog || showError) return;
  if (re.url.test(s) && !ticket) {
    mode = 'vvk'
    loadCustomerDataFromTicket(s);
    notifyer.notify('Ticket akzeptiert', 'success')
  }

  if (s.startsWith('{') && s.endsWith('}') && !ticket) {
    mode = 'ak'
    loadCustomerDataFromForm(s);
    notifyer.notify('Gast registriert', 'success')
  }

  if (re.hp.test(s) && !hp) {
    if (!ticket) {
      notifyer.notify('Bitte zuerst Ticket scannen', 'error')
      reset()
      return
    }

    hp = await getHeadPhoneFromQRString(s);
    if (!hp) {
      notifyer.notify('Kopfhörer nicht in Datenbank gefunden', 'error')
      resetScanner();
      return
    }
    if (hp.lent) {
      notifyer.notify('Kopfhörer ist bereits verliehen', 'error')
      resetScanner();
      hp = undefined;
      return
    }
    notifyer.notify(`Kopfhörer erfolgreich gescannt ${(hp as HeadPhoneRecord).qr}`, 'success')
    await linkTicketToHP();
    delayedReset(1000)
  }
}

async function registerCustomer() {
  createTicket();
  dialog = false;
  checkoutDialog = true;
}

//side effects: customerData, dialog, showError
function loadCustomerDataFromForm(s: string) {
  try {
    customerData = JSON.parse(s);
    dialog = true;
    console.log(customerData)
  } catch (e) {
    if (!showError) {
      showError = true;
      setTimeout(() => {
        showError = false;
        resetScanner();
      }, 3000);
    }
  }
}

//side effects: ticket, customerData, dialog, showError
async function loadCustomerDataFromTicket(s: string) {
  const id = s.split('/').pop()!;
  ticket = await pb.collection('tickets').getOne<TicketRecord>(id);
  if (!ticket) {
    notifyer.notify('Ticket nicht in Datenbank gefunden', 'error')
    resetScanner();
    return
  }
  if (ticket.used) {
    notifyer.notify('Ticket wurde bereits benutzt', 'error')
    reset();
    return
  }
  dialog = true;
}

//side effects: hp
async function getHeadPhoneFromQRString(s: string) {
  return pb.collection('hp').getFirstListItem<HeadPhoneRecord>(`qr = "${s}"`);
}

async function createTicket() {
  let data = $$(customerData).value as any;
  data.filled = true;
  const record = await pb.collection('tickets').create<TicketRecord>(data);
  console.log(record)
  ticket = record;
}

async function sell() {
  if (!ticket) {
    reset();
    notifyer.notify('Verkauf fehlgeschlagen: Kein Ticket ausgewählt', 'error')
    return
  };
  ticket = await pb.collection('tickets').update<TicketRecord>(ticket!.id, { sold: true });
  if (!ticket.sold) {
    reset();
    notifyer.notify('Verkauf fehlgeschlagen: Ticket konnte nicht als verkauft markiert werden', 'error')
    return
  }
  const payload = [{
    amount: 1,
    itemId: pfandItem.id,
  }]
  if (mode === 'ak') {
    payload.push({
      amount: 1,
      itemId: akItem.id,
    })
  }
  pb.checkout(payload)
    .then(() => {
      checkoutDialog = false
    }).catch(() => {
      reset();
      notifyer.notify('Verkauf fehlgeschlagen: Transaktion konnte nicht erstellt werden', 'error')
    })
}

async function linkTicketToHP() {
  const link = await pb.collection('ticket_hp').create<TicketHeadPhoneRecord>({
    ticket: ticket!.id,
    hp: hp!.id,
  });
  console.log(link)
  if (link.ticket !== ticket!.id || link.hp !== hp!.id) {
    notifyer.notify('Kopfhörer konnte nicht mit Ticket verknüpft werden', 'error')
    resetScanner();
    return
  }
}



</script>

<template>
  <Scanner 
    class="full-screen" 
    :overlaypath="overlay" 
    @onScan="onScan($event)" 
    :reset="scannerReset"
    @update:reset="scannerReset = $event" 
  />
    <v-dialog v-model="dialog" :persistent="true">
    <v-card v-if="!ticket" class="pa-4">
      <CustomerForm 
        class="pt-4" 
        v-model="customerData" 
        @submit="registerCustomer" 
        submitText="Bestätigen"
        cancelText="Abbrechen" 
      />
    </v-card>
    <Ticket 
      v-if="ticket?.vvk" 
      :id="ticket?.id" 
      submitText="Bestätigen" 
      @update="checkoutDialog = true; dialog = false;"
      @noticket="delayedReset(3000)"
    />
  </v-dialog>
  <v-dialog v-model="showError" :persistent="true">
    <v-card 
      class="text-red-darken-4 text-center" 
      color="red-lighten-3" 
      text="Der Code ist nicht Valide" 
    />
  </v-dialog>
  <Checkout
    :requested="price" 
    :shown="checkoutDialog" 
    @paied="sell" 
    @cancled="reset()" 
  />
  <HelpBtn from="ak" msg="hodor" icon="mdi-door" :sync="true" style="position: fixed; left: 1rem; bottom: 8rem; z-index: 5001;"/>
</template>