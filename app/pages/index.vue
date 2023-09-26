<script setup lang="ts">
const router = useRouter()

definePageMeta({
  layout: 'admin',
});

const buttons = [
  {
    text: 'Kasse',
    onClick: () => router.push({path:'/kasse'})
  },
  {
    text: 'Tickets',
    onClick: () => router.push({path:'/admin/tickets'})
  }
]

let showLogin = $ref(false)

</script>

<template>
    <div class="buttons">
      <v-btn 
        v-for="btn in buttons" :key="btn.text"
        class="mx-auto my-5" 
        width="150px" height="50px" 
        outline dark large 
        @click="btn.onClick()"
      >{{ btn.text }}
      </v-btn>
    </div>

</template>

<style scoped>
.title {
  margin: 0;
  font-size: 24px;
  color: white;
}

.buttons {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: calc(100vh - 60px);
}
</style>