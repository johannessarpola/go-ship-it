// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from 'next'

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse
) {
  const repos = await (await fetch(`http://localhost:8080/repositories`)).json()
  res.status(200).json(repos)
}
