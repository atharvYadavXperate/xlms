import { useEffect, useState } from "react";
import { type User } from "../pages/dashboard/Dashboard";
import type { MonthInfo } from "../components/Calender";

type Day = {
  month: string;
  date: number;
  day: string;
};

function DayHeaders({ days }: { days: Day[] }) {
  return (
    <>
      {days.map((day, idx) => (
        <th
          key={idx}
          className="p-2 text-xs font-medium text-center border-b text-slate-600 whitespace-nowrap"
        >
          {day.day}
        </th>
      ))}
    </>
  );
}

export default function Table({users,month,}: {users: User[];month: MonthInfo | null;}) {
  const [days, setDays] = useState<Day[]>([]);
  function generateDaysForMonth(month: MonthInfo): Day[] {
    const result: Day[] = [];
    const start = new Date(month.year, month.month, 1);
  
    for (let i = 0; i < month.days; i++) {
      const current = new Date(start);
      current.setDate(start.getDate() + i);
  
      result.push({
        day: current.toLocaleDateString("en-US", { weekday: "short" }),
        date: current.getDate(),
        month: current.toLocaleDateString("en-US", { month: "long" }),
      });
    }
  
    return result;
  }
  
  useEffect(() => {
    if (!month) return;
    setDays(generateDaysForMonth(month));
  }, [month]);

  if (!month) return null;

  return (
    <div className="w-full">
      <div className="relative w-full overflow-x-auto bg-white shadow-md scroll-hide">
        <table className="text-left border-collapse min-w-max">
          <thead className="sticky top-0 z-10">
            <tr>
              <th className="sticky left-0 z-20 p-4 text-sm font-medium bg-white border-b text-slate-600">
                Name
              </th>
              <DayHeaders days={days} />
            </tr>
          </thead>

          <tbody>
            {users.map((user, idx) => (
              <tr key={idx} className="border-b border-gray-200">
                <td className="sticky left-0 z-10 p-4 text-sm font-medium bg-white text-slate-800">
                  {user.full_name}
                </td>

                {days.map((day, i) => (
                  <td
                    key={i}
                    className={`cursor-pointer p-2 text-center text-sm whitespace-nowrap w-[60px]
                      ${
                        day.day === "Sat" || day.day === "Sun"
                          ? "bg-gray-200"
                          : ""
                      }
                      hover:bg-gray-300`}
                  >
                    {day.date}
                  </td>
                ))}
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}
