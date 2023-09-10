<script setup lang="ts">
import { CustomerData } from 'types/form';

const pb = usePocketbase()

function decodeQRString(s:string):CustomerData{
  return JSON.parse(s) //TODO: type safe
}

const data = {
    "firstName": "test",
    "lastName": "test",
    "street": "test",
    "housenumber": "test",
    "zipCode": "test",
    "place": "test",
    "email": "test@example.com",
    "sold": true,
    "validated": true,
    "used": true
};


</script>

<template>
  <h1>Hallo</h1>
  <v-btn @click="pb.collection('tickets').create(data)">create</v-btn>
</template>