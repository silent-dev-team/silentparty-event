
export const useNotifyer = defineStore('notifyer', {
  state: () => ({
    show: false,
    message: '',
    color: 'info',
    duration: 5000,
  }),
  getters: {
    notification: (state) => {
      return {
        message: state.message, 
        color: state.color,
        duration: state.duration,
      }
    },
  },
  actions: {
    notify(message: string, color: 'info'|'success'|'error' = 'info', duration: number = 5000) {
      this.message = message
      this.color = color
      this.duration = duration
      this.show = true
      setTimeout(() => {
        this.$reset()
      }, this.duration + 250)
    }
  },
})