<script setup lang="ts">
const router = useRouter()

definePageMeta({
  layout: 'admin',
});

let showLogin = $ref(false)

</script>

<template>
    <div class="container ma-3 text-center">
        <h1>Silent Party App</h1>
        <p>Hallo lieber Helfer, dies ist die Silent Party App. Hier kannst du unsere Gäste, Kassen und Kopfhörer verwalten. Mit dieser intuitiven App wird die Organisation deiner Veranstaltung zum Kinderspiel. Verfolge den Status der Kopfhörerverleih, behalte den Überblick über die Einnahmen und sorge für ein unvergessliches Erlebnis für alle Gäste.</p>
        <p>Unsere App ist darauf ausgelegt, dir maximale Kontrolle und Übersichtlichkeit zu bieten, damit deine Party ein voller Erfolg wird. Viel Spaß bei der Nutzung unserer App!</p>
    </div>
</template>

<style scoped>
        body {
            font-family: 'Arial', sans-serif;
            margin: 0;
            padding: 0;
            background-color: #f4f4f4;
            color: #333;
            text-align: center;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
        }

        .container {
            max-width: 600px;
            background-color: white;
            padding: 40px;
            border-radius: 8px;
            box-shadow: 0 4px 8px rgba(0,0,0,0.1);
        }

        h1 {

            margin-bottom: 20px;
        }

        p {
            font-size: 1.1em;
            line-height: 1.6;
            margin-bottom: 20px;
        }

        @media (max-width: 600px) {
            .container {
                padding: 20px;
            }

            p {
                font-size: 1em;
            }
        }
</style>