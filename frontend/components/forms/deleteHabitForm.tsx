import React from "react";
import IDeleteHabitForm from "../../shared/interfaces/IDeleteHabitForm";
import { ModalTypeEnum } from "../../shared/enum/modalTypeEnum";

const DeleteHabitForm: React.FC<IDeleteHabitForm> = ({ habit, onSubmit, onModalClose }) => {
	const handleConfirm = () => {
		onSubmit(habit);
        onModalClose(ModalTypeEnum.DeleteHabitModal);
	};

	return (
		<div className="deleteForm">
			<p>{`Are you sure you want to delete: ${habit.name}`}</p>
			<div className="button-group">
				<button className="btn btn-danger" data-bs-dismiss="modal" onClick={handleConfirm}>
					Confirm
				</button>
			</div>
		</div>
	);
};

export default DeleteHabitForm;
