<script setup lang="ts">
const router = useRouter()

const buttons = [
  {
    text: 'Kasse',
    onClick: () => router.push({path:'/kasse'})
  },
  {
    text: 'Ticket',
    onClick: () => router.push({path:'/ticket'})
  }
]

</script>

<template>
  <div class="frame">
    <header class="header">
      <h1 class="title">Silent Party Hannover</h1>
    </header>
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
  </div>
</template>

<style scoped>
.frame {
  width: 100%;
  height: 100vh;
  background: linear-gradient(to bottom, #0054bb, #00c6ff);
}

.header {
  display: flex;
  justify-content: flex-start;
  align-items: center;
  height: 60px;
  padding: 0 20px;
}

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