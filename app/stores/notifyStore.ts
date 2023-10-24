
type Color = 'info'|'success'|'error'

export const useNotifyer = defineStore('notifyer', {
  state: () => ({
    show: false,
    message: '' as string,
    color: 'info' as Color,
    duration: 5000 as number,
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
    notify(message: string, color: Color = 'info', duration: number = 5000) {
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