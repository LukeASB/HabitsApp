import React, { Dispatch, SetStateAction, useState } from "react";
import IHabit from "../shared/interfaces/IHabit";
import ICalendar from "../shared/interfaces/ICalendar";

const daysOfWeek = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"];

const Calendar: React.FC<ICalendar> = ({ currentSelectedHabit, completionDatesCounter, setCompletionDatesCompletionDatesCounter, setCompletionDates, completionDates }) => {
	const [currentDate, setCurrentDate] = useState(new Date());

	const GenerateCalendarGrid = (currentDate: Date) => {
		const getDaysInMonth = (year: number, month: number) => new Date(year, month + 1, 0).getDate();
		const getFirstDayOfMonth = (year: number, month: number) => new Date(year, month, 1).getDay();

		const year = currentDate.getFullYear();
		const month = currentDate.getMonth();

		const daysInMonth = getDaysInMonth(year, month);
		const firstDay = getFirstDayOfMonth(year, month);

		const calendarDays: (number | null)[] = [];
		for (let i = 0; i < firstDay; i++) calendarDays.push(null);
		for (let day = 1; day <= daysInMonth; day++) calendarDays.push(day);
		while (calendarDays.length % 7 !== 0) calendarDays.push(null);

		const weeks: (number | null)[][] = [];
		for (let i = 0; i < calendarDays.length; i += 7) {
			weeks.push(calendarDays.slice(i, i + 7));
		}

		return { weeks, month, year };
	};

	const handleDaySelection = (habit: IHabit | null, year: number, month: number, day: number) => {
		if (!habit) return;

		const completedDate = `${year}-${String(month + 1).padStart(2, "0")}-${String(day).padStart(2, "0")}`;

		const removeCompletedDay = () => {
			setCompletionDatesCompletionDatesCounter(completionDatesCounter - 1);
			habit.completionDates = habit.completionDates.filter((date) => date !== completedDate);

			setCompletionDates(habit.completionDates);
		};

		const addCompletedDay = () => {
			setCompletionDatesCompletionDatesCounter(completionDatesCounter + 1);
			habit.completionDates.push(completedDate);

			setCompletionDates(habit.completionDates);
		};

		if (habit.completionDates.includes(completedDate)) {
			removeCompletedDay();
			return;
		}

		addCompletedDay();
	};

	const isCompletedDay = (day: number) => completionDates.includes(`${year}-${String(month + 1).padStart(2, "0")}-${String(day).padStart(2, "0")}`);

	// Handlers for navigating months
	const prevMonth = () => setCurrentDate((prev) => new Date(prev.getFullYear(), prev.getMonth() - 1, 1));
	const nextMonth = () => setCurrentDate((prev) => new Date(prev.getFullYear(), prev.getMonth() + 1, 1));

	const { weeks, month, year } = GenerateCalendarGrid(currentDate);

	return (
		<div id="calendar" className="calendar">
			{currentSelectedHabit && <h1>{currentSelectedHabit?.name}</h1>}
			<div className="calendar-header d-flex justify-content-between align-items-center mb-3">
				<button onClick={prevMonth} className="btn btn-secondary">
					&larr;
				</button>
				<h4>
					{currentDate.toLocaleString("default", { month: "long" })} {year}
				</h4>
				<button onClick={nextMonth} className="btn btn-secondary">
					&rarr;
				</button>
			</div>

			<div className="calendar-grid">
				<div id="weekday-headers" className="d-flex">
					{daysOfWeek.map((day) => (
						<div className="day-header flex-fill text-center fw-bold" key={day} style={{ width: "calc(100% / 7)" }}>
							{day}
						</div>
					))}
				</div>

				{weeks.map((week, rowIndex) => (
					<div id="calendar-row" key={rowIndex} className="d-flex">
						{week.map((day, colIndex) => {
							if (!day) return <div key={colIndex} className="day-cell flex-fill text-center p-2 empty" style={{ width: "calc(100% / 7)" }}></div>;

							return (
								<div
									key={colIndex}
									className={`day-cell flex-fill text-center p-2 ${day ? "day" : "empty"}`}
									style={{
										width: "calc(100% / 7)",
										backgroundColor: isCompletedDay(day) ? "green" : "transparent",
										color: isCompletedDay(day) ? "white" : "inherit",
										cursor: "pointer",
									}}
									onClick={() => day && handleDaySelection(currentSelectedHabit, year, month, day)}
								>
									{day || ""}
								</div>
							);
						})}
					</div>
				))}
			</div>

			<div className="completionDatesCounter mt-3">
				<p>Selected Days: {completionDatesCounter}</p>
			</div>
		</div>
	);
};

export default Calendar;
