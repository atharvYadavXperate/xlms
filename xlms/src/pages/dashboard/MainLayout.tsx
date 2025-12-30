import React from 'react'
import Header from '../../components/Header';
import { Outlet } from "react-router-dom";
import Hamburger from '../../components/Hamburger';
import "../../App.css"
export default function MainLayout() {
  return (
    <div>
      <Header />
      <main className='flex flex-col md:flex-row flex-1 scroll-hide'>
        <Hamburger />
        <div className="overflow-scroll scrollbar-thin">
          <Outlet />
        </div>
      </main>
    </div>
  )
}
