import styles from '../styles/Home.module.css'
import { useState, useEffect } from 'react'
import { Text } from '@mantine/core';

import Repositories from '../components/repositories'
import { Paper } from '@mantine/core'

export default function Home() {
  const [data, setData] = useState([])
  const [isLoading, setLoading] = useState(false)


  useEffect(() => {
    setLoading(true)
    fetch('/api/repositories')
      .then((res) => res.json())
      .then((data) => {
        setData(data)
        setLoading(false)
      })
  }, [])

  if (isLoading) return <p>Loading...</p>
  if (!data) return <p>No repositories found</p>

  return (
    <Paper shadow="xs" p="md">
      <Text>Ship releases with this tool</Text>
      <Repositories repos={data}></Repositories>
      
    </Paper>
  )
}