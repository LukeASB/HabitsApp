import React, { useState } from "react";
import ICreateHabitForm from "../../shared/interfaces/ICreateHabitForm";
import IHabit from "../../shared/interfaces/IHabit";
import { HabitsModel } from "../../model/habitsModel";
import ICreateHabitFormData from "../../shared/interfaces/ICreateHabitFormData";
import { ModalTypeEnum } from "../../shared/enum/modalTypeEnum";
import ICreateHabitFormError from "../../shared/interfaces/ICreateHabitFormError";

const CreateHabitForm: React.FC<ICreateHabitForm> = ({ onSubmit, onModalClose }) => {
	const form: ICreateHabitFormData = { name: "", daysTarget: 0 };
	const [formData, setFormData] = useState<ICreateHabitFormData>(form);
	const [errors, setErrors] = useState<ICreateHabitFormError>({ name: "", daysTarget: "" });

	const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		const { name, value } = e.target;
		setFormData((prevData) => ({
			...prevData,
			[name]: name === "daysTarget" ? parseInt(value) || 0 : value,
		}));
	};

	const validateForm = () => {
		let isValid = true;
		const newErrors: ICreateHabitFormError = { name: "", days: "", daysTarget: "" };
		const habitErrors = HabitsModel.processCreateHabit({ ...formData });

		habitErrors.forEach((error) => {
			if (error.name) {
				isValid = false;
				return (newErrors.name = error.name);
			}

			if (error.daysTarget) {
				isValid = false;
				return (newErrors.daysTarget = error.daysTarget);
			}
		});

        if (!isValid) {
            setErrors({ ...newErrors })
        } else {
            setErrors({ name: "", days: "", daysTarget: "" });
        }

		return isValid;
	};

	const handleSubmit = (e: React.FormEvent) => {
		e.preventDefault();
		if (!validateForm()) return;

		const createdHabit: IHabit = { habitId: "", days: 0, createdAt: 0, numberOfDays: 0, completionDates: [], ...formData };
		onSubmit(createdHabit);
		onModalClose(ModalTypeEnum.CreateHabitModal);
		setFormData({ name: "", days: 0, daysTarget: 0 });
	};

	return (
		<div id="createHabitForm" className="createHabitForm">
			<div className="shadow-sm bg-white rounded">
				<div className="form-group mb-3">
					<label htmlFor="name" className="form-label">
						Name
					</label>
					<input type="text" id="name" name="name" value={formData.name} onChange={handleChange} placeholder="Enter habit name" className="form-control" />
					<div className="error text-danger">{errors.name}</div>
				</div>
				<div className="form-group mb-3">
					<label htmlFor="daysTarget" className="form-label">
						Days Target
					</label>
					<input
						type="text"
						id="daysTarget"
						name="daysTarget"
						value={formData.daysTarget}
						onChange={handleChange}
						placeholder="Enter target days"
						className="form-control"
						inputMode="numeric"
						pattern="[0-9]*"
					/>
					<div className="error text-danger">{errors.daysTarget}</div>
				</div>
				<button type="submit" className="btn btn-primary w-100" onClick={handleSubmit}>
					Create Habit
				</button>
			</div>
		</div>
	);
};

export default CreateHabitForm;
