<script setup lang="ts">
const { title,price,amount } = $defineProps<{
  title: string;
  price: number;
  amount: number;
}>();
</script>

<template>
    <div class="d-flex w-100 justify-space-between px-4 pb-1 item">
        <div class="self">
          {{amount}}x {{ title }}
        </div> 
        <div style="align-self: end; font-weight: bolder;">
          {{parseFloat(""+price).toFixed(2)}} €
        </div>
    </div>
</template>

<style scoped>
.item {
  font-size: 1.2rem;
}

@media (max-width: 800px) {
  .item {
    font-size: 1rem;
  }
}
</style>
```