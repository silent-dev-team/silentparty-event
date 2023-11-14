
// write a function that checks if the party has started
// if the party has started, return true

const runtimeConfig = useRuntimeConfig();

export function usePartytime() {
    const partytime = new Date(runtimeConfig.public.partytime)
    const partyHasStarted = () => {
        if (partytime < new Date()) return true;
        return false;
    }
    return partyHasStarted
}