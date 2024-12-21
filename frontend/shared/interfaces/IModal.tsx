import { ModalTypeEnum } from "../enum/modalTypeEnum";

interface IModal {
    id: string;
    title: string;
    body: JSX.Element;
    showModal: boolean;
    modalType: ModalTypeEnum;
    onModalOpen: (modalType: ModalTypeEnum) => void;
    onModalClose: (modalType: ModalTypeEnum) => void;
}

export default IModal;