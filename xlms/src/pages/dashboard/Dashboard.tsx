import { useEffect, useState } from "react";
import Table from "../../components/table";
import api from "../../api/api";
import "../../App.css";
import DashboardIcon from "@mui/icons-material/Dashboard";
import { Toaster, toast } from "react-hot-toast";
import Calendar, { type MonthInfo } from "../../components/Calender";

export type User = {
  id: number;
  created_at: Date;
  email: string;
  full_name: string;
  is_approved: boolean;
  role_id: number;
};

export default function Dashboard() {
  const [allUsers, setAllUsers] = useState<User[]>([]);
  const [users, setUsers] = useState<User[]>([]);
  const [managerList, setManagerList] = useState<User[]>([]);
  const [search, setSearch] = useState("");
  const [monthInfo, setMonthInfo] = useState<MonthInfo | null>(null);

  async function loadAllUsers() {
    let page = 1;
    const limit = 20;

    const userMap = new Map<number, User>();
    const managerMap = new Map<number, User>();

    try {
      while (true) {
        const res = await api.get(
          `/users/getusers?page=${page}&limit=${limit}`
        );

        const fetchedUsers: User[] = res.data.data;

        if (!fetchedUsers || fetchedUsers.length === 0) break;

        fetchedUsers.forEach((user) => {
          userMap.set(user.id, user);

          if (user.role_id !== 3) {
            managerMap.set(user.id, user);
          }
        });

        setAllUsers(Array.from(userMap.values()));
        setManagerList(Array.from(managerMap.values()));

        if (fetchedUsers.length < limit) break;

        page++;
      }
    } catch (error) {
      toast.error("Failed to load users");
      throw error;
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

    const q = search.toLowerCase();

    const filtered = allUsers.filter(
      (user) =>
        user.full_name.toLowerCase().includes(q) ||
        user.email.toLowerCase().includes(q)
    );

    setUsers(filtered);
  }, [search, allUsers]);

  useEffect(() => {
    toast.promise(loadAllUsers(), {
      loading: "Loading users...",
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

      <div className="flex items-center gap-3 p-2">
        <input
          onChange={searchOnChange}
          type="search"
          placeholder="Search employee"
          className="p-2 border-2 border-gray-400 rounded"
        />

        <Calendar setMonth={setMonthInfo} />
      </div>

      <Table
        users={users}
        managerList={managerList}
        month={monthInfo}
      />
    </div>
  );
}
