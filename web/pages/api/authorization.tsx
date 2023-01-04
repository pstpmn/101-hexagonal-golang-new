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
    const {cookies} = req
    try {
        const { data } = await axios.get(
            'http://127.0.0.1:8000/authorization',
            {
                headers: {
                    "Content-Type":"application/json",
                    "authorize-token":cookies['token'],
                }
            }
        )
        return  res.send(data)
    } catch (err) {
        return  res.send(err.response.data)
    }
}