import { ModalTypeEnum } from "../enum/modalTypeEnum";
import IHabit from "./IHabit";

export default interface ICreateHabitForm {
    onSubmit: (updatedHabit: IHabit) => void;
    onModalOpen: (modalType: ModalTypeEnum) => void;
    onModalClose: (modalType: ModalTypeEnum) => void;

}