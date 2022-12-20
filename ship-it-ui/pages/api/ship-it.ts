// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from 'next'

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse
) {
  if(req.method === 'POST') {
    console.log(req.body)
    const resp = await fetch('http://localhost:8080/ship-it', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(req.body),
    })
    res.status(resp.status).send("ok")
  } else {
    res.status(405).send("Only POST supported")
  }
}
