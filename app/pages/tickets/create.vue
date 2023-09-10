<script setup lang="ts">
import { RecordModel } from 'pocketbase';

const pb = usePocketbase()

let form = $ref<CustomerData>({
  firstName: '',
  lastName: '',
  street: '',
  housenumber: '',
  zipCode: '',
  place: '',
  email: '',
});

function loadDataFromQRString(s:string){
  form = JSON.parse(s)
}

async function createTicket(){
  const rec = await pb.collection('tickets').create<RecordModel & Ticket>(form)
  console.log(rec)
}

</script>

<template>
  <customer-form v-model="form" @submit="createTicket" submitText="Ticket erstelen"/>
</template>~/types/general~/types~/types