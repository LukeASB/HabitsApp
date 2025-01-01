import IHabitsButton from "../shared/interfaces/IHabitsButton";
import Modal from "./modal/modal";

const HabitsButtons: React.FC<IHabitsButton> = ({ icon, modal }) => {
	return (
		<div>
			<button type="button" className="btn btn-dark btn-custom robo popup-trigger popmake-680" data-popup-id="680" data-do-default="0" onClick={() => modal.onModalOpen(modal.modalType)}>
				<i className={`bi bi-${icon}`}></i>
			</button>
			<Modal
				id={modal.id}
				title={modal.title}
				body={modal.body}
				showModal={modal.showModal}
				modalType={modal.modalType}
				onModalOpen={() => modal.onModalOpen(modal.modalType)}
				onModalClose={() => modal.onModalClose(modal.modalType)}
			/>
		</div>
	);
};

export default HabitsButtons;
