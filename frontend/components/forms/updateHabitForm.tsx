import React, { useState } from "react";
import IUpdateHabitForm from "../../shared/interfaces/IUpdateHabitForm";
import IUpdateHabitFormData from "../../shared/interfaces/IUpdateHabitFormData";
import IUpdateHabitFormError from "../../shared/interfaces/IUpdateHabitFormError";
import { HabitsModel } from "../../model/habitsModel";
import { ModalTypeEnum } from "../../shared/enum/modalTypeEnum";

const UpdateHabitForm: React.FC<IUpdateHabitForm> = ({ habit, onSubmit, onModalClose }) => {
	const form: IUpdateHabitFormData = { name: "", daysTarget: 0 };
	const [formData, setFormData] = useState<IUpdateHabitFormData>(form);
	const [errors, setErrors] = useState<IUpdateHabitFormError>({ name: "", daysTarget: "" });

	const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		const { name, value } = e.target;
		setFormData((prevData) => ({
			...prevData,
			[name]: name === "days" || name === "daysTarget" ? parseInt(value) || 0 : value,
		}));
	};

	const validateForm = () => {
		let isValid = true;
		const newErrors: IUpdateHabitFormError = { name: "", daysTarget: "" };
		const habitErrors = HabitsModel.processUpdateHabit({ ...formData });

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
            setErrors({ ...newErrors });
        } else {
            setErrors({ name: "", daysTarget: "" });
        }

		return isValid;
	};

	const handleSubmit = (e: React.FormEvent) => {
		e.preventDefault();
		if (!validateForm()) return;

		const updatedHabit = { ...habit, ...formData };
		onSubmit(updatedHabit);
		onModalClose(ModalTypeEnum.UpdateHabitModal);
		setFormData({ name: "", daysTarget: 0 });
	};

	return (
		<div id="updateHabitForm" className="updateHabitForm">
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
				<button type="submit" className="btn btn-primary w-100" data-bs-dismiss="modal" onClick={handleSubmit}>
					Update Habit
				</button>
			</div>
		</div>
	);
};

export default UpdateHabitForm;
