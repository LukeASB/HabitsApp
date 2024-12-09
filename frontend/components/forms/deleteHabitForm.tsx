import React from "react";
import IHabit from "../../shared/interfaces/IHabit";

interface IDeleteHabitForm {
  habit: IHabit;
  modalId: string;
  onSubmit: (updatedHabit: IHabit) => void;
}

const DeleteHabitForm: React.FC<IDeleteHabitForm> = ({ habit, modalId, onSubmit }) => {
  const handleConfirm = () => {
    onSubmit(habit);
  };

  return (
    <div className="deleteForm">
      <p>{`Are you sure you want to delete: ${habit.name}`}</p>
      <div className="button-group">
        <button className="btn btn-danger" data-bs-dismiss="modal" onClick={(handleConfirm)}>
          Confirm
        </button>
      </div>
    </div>
  );
};

export default DeleteHabitForm;
