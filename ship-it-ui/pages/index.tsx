import styles from '../styles/Home.module.css'
import { useState, useEffect } from 'react'
import Repositories from '../components/repositories'

export default function Home() {
  const [data, setData] = useState([])
  const [isLoading, setLoading] = useState(false)


  useEffect(() => {
    setLoading(true)
    fetch('/api/repositories')
      .then((res) => res.json())
      .then((data) => {
        console.log(data)
        setData(data)
        setLoading(false)
      })
  }, [])

  if (isLoading) return <p>Loading...</p>
  if (!data) return <p>No repositories found</p>

  return (
    <>
      <h1>Hello world! </h1>
      <Repositories repos={data}></Repositories>
    </>
  )
}