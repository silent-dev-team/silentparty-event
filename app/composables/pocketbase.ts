import PocketBase from 'pocketbase'

const runtimeConfig = useRuntimeConfig();
const pb = new PocketBase(runtimeConfig.public.pocketbase)

export const usePocketbase = () => {
    return pb
}