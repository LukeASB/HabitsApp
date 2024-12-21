import IHabit from "./IHabit";

export default interface ICreateHabitForm {
    onSubmit: (updatedHabit: IHabit) => void;
    onModalOpen: () => void;
    onModalClose: () => void;

}