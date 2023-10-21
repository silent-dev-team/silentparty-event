import PocketBase, { RecordModel } from 'pocketbase'

type Position = {
    itemId: string;
    amount: number;
};

type PocketBaseExtended = PocketBase & {
    checkout: (positions:Position[]) => Promise<RecordModel>
    unlink: (qr:string) => Promise<RecordModel>
}

const runtimeConfig = useRuntimeConfig();
const pb = new PocketBase(runtimeConfig.public.pocketbase) as PocketBaseExtended

pb.checkout = async (positions:Position[]):Promise<RecordModel> => {
    const payload = {
        positions: positions
    }
    const res = await fetch(runtimeConfig.public.pocketbase+'/api/v1/new-transaction', {
        method: 'POST',
        body: JSON.stringify(payload),
        headers: {
        'Content-Type': 'application/json',
        },
    })
    if (!res.ok) {
        throw new Error('Network response was not ok')
    }
    const data = await res.json()
    return data
}

pb.unlink = async (qr:string):Promise<RecordModel> => {
    const payload = {
        qr: qr
    }
    const res = await fetch(runtimeConfig.public.pocketbase+'/api/collections/ticket_hp/unlink', {
        method: 'POST',
        body: JSON.stringify(payload),
        headers: {
        'Content-Type': 'application/json',
        },
    })
    if (!res.ok) {
        throw new Error('Network response was not ok')
    }
    const data = await res.json()
    return data
}


export const usePocketbase = () => {
    return pb
}