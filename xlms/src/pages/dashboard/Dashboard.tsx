import { useEffect, useState } from 'react'
import Table from '../../components/table'
import api from '../../api/api'
import "../../App.css"
import DashboardIcon from '@mui/icons-material/Dashboard';

export type User = {
  created_at: Date,
  email: string,
  full_name: string,
  is_approved: boolean,
  role_id: number
}

export default function Dashboard() {
  const [users, setUsers] = useState<User[]>([]);
  async function loadUsers() {
    const res = await api.get("/users/getusers");
    setUsers(res.data.data)
    console.log(res.data)
  }

  useEffect(() => {
    loadUsers()
  }, [])

  return (
    <div className="overflow-hidden">
      <div className='flex items-center p-3'>
        <DashboardIcon color="primary" fontSize="medium"></DashboardIcon>
        <h1 className="ml-3 text-2xl font-semibold text-blue-400">Dashboard</h1>
      </div>
      <div>
        <input type="search" placeholder='Search Employee' className="p-2 mx-2 border-2 border-gray-400 rounded"/>
      </div>
      <Table users={users} />
    </div>
  )
}
