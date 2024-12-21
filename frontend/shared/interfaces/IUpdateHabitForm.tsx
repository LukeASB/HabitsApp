import { ModalTypeEnum } from "../enum/modalTypeEnum";
import IHabit from "./IHabit";

export default interface IUpdateHabitForm {
    habit: IHabit;
    onSubmit: (updatedHabit: IHabit) => void;
    onModalClose: (modalType: ModalTypeEnum) => void;
}