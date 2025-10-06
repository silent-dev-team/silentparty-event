
export const useSettingsStore = defineStore('settings', {
  state: () => ({
    showIDCardPreview: false as boolean,
  }),
  getters: {
    idCardPreview: (state) => state.showIDCardPreview,
  },
  actions: {
    toggleIDCardPreview() {
      this.showIDCardPreview = !this.showIDCardPreview
    },
    setIDCardPreview(value: boolean) {
      this.showIDCardPreview = value
    },
  }
})
