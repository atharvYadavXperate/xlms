import { useEffect, useState } from "react";
import Table from "../../components/table";
import api from "../../api/api";
import "../../App.css";
import DashboardIcon from "@mui/icons-material/Dashboard";
import { Toaster, toast } from "react-hot-toast";
import Calendar, { type MonthInfo } from "../../components/Calender";

export type User = {
  created_at: Date;
  email: string;
  full_name: string;
  is_approved: boolean;
  role_id: number;
};

export default function Dashboard() {
  const [allUsers, setAllUsers] = useState<User[]>([]);
  const [users, setUsers] = useState<User[]>([]);
  const [search, setSearch] = useState("");

  const [monthInfo, setMonthInfo] = useState<MonthInfo | null>(null);
  const [isFiltered, setFiltered] = useState<boolean>(false)
  const [searchString, setSearchString] = useState<string>("")

  async function loadAllUsers() {
    let page = 1;
    const limit = 20;

    while (true) {
      const res = await api.get(
        `/users/getusers?page=${page}&limit=${limit}`
      );

      const fetchedUsers: User[] = res.data.data;

      if (fetchedUsers.length === 0) break;

      setAllUsers(prev => [...prev, ...fetchedUsers]);

      if (fetchedUsers.length < limit) break;

      page++;
    }
  }


  const searchOnChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSearch(e.target.value);
  };


  useEffect(() => {
    if (!search.trim()) {
      setUsers(allUsers);
      return;
    }

    const filtered = allUsers.filter(user =>
      user.full_name.toLowerCase().includes(search.toLowerCase()) ||
      user.email.toLowerCase().includes(search.toLowerCase())
    );

    setUsers(filtered);
  }, [search, allUsers]);


  useEffect(() => {
    toast.promise(loadAllUsers(), {
      loading: "Loading users",
      success: "Users loaded",
      error: "Failed to load users",
    });
  }, []);
  

  return (
    <div className="overflow-hidden">
      <Toaster />

      <div className="flex items-center p-3">
        <DashboardIcon color="primary" fontSize="medium" />
        <h1 className="ml-3 text-2xl font-semibold text-blue-400">
          Dashboard
        </h1>
      </div>

      <div className="flex items-center">
        <input
          onChange={searchOnChange}
          type="search"
          placeholder="Search Employee"
          className="p-2 mx-2 border-2 border-gray-400 rounded"
        />

        <Calendar setMonth={setMonthInfo} />
      </div>

      <Table month={monthInfo} users={users} />
    </div>
  );
}
