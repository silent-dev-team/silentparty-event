<template>
  <Login v-model="showLogin" />
  <v-app id="inspire">
    <v-app-bar height="50" scroll-behavior="elevate" class="header">
      <v-toolbar-title>Silent App</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn v-if="!pb.authStore.isAdmin" @click="showLogin = true">
        login
      </v-btn>
      <v-btn v-if="pb.authStore.isAdmin" @click="logout()">
        logout
      </v-btn>
    </v-app-bar>

    <v-main>
      <div class="bg-blue-grey-lighten-5" style="height: 100vh; width: 100%; position: fixed;">
        <slot v-if="pb.authStore.isAdmin" />
      </div>
    </v-main>

    <HelpBtn :from="page" msg="Glocke gedrückt" icon="mdi-bell" :sync="false" style="position: fixed; left: 1rem; bottom: 4rem; z-index: 5000;" />

    <v-bottom-navigation height="50" :active="navbar">
      <v-btn 
        position="absolute" 
        fab
        icon="mdi-menu-down" 
        variant="plain"
        size="large"
        @click="navbar = false"
        style="bottom: 0; left: 0; z-index: 1000;"
      ></v-btn>
      <v-btn value="KH" @click="push('/admin/hp')">
        <v-icon>mdi-headphones</v-icon>

        <span>KH</span>
      </v-btn>

      <v-btn value="tickets" @click="push('/admin/tickets/ak')">
        <v-icon>mdi-ticket-confirmation</v-icon>

        <span>AK</span>
      </v-btn>

      <v-btn value="vvk" @click="push('/admin/tickets/vvk')">
        <v-icon>mdi-ticket</v-icon>

        <span>VVK</span>
      </v-btn>

      <v-btn value="bar" @click="push('/admin/bar')">
        <v-icon>mdi-cash-register</v-icon>

        <span>Bar</span>
      </v-btn>

      <v-btn value="bar" @click="push('/admin/door')">
        <v-icon>mdi-door</v-icon>

        <span>Tür</span>
      </v-btn>
    </v-bottom-navigation>
    <v-btn 
      v-if="!navbar"
      position="fixed" 
      fab
      icon="mdi-menu-up" 
      variant="plain"
      size="large"
      @click="navbar = true"
      style="bottom: 0; left: 12px; z-index: 1000;"
    ></v-btn>
  </v-app>
</template>

<script setup>
const pb = usePocketbase()
const router = useRouter()
const shopStore = useShopStore()
shopStore.loadShop()

let page = $computed(() => router.currentRoute.value.path.split('/').at(-1))

const unsubscripe = await pb.collection('shop_items').subscribe('*', function (e) {
    console.log(e.action);
    console.log(e.record);
    shopStore.loadShop();
});

onBeforeMount(() => {
  shopStore.loadShop();
});

onUnmounted(() => {
    unsubscripe();
});

function push(path) {
  navbar = false
  router.push({ path })
}

let showLogin = $ref(false)
let navbar = $ref(false)

if (!pb.authStore.token) {
  pb.admins.authRefresh()
    .then(res => console.log(res))
    .catch(err => showLogin = true)
}
function logout() {
  pb.authStore.clear()
  window.location.reload()
}
</script>

<script>
export default {
  data: () => ({ drawer: null }),
}
</script>

<style scoped>
.frame {
  width: 100vw;
  display: flex;
  flex-direction: column;
  background: linear-gradient(to bottom, #0054bb, #00c6ff);
}

.header {
  width: 100%;
  background-color: #0054bb;
  color: aliceblue;
  display: flex;
  align-items: center;
  padding: 0 1rem;
}
</style>
