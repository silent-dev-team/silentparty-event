

export const useKasseStore1 = defineStore('kasse1', {
  state: () => ({
    _count: 0,
  }),
  getters: {
    count: (state) => state._count,
  },
  actions: {
    increment() {
      this._count++
    },
  },
})

export const useKasseStore2 = defineStore('kasse2', () => {
  const count = ref(0)
  const increment = () => count.value++
  function $reset() {
    count.value = 0
  }
  return { count, increment }
})