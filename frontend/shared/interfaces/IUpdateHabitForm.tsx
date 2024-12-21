import IHabit from "./IHabit";

export default interface IUpdateHabitForm {
    habit: IHabit | null;
    onSubmit: (updatedHabit: IHabit) => void;
    onModalOpen: () => void;
    onModalClose: () => void;
}