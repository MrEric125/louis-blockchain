import { useState } from 'react'
import reactLogo from './assets/react.svg'
import './App.css'
import { Navbar,Welcome,Loader,Services,Transactions } from './components'
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

var Web3=require('web3')



var web3=new Web3(Web3.givenProvider|| "ws://localhost:8545")

// 修改provier
web3.setProvider(new web3.providers.WebsocketProvider("ws://localhost:8545"))




export default App
