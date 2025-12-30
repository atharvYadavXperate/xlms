import  { useEffect, useState } from 'react'
import Table from '../../components/table'
import api from '../../api/api'


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
    <div>
      <Table users={users}/>
    </div>
  )
}
