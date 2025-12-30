import { BrowserRouter, Routes, Route } from "react-router-dom";
import Home from "./pages/Home";
import Login from "./pages/auth/Login";
import Register from "./pages/auth/Register";
import MainLayout from "./pages/dashboard/MainLayout";
import Dashboard from "./pages/dashboard/dashboard";
import Calender from "./pages/dashboard/Calender";
function App() {
  return (
    <BrowserRouter>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route element={<MainLayout />}>
            <Route path="/dashboard" element={<Dashboard/>}/>
            <Route path="/calender" element={<Calender/>}/>
          </Route>
        </Routes>
    </BrowserRouter>
  );
}

export default App;
