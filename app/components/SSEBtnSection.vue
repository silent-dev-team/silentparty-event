<script setup lang="ts">
const settings = useSettingsStore()
</script>

<template>
    <v-dialog :model-value="!settings.stationName" persistent max-width="600">
        <v-card>
            <v-card-title class="text-h5">An welcher Station wird das Gerät eingesetzt?</v-card-title>
            <v-card-text>
                <p>Die Station wird in Benachrichtigungen und Hilfemeldungen verwendet.</p>
            </v-card-text>
            <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="primary" text @click="$router.push('/admin/settings')">Go to Settings</v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
    <HelpBtn :from="settings.stationName" msg="hodor" channel="default" icon="mdi-door" :sync="true" :global="true" style="position: fixed; left: 1rem; bottom: 12rem; z-index: 5001;"/>
    <HelpBtn :from="settings.stationName" msg="👁️Awareness👁️" channel="awareness" icon="mdi-account-eye" :sync="true" style="position: fixed; left: 1rem; bottom: 8rem; z-index: 5000;" />
    <HelpBtn :from="settings.stationName" msg="🔔Glocke🔔" channel="default" icon="mdi-bell" :sync="true" style="position: fixed; left: 1rem; bottom: 4rem; z-index: 5000;" />
</template>
