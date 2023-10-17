// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  //srcDir: 'app/',
  ssr: false,
  runtimeConfig: {
    public: {
      pocketbase: process.env.NUXT_POCKETBASE_URL,
    }
  },
  modules: [
    '@pinia/nuxt',
    '@vue-macros/nuxt',
    '@vite-pwa/nuxt'
  ],
  pinia: {
    autoImports: [
      // automatically imports `defineStore`
      'defineStore', // import { defineStore } from 'pinia'
      ['defineStore', 'definePiniaStore'], // import { defineStore as definePiniaStore } from 'pinia'
    ],
  },
  css: [
    'vuetify/lib/styles/main.sass',
    '@mdi/font/css/materialdesignicons.min.css',
  ],
  pwa: {
    manifest: {
      name: 'Silent App',
      lang: 'de',
      start_url: 'https://app.silentparty-hannover.de',
    }
  },
  build: {
    transpile: ['vuetify'],
  },
  vite: {
    define: {
      'process.env.DEBUG': false,
    },
  },
  imports: {
    autoImport: true,
    addons: {
      vueTemplate: true
    },
  },
})
