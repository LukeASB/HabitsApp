import { useState } from "react";

const daysOfWeek = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"];

const Calendar: React.FC = () => {
  const [currentDate, setCurrentDate] = useState(new Date());
  const [selectedDays, setSelectedDays] = useState<{ [key: string]: number[] }>({}); // Store selected days for each month/year
  const [counter, setCounter] = useState(0); // Track how many days are selected

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

  // Handle the day click event
  const handleDayClick = (day: number) => {
    const dateKey = `${year}-${month}`; // Use the combination of year and month as the key
    const selectedForMonth = selectedDays[dateKey] || []; // Get the selected days for this month or default to empty array

    if (selectedForMonth.includes(day)) {
      // If the day is already selected, remove it
      const updatedDays = selectedForMonth.filter((selectedDay) => selectedDay !== day);
      setSelectedDays({ ...selectedDays, [dateKey]: updatedDays });
      setCounter(counter - 1);
    } else {
      // If the day is not selected, add it
      const updatedDays = [...selectedForMonth, day];
      setSelectedDays({ ...selectedDays, [dateKey]: updatedDays });
      setCounter(counter + 1);
    }
  };

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
            {week.map((day, colIndex) => {
              if (!day) return <div key={colIndex} className="day-cell flex-fill text-center p-2 empty" style={{ width: "calc(100% / 7)" }}></div>;

              return (
                <div
                  key={colIndex}
                  className={`day-cell flex-fill text-center p-2 ${day ? "day" : "empty"}`}
                  style={{
                    width: "calc(100% / 7)", // Divide evenly into 7 columns
                    backgroundColor: selectedDays[`${year}-${month}`]?.includes(day) ? 'green' : '', // Set the background color for selected days
                    color: selectedDays[`${year}-${month}`]?.includes(day) ? 'white' : '', // Change text color for selected days
                    cursor: 'pointer', // Make it clear it's clickable
                  }}
                  onClick={() => day && handleDayClick(day)} // Only trigger click for valid days
                >
                  {day || ""}
                </div>
              )
            })}
          </div>
        ))}
      </div>

      <div className="counter mt-3">
        <p>Selected Days: {counter}</p>
      </div>
    </div>
  );
};

export default Calendar;
