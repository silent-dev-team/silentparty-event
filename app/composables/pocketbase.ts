import PocketBase from 'pocketbase'

const runtimeConfig = useRuntimeConfig();
const url = runtimeConfig.public.pocketbase as string
const pb = new PocketBase(url)

export const usePocketbase = () => {
    return pb
}