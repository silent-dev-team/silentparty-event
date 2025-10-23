
export const useSettingsStore = defineStore('settings', {
  state: () => ({
    showIDCardPreview: false as boolean,
    noInteraction: false as boolean,
    stationName: '' as string,
    showDashboard: false as boolean,
    showNavBar: true as boolean,
    admin: false as boolean,
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
    toggleDashboard() {
      this.showDashboard = !this.showDashboard
    },
    setDashboard(value: boolean) {
      this.showDashboard = value
    },
    toggleNavBar() {
      this.showNavBar = !this.showNavBar
    },
    setNavBar(value: boolean) {
      this.showNavBar = value
    },
  },
  persist: true,
})
