import { ModalTypeEnum } from "../enum/modalTypeEnum";

export default interface IGenericErrorForm {
    error: string;
    modalType: ModalTypeEnum;
    onSubmit: (error: string) => void;
    onModalClose: (modalType: ModalTypeEnum) => void;
}
  