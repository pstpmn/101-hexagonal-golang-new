import axios from 'axios'
import { NextApiRequest, NextApiResponse } from 'next'


type Response = {
    msg:string ,
    status: boolean,
    result?: {
        firstName:string ,
        lastName:string ,
        // token:string,
        username :string
    },
}

export default async (req: NextApiRequest, res: NextApiResponse) => {
    res.setHeader('Set-Cookie', 'token=deleted; path=/; expires=Thu, 01 Jan 1970 00:00:00 GMT')
    return res.send(200)
}