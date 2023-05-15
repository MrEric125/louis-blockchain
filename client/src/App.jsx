import { useState } from 'react'
import reactLogo from './assets/react.svg'
import './App.css'
import { Navbar,Welcome,Loader,Services,Transactions } from './components'
const element = <h1>Hello, world!</h1>;

const App=()=> {
  return (
    <div className="min-h-screen">
      <div className='gradient-bg-welcome'>
        <Navbar>
          
        </Navbar>
        <Welcome></Welcome>

      </div>
      <Services></Services>
      <Transactions></Transactions>

    </div>
  )
}

export default App
