// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from 'next'
import type Repository from '../../models/Repository'

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<Repository[]>
) {
  const repos = await (await fetch(`http://localhost:8080/repositories`)).json()
  res.status(200).json(repos)
}
