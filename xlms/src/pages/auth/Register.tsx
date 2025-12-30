import React, { useState } from "react";
import api from "../../api/api";
import { Toaster, toast } from "react-hot-toast";
import type { User } from "../../types/user"
import Spinner from "../../components/Spinner";
import { Link } from 'react-router-dom'

export default function Register() {
  const [form, setForm] = useState({
    full_name: "",
    email: "",
    role: 0,
  });

  const [loading, setLoading] = useState(false);

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>
  ) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    console.log(form);
    setLoading(true)
    try {
      const res = await api.post("/users/register", {
        full_name: form.full_name,
        email: form.email,
        role_id: Number(form.role)
      })
      const user: User = res.data
      console.log(user)
      toast.success(res.data.message)
    } catch (err: any) {
      toast.error(err.response?.data.message || err.message)
      console.error("Register error:", err.response?.data || err.message);
    }
    setLoading(false)
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <Toaster 
      position="top-right"
      reverseOrder={false}
      />
      <div className="bg-white p-8 rounded-lg shadow-md w-full max-w-md">
        <h2 className="text-2xl font-semibold text-center mb-6">
          LMS Registration
        </h2>

        <form onSubmit={handleSubmit} className="space-y-4">
          <input
            type="text"
            name="full_name"
            placeholder="Full Name"
            value={form.full_name}
            onChange={handleChange}
            className="w-full px-4 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
            required
          />

          <input
            type="email"
            name="email"
            placeholder="Email Address"
            value={form.email}
            onChange={handleChange}
            className="w-full px-4 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
            required
          />

          <select
            name="role"
            value={form.role}
            onChange={handleChange}
            className="w-full px-4 py-2 border rounded bg-white focus:outline-none focus:ring-2 focus:ring-blue-500"
            required
          >
            <option >Select Role</option>
            <option value="3">User</option>
            <option value="2">Manager</option>
            <option value="1">Admin</option>
          </select>

          <button
            disabled={loading}
            type="submit"
            className="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700 transition flex justify-center"
          >
             {loading ? <Spinner size="sm" color="text-white" /> : "Register"}
          </button>
          <Link to={"/login"} className="text-blue-800 my-2">Already have account</Link>
        </form>
      </div>
    </div>
  );
}
