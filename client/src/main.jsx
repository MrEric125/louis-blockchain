import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import './index.css'

const name = 'Josh Perez';
const element = <h1>Hello, {name}</h1>;

ReactDOM.createRoot(document.getElementById('root')).render(

  <React.StrictMode>
    
    <App />
  </React.StrictMode>
)
