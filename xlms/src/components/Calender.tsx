import { useEffect, useState } from "react";
import ChevronRightIcon from "@mui/icons-material/ChevronRight";
import KeyboardArrowLeftIcon from "@mui/icons-material/KeyboardArrowLeft";

export type MonthInfo = {
  month: number;
  days: number;
  year: number;
  monthStartDate: Date;
};

type CalendarProps = {
  setMonth: (month: MonthInfo) => void;
};

export default function Calendar({ setMonth }: CalendarProps) {
  const [monthInfo, setMonthInfo] = useState<MonthInfo | null>(null);

  const getStartDateOfMonth = (date: Date): Date =>
    new Date(date.getFullYear(), date.getMonth(), 1);

  const getDaysInMonth = (date: Date): number =>
    new Date(date.getFullYear(), date.getMonth() + 1, 0).getDate();

  const buildMonthInfo = (date: Date): MonthInfo => ({
    month: date.getMonth(),
    year: date.getFullYear(),
    days: getDaysInMonth(date),
    monthStartDate: getStartDateOfMonth(date),
  });

  const nextMonth = () => {
    if (!monthInfo) return;

    const nextDate = new Date(monthInfo.year, monthInfo.month + 1, 1);
    const info = buildMonthInfo(nextDate);

    setMonthInfo(info);
    setMonth(info); // ✅ send to parent
  };

  const prevMonth = () => {
    if (!monthInfo) return;

    const prevDate = new Date(monthInfo.year, monthInfo.month - 1, 1);
    const info = buildMonthInfo(prevDate);

    setMonthInfo(info);
    setMonth(info); // ✅ send to parent
  };

  // Init
  useEffect(() => {
    const info = buildMonthInfo(new Date());
    setMonthInfo(info);
    setMonth(info); // ✅ initial sync with parent
  }, []);

  if (!monthInfo) return null;

  const monthName = monthInfo.monthStartDate.toLocaleDateString("en-US", {
    month: "long",
  });

  return (
    <div className="flex items-center gap-2 ml-3">
      <KeyboardArrowLeftIcon
        color="primary"
        fontSize="large"
        className="cursor-pointer"
        onClick={prevMonth}
      />

      <span className="w-40 font-semibold text-center">
        {monthName} - {monthInfo.year}
      </span>

      <ChevronRightIcon
        color="primary"
        fontSize="large"
        className="cursor-pointer"
        onClick={nextMonth}
      />
    </div>
  );
}
