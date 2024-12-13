import React, { useState } from "react";
import IHabit from "../../shared/interfaces/IHabit";
import { mockhabits } from "../../data/mock_habits";
import ICreateHabitForm from "../../shared/interfaces/ICreateHabitForm";

const CreateHabitForm: React.FC<ICreateHabitForm> = ({ onSubmit }) => {
	let habit: IHabit = {
		id: mockhabits.length > 0 ? String(parseInt(mockhabits[mockhabits.length - 1]?.id) + 1) : "1",
		createdAt: Date.now(),
		name:
			mockhabits.length > 0
				? `Added Habit_${String(parseInt(mockhabits[mockhabits.length - 1].id) + 1)}`
				: "Added Habit_1",
		days: 0,
		daysTarget: 30,
		numberOfDays: 1,
		completionDates: ["2024-12-20", "2024-12-19", "2024-12-17"],
	};

	const [formData, setFormData] = useState({
		name: habit?.name || "",
		days: habit?.days || 0,
		daysTarget: habit?.daysTarget || 0,
	});

	const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		const { name, value } = e.target;
		setFormData((prevData) => ({
			...prevData,
			[name]: name === "days" || name === "daysTarget" ? parseInt(value) || 0 : value,
		}));
	};

	const handleSubmit = (e: React.FormEvent) => {
		e.preventDefault();
		const updatedHabit = { ...habit, ...formData };
		onSubmit(updatedHabit);
	};

	return (
		<>
			<div className="form-group">
				<label htmlFor="name">Name</label>
				<input
					type="text"
					id="name"
					name="name"
					value={formData.name}
					onChange={handleChange}
					placeholder="Enter habit name"
					className="form-control"
				/>
			</div>
			<div className="form-group">
				<label htmlFor="days">Days</label>
				<input
					type="number"
					id="days"
					name="days"
					value={formData.days}
					onChange={handleChange}
					placeholder="Enter current days"
					className="form-control"
				/>
			</div>
			<div className="form-group">
				<label htmlFor="daysTarget">Days Target</label>
				<input
					type="number"
					id="daysTarget"
					name="daysTarget"
					value={formData.daysTarget}
					onChange={handleChange}
					placeholder="Enter target days"
					className="form-control"
				/>
			</div>
			<button type="submit" className="btn btn-primary" data-bs-dismiss="modal" onClick={handleSubmit}>
				{`Create Habit`}
			</button>
		</>
	);
};

export default CreateHabitForm;
