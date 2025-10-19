<script setup lang="ts">
import { Overlay } from '@/types/enums';
import type { RecordModel } from 'pocketbase';
import { useUtils } from '~/composables/utils';
import { useSettingsStore } from '~/stores/settingStore';

definePageMeta({
  layout: 'admin',
});

const settingsStore = useSettingsStore();

const pb = usePocketbase();
const notifyer = useNotifyer();
const router = useRouter();

// ---- Reactive Variables ----
let scannerReset = $ref(false);
let dialog = $ref(false);
let checkoutDialog = $ref(false);
let showError = $ref(false);

let ticket = $ref<TicketRecord>();
const ticketChecked = ref(false);
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

const {senetizeTicketCode} = useUtils();

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
    ticketChecked.value = false;
    await loadCustomerDataFromTicket(s);
    // notifyer.notify('Ticket akzeptiert', 'success')
  }

  if (s.startsWith('{') && s.endsWith('}') && !ticket) {
    mode = 'ak'
    ticketChecked.value = false;
    loadCustomerDataFromForm(s);
    notifyer.notify('Gast registriert', 'success')
  }

  if (re.hp.test(s) && !hp) {
    if (!ticket) {
      notifyer.notify('Bitte zuerst Ticket scannen', 'error')
      reset()
      return
    }

    try {
      hp = await getHeadPhoneFromQRString(s);
    } catch (e) {
      if (e instanceof Error) {
        if (e.message.includes('404')) {
          notifyer.notify('Kopfhörer nicht in Datenbank gefunden', 'error')
          resetScanner();
          return
        }
      }
      notifyer.notify('Fehler beim Laden des Kopfhörers', 'error')
      resetScanner();
      return
    }
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
  console.log('loading ticket from', s)
  const id = s.split('/').pop()!;
  try {
    if (id.includes("?c=")) {
      const code = id.split('?c=')[1];
      console.log('code:',code)
      const no = senetizeTicketCode(code);
      console.log('ticket number:',no)
      if (!no) throw new Error('Ticketnummer nicht gefunden');
      ticket = await pb.collection('tickets').getFirstListItem<TicketRecord>(`no = "${no}"`);
    } else {
      ticket = await pb.collection('tickets').getOne<TicketRecord>(id);
    }
    if (!ticket) {
      throw new Error('Ticket nicht gefunden');
    }
    if (ticket.used) {
      delayedReset(1000)
      notifyer.notify('Ticket wurde bereits benutzt', 'error')
      return
    }
  } catch (e) {
    resetScanner();
    notifyer.notify('Ticket nicht in Datenbank gefunden', 'error')
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
    notifyer.notify('Ausleihe fehlgeschlagen: Kein Ticket ausgewählt', 'error')
    return
  };
  ticket = await pb.collection('tickets').update<TicketRecord>(ticket!.id, { sold: true });
  if (!ticket.sold) {
    reset();
    notifyer.notify('Ausleihe fehlgeschlagen: Ticket konnte nicht als verkauft markiert werden', 'error')
  }
  checkoutDialog = false
}

async function linkTicketToHP() {
  let link: TicketHeadPhoneRecord
  try {
    link = await pb.collection('ticket_hp').create<TicketHeadPhoneRecord>({
      ticket: ticket!.id,
      hp: hp!.id,
    })
  } catch (e) {
    notifyer.notify('Kopfhörer konnte nicht mit Ticket verknüpft werden', 'error')
    resetScanner();
    return
  }
  console.log(link)
  if (link.ticket !== ticket!.id || link.hp !== hp!.id) {
    notifyer.notify('Kopfhörer konnte nicht mit Ticket verknüpft werden', 'error')
    resetScanner();
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
      reset();
      notifyer.notify('Ausleihe erfolgreich', 'success')
    })
    .catch(() => {
      reset();
      notifyer.notify('Verkauf fehlgeschlagen: Transaktion konnte nicht erstellt werden', 'error')
    })
}

function submitTicket() {
  ticketChecked.value = true;
  dialog = false;
  // checkoutDialog = true;
}

const {startLogging, stopLogging} = useKeyLogger((logs) => {
  if (logs === enter) {
    if (!ticket?.vvk) return
    if (ticketChecked.value) {
      sell()
    }
  }
}, {timeout: 500});

onMounted(() => {
  startLogging();
})

onUnmounted(() => {
  stopLogging();
  // unsubscribe()
})

</script>

<template>
  <div class="subappmarker">
    <v-icon class="mt-1">mdi-ticket</v-icon>
    <!-- <span class="ml-2" style="position: relative; top: 4px;">AK {{ unusedVvkTickets }}</span> -->
  </div>
  <v-btn 
    style="position: absolute; right: 1rem; top: 5rem;z-index: 100;"
    :color="settingsStore.showIDCardPreview ? 'green lighten-2' : 'transparent-white'"
    icon="mdi-card-account-details"
    size="large"
    @click="settingsStore.toggleIDCardPreview()"
    >
  </v-btn>
  <Scanner 
    v-touch="{
      right: () => router.push('/admin/hp'),
    }"
    class="full-screen" 
    :overlaypath="overlay" 
    @scan="onScan($event)" 
    :reset="scannerReset"
    @update:reset="scannerReset = $event" 
  />
  <v-dialog v-model="dialog" :persistent="true"  max-width="800">
    <v-card v-if="!ticket" class="pa-4">
      <CustomerForm 
        class="pt-4" 
        v-model="customerData" 
        @submit="registerCustomer" 
        @cancel="reset()"
        submitText="Bestätigen"
        cancelText="Abbrechen" 
      />
    </v-card>
    <Ticket 
      v-if="ticket" 
      :id="ticket?.id" 
      class="w-100"
      submitText="Bestätigen" 
      cancle-text="Abbrechen"
      @update="submitTicket"
      @cancel="reset()"
      @noticket="delayedReset(3000)"
      :preview="settingsStore.showIDCardPreview"
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
    :preset="price"
    :shown="checkoutDialog" 
    @paied="sell" 
    @cancled="reset()" 
  />
  <SSEBtnSection />
</template>
