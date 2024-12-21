interface IModal {
    id: string;
    title: string;
    body: JSX.Element;
    showModal: boolean;
    onModalOpen: () => void;
    onModalClose: () => void;
}

export default IModal;