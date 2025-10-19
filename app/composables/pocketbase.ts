import PocketBase, { type RecordModel } from 'pocketbase'

export type Position = {
    itemId: string;
    amount: number;
};

type PocketBaseExtended = PocketBase & {
    checkout: (positions:Position[]) => Promise<RecordModel>
    unlink: (qr:string) => Promise<RecordModel>
    fileUrl: (collectionIdOrName:string, recordId: string, filename: string|null|undefined) => string  
    fileUrlFromRecord: (record:RecordModel, filename: string) => string
    userstats: () => Promise<IUserStats>
}

const runtimeConfig = useRuntimeConfig();
console.log("PocketBase URL:", runtimeConfig.public.pocketbase);
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

pb.fileUrl = (collectionIdOrName:string, recordId: string, filename: string|null|undefined) :string => {
    if (!filename) {
        return ""
    }
    return runtimeConfig.public.pocketbase+`/api/files/${collectionIdOrName}/${recordId}/${filename}`
}

pb.fileUrlFromRecord = (record:RecordModel, filename: string) :string => {
    return runtimeConfig.public.pocketbase+`/api/files/${record.collectionId}/${record.collectionId}/${filename}`
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

pb.userstats = async ():Promise<IUserStats> => {
    const res = await fetch(runtimeConfig.public.pocketbase+'/api/v1/userstats')
    if (!res.ok) {
        throw new Error('Network response was not ok')
    }
    const data = await res.json() as IUserStats
    return data
}


export const usePocketbase = () => {
    return pb
}
