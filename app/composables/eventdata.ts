import type { UnsubscribeFunc } from "pocketbase"

export function useEventData() {
  const pb = usePocketbase()
  const userStats = ref<IUserStats | null>(null)

  const refresh = async () => {
    userStats.value = await pb.userstats()
  }

  let unsubscribe: UnsubscribeFunc | null = null
  let interval: ReturnType<typeof setInterval> | null = null

  onMounted(async () => {
    await refresh()
    unsubscribe = await pb.realtime.subscribe('userstats', (e) => {
      userStats.value = e
    })
    interval = setInterval(() => {
      refresh()
    }, 10 * 1000) // every 10 seconds
  })

  onBeforeUnmount(() => {
    if (unsubscribe) {
      unsubscribe()
    }
    if (interval) {
      clearInterval(interval)
    }
  })

  function setAKContingent(value: number) {
    pb.collection('numbers').getFirstListItem('key="overbooks"').then((record) => {
      record.value = value
      pb.collection('numbers').update(record.id, record)
    })
  }

  return {
    data: readonly(userStats), refresh, setAKContingent
  }
}
