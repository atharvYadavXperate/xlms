import React, { useEffect, useState } from "react";
import  {type User} from "../pages/dashboard/Dashboard"

type Day = {
    month: string;
    date: number;
    day: string;
};



function DayHeaders({ days }: { days: Day[] }) {
    return (
        <>
            {days.map((ele, idx) => (
                <th
                    key={idx}
                    className="p-2 border-b text-xs font-medium text-slate-600 text-center whitespace-nowrap "
                >
                    {ele.day}
                </th>
            ))}
        </>
    );
}

export default function Table({ users }: { users: User[] }) {
    const [days, setDays] = useState<Day[]>([]);
    
    function generateNext30Days(): Day[] {
        const result: Day[] = [];
        const today = new Date();

        for (let i = 0; i < 30; i++) {
            const current = new Date(today);
            current.setDate(today.getDate() + i);

            result.push({
                day: current.toLocaleDateString("en-US", { weekday: "short" }),
                date: current.getDate(),
                month: current.toLocaleDateString("en-US", { month: "long" }),
            });
        }

        return result;
    }

   
    useEffect(() => {
        setDays(generateNext30Days());
        
    }, []);

    return (
        <div className="w-full">
            <div className="relative w-full overflow-x-auto bg-white shadow-md scroll-hide">
                <table className="min-w-max border-collapse text-left">
                    <thead className="sticky top-0 z-10">
                        <tr>
                            <th className="p-4 border-b text-sm font-medium text-slate-600 sticky left-0 z-20 bg-white ">
                                Name
                            </th>
                            <DayHeaders days={days} />
                        </tr>
                    </thead>
                    <tbody>
                        {users.map((user, idx) => (
                            <tr
                                key={idx}
                                className="border-b border-gray-200"
                            >
                                <td className="p-4 text-sm font-medium text-slate-800 sticky left-0 bg-white z-10">
                                    {user.full_name}
                                </td>

                                {days.map((day, i) => (
                                    <td
                                        key={i}
                                        className={`cursor-pointer p-2 text-center text-sm text-slate-700 whitespace-nowrap w-[60px] h-[60pxs] ${(day.day == 'Sat' || day.day == 'Sun') ? "bg-gray-300" : ""} hover:bg-gray-300`}
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
