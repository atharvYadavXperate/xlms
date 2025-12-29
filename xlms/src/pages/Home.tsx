import React from 'react'
import { Link } from 'react-router-dom'
import "../App.css"
export default function Home() {
    return (
        <div className="min-h-screen flex items-center justify-center bg-gray-100">
            <div className="bg-white shadow-md rounded-lg p-10 w-full max-w-md text-center">

                <h1 className="text-3xl font-bold text-gray-800 mb-2">
                    Leave Management System
                </h1>

                <p className="text-gray-600 mb-8">
                    Track Your Leave
                </p>

                <div className="flex gap-4 justify-center">
                    <Link
                        to="/login"
                        className="px-6 py-2 bg-blue-600 text-white rounded hover:bg-blue-700 transition"
                    >
                        Login
                    </Link>

                    <Link
                        to="/register"
                        className="px-6 py-2 border border-blue-600 text-blue-600 rounded hover:bg-blue-50 transition"
                    >
                        Register
                    </Link>
                </div>
            </div>
        </div>
    )
}
