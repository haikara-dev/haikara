import type { NextPage } from 'next'
import Head from 'next/head'

const Home: NextPage = () => {
  return (
    <div>
      <Head>
        <title>dailyfj</title>
        <meta name="description" content="dailyfj" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main>
        <h1>dailyfj</h1>
      </main>

      <footer>
        <a
          href="https://cubdesign.com"
          target="_blank"
          rel="noopener noreferrer"
        >
          ©︎cubdesign
        </a>
      </footer>
    </div>
  )
}

export default Home
