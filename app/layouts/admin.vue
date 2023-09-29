<template>
  <Login v-model="showLogin" />
  <v-app id="inspire">
    <v-app-bar scroll-behavior="elevate" class="header">
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

    <v-bottom-navigation>
      <v-btn value="KH">
        <v-icon>mdi-headphones</v-icon>

        <span>KH</span>
      </v-btn>

      <v-btn value="tickets" @click="toTickets()">
        <v-icon>mdi-ticket-confirmation</v-icon>

        <span>Tickets</span>
      </v-btn>

      <v-btn value="vvk" @click="toVVK()">
        <v-icon>mdi-ticket</v-icon>

        <span>VVK</span>
      </v-btn>

      <v-btn value="kasse">
        <v-icon>mdi-cash-register</v-icon>

        <span>Kasse</span>
      </v-btn>
    </v-bottom-navigation>
  </v-app>
</template>

<script setup>
import { ref } from 'vue'
const pb = usePocketbase()
const router = useRouter()

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

function toTickets() {
  router.push({ path: '/admin/tickets' })
}

function toVVK() {
  router.push({ path: '/admin/tickets/vvk' })
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
