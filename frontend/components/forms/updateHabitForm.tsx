import React, { useState } from "react";
import IUpdateHabitForm from "../../shared/interfaces/IUpdateHabitForm";

const UpdateHabitForm: React.FC<IUpdateHabitForm> = ({ habit, onSubmit }) => {
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
		if (!habit) return;
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
				{`Update Habit`}
			</button>
		</>
	);
};

export default UpdateHabitForm;
