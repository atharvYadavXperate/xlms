import React from 'react'
import { Link } from 'react-router-dom'
import "../App.css"
export default function Home() {
    return (
        <div className="flex items-center justify-center min-h-screen bg-gray-100">
            <div className="w-full max-w-md p-10 text-center bg-white rounded-lg shadow-md">

                <h1 className="mb-2 text-3xl font-bold text-gray-800">
                    Leave Management System
                </h1>

                <p className="mb-8 text-gray-600">
                    Track Your Leave
                </p>

                <div className="flex justify-center gap-4">
                    <Link
                        to="/login"
                        className="px-6 py-2 text-white transition bg-blue-600 rounded hover:bg-blue-700"
                    >
                        Login
                    </Link>

                    <Link
                        to="/register"
                        className="px-6 py-2 text-blue-600 transition border border-blue-600 rounded hover:bg-blue-50"
                    >
                        Register
                    </Link>
                </div>
            </div>
        </div>
    )
}
