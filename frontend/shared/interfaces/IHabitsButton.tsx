import IHabit from "./IHabit";
import IModal from "./IModal";

export default interface IHabitsButton {
    icon: string;
    modal: IModal;
    onClick: (habit: IHabit) => void;
}