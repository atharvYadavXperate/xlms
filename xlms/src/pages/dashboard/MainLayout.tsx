import React, { useState } from "react";
import Header from "../../components/Header";
import { Outlet } from "react-router-dom";
import Hamburger from "../../components/Hamburger";
import "../../App.css";
import { Toaster, toast } from "react-hot-toast";

export default function MainLayout() {
  

  return (
    <div className="h-screen overflow-hidden">
      <Header />

      <div className="flex flex-col md:flex-row h-[calc(100vh-64px)] overflow-hidden">

        <aside
          className="
            w-full md:w-24
            md:fixed md:top-16 md:left-0
            md:h-[calc(100vh-64px)]
            border-b md:border-b-0 md:border-r
            bg-white z-40
          "
        >
          <Hamburger />
        </aside>

        <main className="flex-1 overflow-y-auto md:ml-24 scrollbar-thin">
          <Outlet />
        </main>

      </div>
    </div>
  );
}
