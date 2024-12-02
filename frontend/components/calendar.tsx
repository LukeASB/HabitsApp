import React from "react";

const daysOfWeek = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"];

interface CalendarProps {
    selectedDays: { [key: string]: number[] };
    counter: number;
    handleDaySelection: (year: number, month: number, day: number) => void;
    completionDates: string[]; // New property
}

const Calendar: React.FC<CalendarProps> = ({ selectedDays, counter, handleDaySelection, completionDates }) => {
    const currentDate = new Date();
    const year = currentDate.getFullYear();
    const month = currentDate.getMonth();

    const getDaysInMonth = (year: number, month: number) => new Date(year, month + 1, 0).getDate();
    const getFirstDayOfMonth = (year: number, month: number) => new Date(year, month, 1).getDay();

    const daysInMonth = getDaysInMonth(year, month);
    const firstDay = getFirstDayOfMonth(year, month);

    // Generate a grid for the calendar
    const calendarDays: (number | null)[] = [];
    for (let i = 0; i < firstDay; i++) calendarDays.push(null);
    for (let day = 1; day <= daysInMonth; day++) calendarDays.push(day);
    while (calendarDays.length % 7 !== 0) calendarDays.push(null);

    const weeks: (number | null)[][] = [];
    for (let i = 0; i < calendarDays.length; i += 7) {
        weeks.push(calendarDays.slice(i, i + 7));
    }

    // Check if a day is in completionDates
    const isCompletedDay = (day: number) => {
        const formattedDate = `${year}-${String(month + 1).padStart(2, "0")}-${String(day).padStart(2, "0")}`;
        return completionDates.includes(formattedDate);
    };

    return (
        <div className="calendar">
            <div className="calendar-header d-flex justify-content-between align-items-center mb-3">
                <button className="btn btn-secondary">&larr;</button>
                <h4>
                    {currentDate.toLocaleString("default", { month: "long" })} {year}
                </h4>
                <button className="btn btn-secondary">&rarr;</button>
            </div>

            <div className="calendar-grid">
                {/* Weekday Headers */}
                <div className="d-flex">
                    {daysOfWeek.map((day) => (
                        <div
                            className="day-header flex-fill text-center fw-bold"
                            key={day}
                            style={{ width: "calc(100% / 7)" }}
                        >
                            {day}
                        </div>
                    ))}
                </div>

                {/* Calendar Rows */}
                {weeks.map((week, rowIndex) => (
                    <div key={rowIndex} className="d-flex">
                        {week.map((day, colIndex) => {
                            if (!day) {
                                return (
                                    <div
                                        key={colIndex}
                                        className="day-cell flex-fill text-center p-2 empty"
                                        style={{ width: "calc(100% / 7)" }}
                                    ></div>
                                );
                            }

                            return (
                                <div
                                    key={colIndex}
                                    className={`day-cell flex-fill text-center p-2 ${day ? "day" : "empty"}`}
                                    style={{
                                        width: "calc(100% / 7)",
                                        backgroundColor: isCompletedDay(day) || selectedDays[`${year}-${month}`]?.includes(day) ? "green" : "",
                                        color: isCompletedDay(day) || selectedDays[`${year}-${month}`]?.includes(day)? "white" : "",
                                        cursor: "pointer",
                                    }}
                                    onClick={() => day && handleDaySelection(year, month, day)}
                                >
                                    {day || ""}
                                </div>
                            );
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
