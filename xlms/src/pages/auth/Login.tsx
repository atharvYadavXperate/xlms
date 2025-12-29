import React, { useState } from "react";
import Spinner from "../../components/Spinner";
import api from "../../api/api";
import { Toaster, toast } from "react-hot-toast";
export default function Login() {
  const [isRequestOtp, setIsRequestOtp] = useState(true);
  const [email, setEmail] = useState("");
  const [otp, setOtp] = useState("");
  const [isLoading, setLoading] = useState(false);

  function handleEmailChange(e: React.ChangeEvent<HTMLInputElement>) {
    setEmail(e.target.value);
  }

  function handleOtpChange(e: React.ChangeEvent<HTMLInputElement>) {
    setOtp(e.target.value);
  }

  function changeEmail() {
    setIsRequestOtp(true);
    setOtp("");
  }

  async function handleGetOtp(e: React.FormEvent) {
    e.preventDefault();
    setLoading(true);
    try {
      const res = await api.post("/users/otp", {
        email: email.trim(),
      });
      console.log(res.data);
      setIsRequestOtp(false);
      toast.success("OTP generated successfully")
    } catch (err: any) {
      console.error(err);
      toast.error(err.response?.data.message || err.message)
    } finally {
      setLoading(false);
    }
  }

  async function handleVerifyOtp(e: React.FormEvent) {
    e.preventDefault();
    setLoading(true);

    try {
      const res = await api.post("/users/login", {
        email: email.trim(),
        otp: parseInt(otp, 10), 
      });
      console.log(res.data); 
      toast.success(res.data.message)
    } catch (err: any) {
      console.error(err);
      toast.error(err.response?.data.message || err.message)
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <Toaster />
      <div className="bg-white p-8 rounded-lg shadow-md w-full max-w-md">
        <h2 className="text-2xl font-semibold text-center mb-6">
          LMS Login
        </h2>

        {isRequestOtp ? (
          <form onSubmit={handleGetOtp} className="space-y-4">
            <input
              type="email"
              placeholder="Email Address"
              value={email}
              onChange={handleEmailChange}
              className="w-full px-4 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
              required
            />

            <button
              type="submit"
              disabled={isLoading}
              className="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700 transition flex justify-center disabled:opacity-60"
            >
              {isLoading ? <Spinner size="sm" color="text-white" /> : "Get OTP"}
            </button>
          </form>
        ) : (
          <form onSubmit={handleVerifyOtp} className="space-y-4">
            <input
              type="number"
              placeholder="Enter OTP"
              value={otp}
              onChange={handleOtpChange}
              className="w-full px-4 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
              required
            />

            <button
              type="submit"
              disabled={isLoading}
              className="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700 transition flex justify-center disabled:opacity-60"
            >
              {isLoading ? <Spinner size="sm" color="text-white" /> : "Verify OTP"}
            </button>
          </form>
        )}

        {!isRequestOtp && (
          <p className="text-center text-sm text-gray-600 mt-6">
            OTP sent to <b>{email}</b>
            <button
              type="button"
              onClick={changeEmail}
              className="ml-2 text-blue-600 hover:underline"
            >
              change
            </button>
          </p>
        )}
      </div>
    </div>
  );
}
