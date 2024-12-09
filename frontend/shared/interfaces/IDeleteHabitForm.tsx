import IHabit from "./IHabit";

export default interface IDeleteHabitForm {
    habit: IHabit;
    modalId: string;
    onSubmit: (updatedHabit: IHabit) => void;
}
  