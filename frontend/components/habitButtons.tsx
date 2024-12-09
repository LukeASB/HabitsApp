import IHabit from "../shared/interfaces/IHabit";
import IModal from "../shared/interfaces/IModal";
import Modal from "./modal/modal";

interface IHabitsButton {
    icon: string;
    modal: IModal
    onClick: (habit: IHabit) => void;
}

const HabitsButtons: React.FC<IHabitsButton> = ({ icon, modal, onClick}) => {
    return (
        <div>
            <button
                type="button"
                className="btn btn-dark btn-custom robo popup-trigger popmake-680"
                data-popup-id="680"
                data-do-default="0"
                data-bs-toggle="modal"
                data-bs-target={`#${modal.id}`}
            >
                <i className={`bi bi-${icon}`}></i>
            </button>
            <Modal
                id={`${modal.id}`}
                title={`${modal.title}`}
                body={modal.body}
            />
            
        </div>

    );
}

export default HabitsButtons;