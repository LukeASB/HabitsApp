import { ModalTypeEnum } from "../enum/modalTypeEnum";
import IHabit from "./IHabit";

export default interface IDeleteHabitForm {
    habit: IHabit;
    onSubmit: (updatedHabit: IHabit) => void;
    onModalClose: (modalType: ModalTypeEnum) => void;
}
  