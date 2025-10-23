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

  async function setOverbooks(value: number) {
    const {notify} = useNotifyer()
    try {
      const overbooksRecord = await pb.collection('numbers').getFirstListItem('key = "overbooks"');
      await pb.collection('numbers').update(overbooksRecord.id, { value });
      notify('Overbooks Eintrag aktualisiert', 'success');
    } catch (e) {
      notify('Erstelle Overbooks Eintrag', 'info');
      try {
        await pb.collection('numbers').create({ key: 'overbooks', value });
      } catch (e) {
        notify('Fehler beim Erstellen des Overbooks Eintrags', 'error');
        throw e;
      }
    }
  }

  return {
    data: readonly(userStats), refresh, setOverbooks
  }
}
