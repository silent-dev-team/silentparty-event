
export const useSettingsStore = defineStore('settings', {
  state: () => ({
    showIDCardPreview: false as boolean,
    noInteraction: false as boolean,
  }),
  getters: {
    idCardPreview: (state) => state.showIDCardPreview,
    interactionDisabled: (state) => state.noInteraction,
  },
  actions: {
    toggleIDCardPreview() {
      this.showIDCardPreview = !this.showIDCardPreview
    },
    setIDCardPreview(value: boolean) {
      this.showIDCardPreview = value
    },
    toggleNoInteraction() {
      this.noInteraction = !this.noInteraction
    },
    setNoInteraction(value: boolean) {
      this.noInteraction = value
    },
  },
  persist: true,
})
