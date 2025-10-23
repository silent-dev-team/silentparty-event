<template>
  <Login v-model="showLogin" />
  <v-app id="inspire">
    <v-app-bar height="50" scroll-behavior="elevate" class="header">
      <v-toolbar-title style="cursor: pointer;" @click="router.push('/')">Silent App</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn icon="mdi-cog" variant="text" @click="router.push('/admin/settings').then(() => settingsStore.setNavBar(false))"></v-btn>
      <v-btn v-if="!pb.authStore.isAdmin" @click="showLogin = true">
        login
      </v-btn>
      <v-btn v-if="pb.authStore.isAdmin" @click="logout()">
        logout
      </v-btn>
    </v-app-bar>

    <v-main >
      <div class="bg-blue-grey-lighten-5" style="height: 100vh; width: 100%; position: fixed;">
        <slot  v-if="pb.authStore.isAdmin"></slot>
      </div>
    </v-main>

    <v-bottom-navigation height="50" :active="settingsStore.showNavBar">
      <v-btn 
        position="absolute" 
        fab
        icon="mdi-menu-down" 
        variant="plain"
        size="large"
        @click="settingsStore.setNavBar(false)"
        style="bottom: 0; left: 0; z-index: 1000;"
      ></v-btn>

      <v-btn value="KH" @click="push('/admin/hp')">
        <v-icon>mdi-headphones-settings</v-icon>
        <span>KH-Scan</span>
      </v-btn>

      <v-btn value="tickets" @click="push('/admin/tickets/ak')">
        <v-icon>mdi-headphones-box</v-icon>
        <span>Ausgabe</span>
      </v-btn>

      <v-btn value="vvk" @click="push('/admin/tickets/vvk')">
        <v-icon>mdi-ticket</v-icon>
        <span>Eingang</span>
      </v-btn>

      <!-- <v-btn value="bar" @click="push('/admin/bar')"> -->
      <!--   <v-icon>mdi-cash-register</v-icon> -->
      <!--   <span>Bar</span> -->
      <!-- </v-btn> -->

      <v-btn value="dashboard" @click="push('/admin/dashboard')">
        <v-icon>mdi-view-dashboard</v-icon>
        <span>Dashboard</span>
      </v-btn>

      <v-btn class="d-none d-sm-flex" value="bar" @click="push('/admin/door')">
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
      @click="settingsStore.setNavBar(true)"
      style="bottom: 0; left: 12px; z-index: 1000;"
    ></v-btn>
  </v-app>
</template>

<script setup>
const pb = usePocketbase()
const router = useRouter()
const shopStore = useShopStore()
const partyHasStarted = usePartytime()
shopStore.loadShop()
const settingsStore = useSettingsStore()


let page = $computed(() => router.currentRoute.value.path.split('/').at(-1))

const unsubscripe = await pb.collection('shop_items').subscribe('*', function (e) {
    console.log(e.action);
    console.log(e.record);
    shopStore.loadShop();
});

watch($$(page), (newPage) => {
  if (newPage === 'admin' || !page) {
    settingsStore.setNavBar(true)
  }
})

onBeforeMount(() => {
  shopStore.loadShop();
  if (!page) {
    settingsStore.setNavBar(true)
  }
});

onUnmounted(() => {
    unsubscripe();
});

function push(path) {
  // settingsStore.setNavBar(false)
  router.push({ path })
}

let showLogin = $ref(false)

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
  background: rgb(12,84,171);
  background: linear-gradient(0deg, rgba(12,84,171,1) 0%, rgba(0,60,131,1) 100%);
  color: aliceblue;
  display: flex;
  align-items: center;
  padding: 0 1rem;
}
</style>
