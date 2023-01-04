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

export default async (req: NextApiRequest, res: NextApiResponse<Response>) => {
    const { body  , headers } = req
    if (req.method === 'POST') {
        try {
            const { data } = await axios.post(
                'http://127.0.0.1:8000/authentication',
                body,
                {headers: headers}
            )
            if(data.status === true){
                res.setHeader("Set-Cookie",`token=${data.result.token}; samesite=-; path=/;`)
            }else {
                res.setHeader('Set-Cookie', 'token=deleted; path=/; expires=Thu, 01 Jan 1970 00:00:00 GMT')
            }
            return  res.send(data)
        } catch (err) {
            return  res.status(500)
        }
    } else {
        return  res.status(404)
    }
}