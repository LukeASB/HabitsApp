import React, { useState } from "react";
import IHabit from "../shared/interfaces/IHabit";
import ICalendar from "../shared/interfaces/ICalendar";
import { FormatDate } from "../shared/helpers/formatDate";

const daysOfWeek = ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"];

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
		for (let i = 0; i < calendarDays.length; i += 7) weeks.push(calendarDays.slice(i, i + 7));

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
			{currentSelectedHabit && <h1 className="text-center text-dark">{currentSelectedHabit?.name}</h1>}
			<div className="calendar-header d-flex justify-content-between align-items-center mb-3">
				<button onClick={prevMonth} className="btn btn-outline-dark">
					<i className="bi bi-chevron-left"></i>
				</button>
				<h4 className="text-dark">
					{currentDate.toLocaleString("default", { month: "long" })} {year}
				</h4>
				<button onClick={nextMonth} className="btn btn-outline-dark">
					<i className="bi bi-chevron-right"></i>
				</button>
			</div>

			<div className="calendar-grid">
				<div id="weekday-headers" className="d-flex">
					{daysOfWeek.map((day) => (
						<div className="day-header flex-fill text-center fw-bold text-dark" key={day} style={{ width: "calc(100% / 7)", borderRadius: "5px", margin: "2px" }}>
							{day}
						</div>
					))}
				</div>

				{weeks.map((week, rowIndex) => (
					<div id="calendar-row" key={rowIndex} className="d-flex">
						{week.map((day, colIndex) => {
							return (
								<div
									key={colIndex}
									className={`day-cell flex-fill text-center p-2 ${day ? "day" : "empty"}`}
									style={{
										width: "calc(100% / 7)",
										height: "50px",
										backgroundColor: day && isCompletedDay(day) ? "green" : "transparent",
										color: day && isCompletedDay(day) ? "white" : "inherit",
										cursor: day ? "pointer" : "default",
										border: "1px solid #ccc",
										borderRadius: "5px",
										margin: "2px",
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

			<div className="card mt-3 p-3 shadow-sm">
				<div className="card-body">
					<p>
						<strong>Completed Days Total:</strong> {completionDatesCounter}
					</p>
					<p>
						<strong>Days Target:</strong> {currentSelectedHabit?.daysTarget}
					</p>
					<p>
						<strong>Created At:</strong> {FormatDate(currentSelectedHabit?.createdAt ? `${currentSelectedHabit?.createdAt}` : "")}
					</p>
				</div>
			</div>
		</div>
	);
};

export default Calendar;
