import React from 'react'
import "../App.css"
import { Link } from 'react-router-dom'

export default function Hamburger() {
  return (
    <div className="w-full
        md:w-72 md:h-screen
        border-b border-gray-200
        md:border-b-0 md:border-r
        bg-white">
        <div className='p-3 hover:bg-gray-100  sm:w-full'><Link to={'/dashboard'}>Dashboard</Link></div>
        <div className='p-3 hover:bg-gray-100  sm:w-full'><Link to={'/calender'}>Calender</Link></div>
    </div>
  )
}
