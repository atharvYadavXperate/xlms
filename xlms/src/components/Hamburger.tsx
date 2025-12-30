import React from "react";
import { NavLink } from "react-router-dom";
import DashboardIcon from "@mui/icons-material/Dashboard";
import CalendarMonthIcon from "@mui/icons-material/CalendarMonth";

export default function Hamburger() {
  return (
    <div className="bg-white border-b border-gray-200 md:h-screen md:border-b-0 md:border-r">
      <NavLink
        to="/dashboard"
        className="flex items-center w-full p-4 text-center md:justify-center md:flex-col sm:flex-row hover:bg-gray-100"
      >
        <DashboardIcon
          color="primary"
          sx={{
            fontSize: {
              xs: 24,
              md: 32,
            },
          }}
        />
        <span className="mt-1 text-sm">Dashboard</span>
      </NavLink>

      <NavLink
        to="/calender"
        className="flex items-center w-full p-4 text-center md:justify-center md:flex-col sm:flex-row hover:bg-gray-100"
      >
        <CalendarMonthIcon
          color="primary"
          sx={{
            fontSize: {
              xs: 24,
              md: 32,
            },
          }}
        />
        <span className="mt-1 text-sm">Calendar</span>
      </NavLink>

    </div>
  );
}
