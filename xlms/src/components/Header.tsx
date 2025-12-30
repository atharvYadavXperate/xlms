import React, { useState } from "react";
import image from "../assets/imgs/logo.png";

export default function Header() {
  const [imgError, setImgError] = useState(false);

  if (imgError) {
    throw new Error("Logo failed to load");
  }

  return (
    <header className="w-full bg-blue-400 shadow-md">
      <div className="
        mx-auto 
        px-4 py-3
        flex flex-col gap-1
        md:flex-row md:items-center md:justify-between
      ">
        <div className="flex items-center gap-3">
          <img
            src={image}
            onError={() => setImgError(true)}
            alt="Xperate Logo"
            className="w-9 h-9 rounded-md bg-white p-1"
          />
          <span className="text-white font-bold text-xl sm:text-2xl tracking-wide">
            Xperate
          </span>
        </div>

        <h1 className="
          text-white font-semibold text-center
          text-base sm:text-lg md:text-xl
          tracking-wide
        ">
          Leave Management System
        </h1>
        <div className="flex justify-end md:justify-normal">
          <div className="
            w-9 h-9 rounded-full
            bg-amber-300 text-black font-semibold
            flex items-center justify-center
            cursor-pointer
            hover:scale-105 transition-transform
          ">
            P
          </div>
        </div>
      </div>
    </header>
  );
}
