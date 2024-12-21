import IModal from "../../shared/interfaces/IModal";
  

const Modal: React.FC<IModal> = (modal: IModal) => {
    return (
        <div
          className={`modal fade ${modal.showModal ? "show d-block" : ""}`}
          id={modal.id}
          tabIndex={-1}
          aria-labelledby={`${modal.id}Label`}
          aria-hidden={!modal.showModal}
          style={{ backgroundColor: modal.showModal ? "rgba(0, 0, 0, 0.5)" : "transparent" }} // Optional: Add dimming effect
        >
          <div className="modal-dialog modal-dialog-centered" role="document">
            <div className="modal-content">
              <div className="modal-header text-black">
                <h5 className="modal-title" id={`${modal.id}Label`}>
                  {modal.title}
                </h5>
                <button
                  type="button"
                  className="btn-close"
                  onClick={() => modal.onModalClose(modal.modalType)} // Handle modal close
                  aria-label="Close"
                ></button>
              </div>
              <div className="modal-body text-black text-start">{modal.body}</div>
            </div>
          </div>
        </div>
      );
};

export default Modal;
