import React from "react";
import IDeleteHabitForm from "../../shared/interfaces/IDeleteHabitForm";
import { ModalTypeEnum } from "../../shared/enum/modalTypeEnum";

const DeleteHabitForm: React.FC<IDeleteHabitForm> = ({ habit, onSubmit, onModalClose }) => {
	const handleConfirm = () => {
		onSubmit(habit);
		onModalClose(ModalTypeEnum.DeleteHabitModal);
	};

    return (
        <div id="deleteHabitForm" className="deleteHabitForm">
            <div className="shadow-sm bg-white rounded p-4">
                <p className="mb-4">{`Are you sure you want to delete: ${habit.name}`}</p>
                <div className="d-flex justify-content-center">
                    <button className="btn btn-danger me-2" data-bs-dismiss="modal" onClick={handleConfirm}>
                        Confirm
                    </button>
                </div>
            </div>
        </div>
    );
};

export default DeleteHabitForm;
