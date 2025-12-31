import React, { useState } from "react";
import image from "../assets/imgs/logo.png";

export default function Header() {
  const [imgError, setImgError] = useState(false);

  if (imgError) {
    throw new Error("Logo failed to load");
  }

  return (
    <header className="w-full bg-white shadow-md">
      <div className="flex flex-col gap-1 px-4 py-3 mx-auto md:flex-row md:items-center md:justify-between">
        <div className="flex items-center gap-3">
          <img
            src={image}
            onError={() => setImgError(true)}
            alt="Xperate Logo"
            className="h-10 p-1 rounded-md"
          />
        </div>

        <h1 className="text-xl font-bold tracking-wide text-white sm:text-2xl">
          Leave Management System
        </h1>
        <div className="flex justify-end md:justify-normal">
          <div className="flex items-center justify-center font-semibold text-black transition-transform rounded-full cursor-pointer w-9 h-9 bg-amber-300 hover:scale-105">
            P
          </div>
        </div>
      </div>
    </header>
  );
}
