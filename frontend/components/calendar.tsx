import { useState } from "react";

const daysOfWeek = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"];

const Calendar: React.FC = () => {
  const [currentDate, setCurrentDate] = useState(new Date());

  const getDaysInMonth = (year: number, month: number) => new Date(year, month + 1, 0).getDate();
  const getFirstDayOfMonth = (year: number, month: number) => new Date(year, month, 1).getDay();

  const prevMonth = () => setCurrentDate((prev) => new Date(prev.getFullYear(), prev.getMonth() - 1, 1));
  const nextMonth = () => setCurrentDate((prev) => new Date(prev.getFullYear(), prev.getMonth() + 1, 1));

  const year = currentDate.getFullYear();
  const month = currentDate.getMonth();
  const daysInMonth = getDaysInMonth(year, month);
  const firstDay = getFirstDayOfMonth(year, month);

  // Generate a grid for the calendar
  const calendarDays: (number | null)[] = [];

  // Add empty cells for days before the first day of the month
  for (let i = 0; i < firstDay; i++) {
    calendarDays.push(null);
  }

  // Add the days of the month
  for (let day = 1; day <= daysInMonth; day++) {
    calendarDays.push(day);
  }

  // Add empty cells to complete the final row (if needed)
  while (calendarDays.length % 7 !== 0) {
    calendarDays.push(null);
  }

  // Group days into weeks for rendering rows
  const weeks = [];
  for (let i = 0; i < calendarDays.length; i += 7) {
    weeks.push(calendarDays.slice(i, i + 7));
  }

  return (
    <div className="calendar">
      <div className="calendar-header d-flex justify-content-between align-items-center mb-3">
        <button onClick={prevMonth} className="btn btn-secondary">&larr;</button>
        <h4>
          {currentDate.toLocaleString("default", { month: "long" })} {year}
        </h4>
        <button onClick={nextMonth} className="btn btn-secondary">&rarr;</button>
      </div>

      <div className="calendar-grid">
        {/* Weekday Headers */}
        <div className="d-flex">
          {daysOfWeek.map((day) => (
            <div className="day-header flex-fill text-center fw-bold" key={day} style={{
                width: "calc(100% / 7)", // Divide evenly into 7 columns
              }}>
              {day}
            </div>
          ))}
        </div>

        {/* Calendar Rows */}
        {weeks.map((week, rowIndex) => (
          <div key={rowIndex} className="d-flex">
            {week.map((day, colIndex) => (
              <div
                key={colIndex}
                className={`day-cell flex-fill text-center p-2 ${day ? "day" : "empty"}`}
                style={{
                  width: "calc(100% / 7)", // Divide evenly into 7 columns
                }}
              >
                {day || ""}
              </div>
            ))}
          </div>
        ))}
      </div>
    </div>
  );
};

export default Calendar;
