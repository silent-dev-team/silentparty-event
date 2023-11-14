<script setup lang="ts">
const pb = usePocketbase()
const notifyer = useNotifyer()

const props = defineProps({
  modelValue: {
    type: Boolean,
    required: true
  }
})

const emit = defineEmits(['update:modelValue'])

let mail = $ref('')
let password = $ref('')

function login(mail:string, password:string){
  pb.admins.authWithPassword(mail.trim(), password.trim())
    .then((res) => {
      notifyer.notify('Login erfolgreich', 'success')
      console.log(res)
      emit('update:modelValue', false)
    })
    .catch((err) => {
      notifyer.notify('Login fehlgeschlagen', 'error')
      console.log(err)
    })
}

</script>

<template>
  <v-dialog width="400px" v-model="props.modelValue" :persistent="true">
    <v-card class="pa-3">
      <v-card-title>
        <h1>Login</h1>
      </v-card-title>
      <v-divider></v-divider>
      <v-card-text>
        Hier kannst du dich einloggen.
        <v-text-field
          class="mt-3"
          v-model="mail"
          label="E-Mail"
          type="email"
          required
        ></v-text-field>
        <v-text-field
          class="mt-3"
          v-model="password"
          label="Passwort"
          type="password"
          required
        ></v-text-field>
      </v-card-text>
      <v-card-actions>
        <v-btn color="primary" @click="emit('update:modelValue',false)">Abbrechen</v-btn>
        <v-spacer></v-spacer>
        <v-btn color="primary" @click="login(mail, password)">Login</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>
