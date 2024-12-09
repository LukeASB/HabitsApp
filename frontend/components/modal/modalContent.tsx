const ModalContent: React.FC<IModalContent> = (modalContent: IModalContent) => {
    return (
        <>
            <h6>{modalContent.content.title}</h6>
            {modalContent.content.description}
        </>
    )
}

export default ModalContent;