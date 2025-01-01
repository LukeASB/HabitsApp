import React, { useState } from "react";
import ICreateHabitForm from "../../shared/interfaces/ICreateHabitForm";
import IHabit from "../../shared/interfaces/IHabit";
import { HabitsModel } from "../../model/habitsModel";
import ICreateHabitFormData from "../../shared/interfaces/ICreateHabitFormData";
import { ModalTypeEnum } from "../../shared/enum/modalTypeEnum";
import ICreateHabitFormError from "../../shared/interfaces/ICreateHabitFormError";

const CreateHabitForm: React.FC<ICreateHabitForm> = ({ onSubmit, onModalClose }) => {
	const form: ICreateHabitFormData = { name: "", days: 0, daysTarget: 0 };
	const [formData, setFormData] = useState<ICreateHabitFormData>(form);
	const [errors, setErrors] = useState<ICreateHabitFormError>({ name: "", days: "", daysTarget: "" });

	const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		const { name, value } = e.target;
		setFormData((prevData) => ({
			...prevData,
			[name]: name === "days" || name === "daysTarget" ? parseInt(value) || 0 : value,
		}));
	};

	const validateForm = () => {
		let isValid = true;
		let newErrors: ICreateHabitFormError = { name: "", days: "", daysTarget: "" };
		const habitErrors = HabitsModel.processCreateHabit({ ...formData });

		habitErrors.forEach((error) => {
			if (error.name) {
				isValid = false;
				return (newErrors.name = error.name);
			}

			if (error.days) {
				isValid = false;
				return (newErrors.days = error.days);
			}

			if (error.daysTarget) {
				isValid = false;
				return (newErrors.daysTarget = error.daysTarget);
			}
		});

		!isValid ? setErrors({ ...newErrors }) : setErrors({ name: "", days: "", daysTarget: "" });

		return isValid;
	};

	const handleSubmit = (e: React.FormEvent) => {
		e.preventDefault();
		if (!validateForm()) return;

		const createdHabit: IHabit = { habitId: "", createdAt: 0, numberOfDays: 0, completionDates: [], ...formData };
		onSubmit(createdHabit);
		onModalClose(ModalTypeEnum.CreateHabitModal);
		setFormData({ name: "", days: 0, daysTarget: 0 });
	};

	return (
		<>
			<div className="form-group">
				<label htmlFor="name">Name</label>
				<input type="text" id="name" name="name" value={formData.name} onChange={handleChange} placeholder="Enter habit name" className="form-control" />
				<div className="error text-danger">{errors.name}</div>
			</div>
			<div className="form-group">
				<label htmlFor="days">Days</label>
				<input type="number" id="days" name="days" value={formData.days} onChange={handleChange} placeholder="Enter current days" className="form-control" />
				<div className="error text-danger">{errors.days}</div>
			</div>
			<div className="form-group">
				<label htmlFor="daysTarget">Days Target</label>
				<input type="number" id="daysTarget" name="daysTarget" value={formData.daysTarget} onChange={handleChange} placeholder="Enter target days" className="form-control" />
				<div className="error text-danger">{errors.daysTarget}</div>
			</div>
			<button type="submit" className="btn btn-primary" onClick={handleSubmit}>
				{`Create Habit`}
			</button>
		</>
	);
};

export default CreateHabitForm;
