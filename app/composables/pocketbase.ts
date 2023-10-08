import PocketBase from 'pocketbase'

const runtimeConfig = useRuntimeConfig();
const apiHost = runtimeConfig.public.pocketbase as string || location.host
const pb = new PocketBase(apiHost)

export const usePocketbase = () => {
    return pb
}