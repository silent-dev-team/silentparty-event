<script setup lang="ts">
import { CustomerData } from 'types/form';

const pb = usePocketbase()

function decodeQRString(s:string):CustomerData{
  return JSON.parse(s) //TODO: type safe
}
</script>

<template>
  <h1>Hallo Welt</h1>
</template>