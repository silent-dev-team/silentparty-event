
export const useSettingsStore = defineStore('settings', {
  state: () => ({
    showIDCardPreview: false as boolean,
    noInteraction: false as boolean,
    stationName: '' as string,
  }),
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
    setStationName(name: string) {
      this.stationName = name
    },
  },
  persist: true,
})
