import { useState, useEffect } from "react";
import Link from "next/link";
import IHabit from "../shared/interfaces/IHabit";
import HabitsButtons from "./habitButtons";
import CreateHabitForm from "./forms/createHabitForm";
import UpdateHabitForm from "./forms/updateHabitForm";
import DeleteHabitForm from "./forms/deleteHabitForm";
import IHabitsNavbar from "../shared/interfaces/IHabitsNavbar";
import IHabitModalTypes from "../shared/interfaces/IHabitModalTypes";
import { ModalTypeEnum } from "../shared/enum/modalTypeEnum";
import { useRouter } from "next/router";
import { AuthModel } from "../model/authModel";

const HabitsNavbar: React.FC<IHabitsNavbar> = ({ showSidebar, setShowSidebar, habit, habitOps: { createHabit, updateHabit, deleteHabit } }) => {
	const router = useRouter();
	const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);
	const [showModal, setShowModal] = useState<IHabitModalTypes>({ createHabitModal: false, updateHabitModal: false, deleteHabitModal: false});

	useEffect(() => {
		const jwt = sessionStorage.getItem("access-token");
		if (!jwt) {
			setIsLoggedIn(false);
			router.push("/login");
			return;
		}

		if (jwt) {
			if (!AuthModel.parseJWT(jwt)) {
				setIsLoggedIn(false);
				router.push("/login");
				return;
			}
		}
		setIsLoggedIn(true);
	}, []);

	const handleOpenModal = (modalType: ModalTypeEnum) => {
		if (modalType === ModalTypeEnum.CreateHabitModal) return setShowModal({ createHabitModal: true, updateHabitModal: false, deleteHabitModal: false });
		if (modalType === ModalTypeEnum.UpdateHabitModal) return setShowModal({ createHabitModal: false, updateHabitModal: true, deleteHabitModal: false });
		if (modalType === ModalTypeEnum.DeleteHabitModal) return setShowModal({ createHabitModal: false, updateHabitModal: false, deleteHabitModal: true });
	};

	const handleCloseModal = (modalType: ModalTypeEnum) => {
		if (modalType === ModalTypeEnum.CreateHabitModal) return setShowModal({ createHabitModal: false, updateHabitModal: false, deleteHabitModal: false });
		if (modalType === ModalTypeEnum.UpdateHabitModal) return setShowModal({ createHabitModal: false, updateHabitModal: false, deleteHabitModal: false });
		if (modalType === ModalTypeEnum.DeleteHabitModal) return setShowModal({ createHabitModal: false, updateHabitModal: false, deleteHabitModal: false });
	};

	return (
		<nav className="navbar navbar-expand-lg navbar-light bg-primary">
			<div className="container-fluid d-flex justify-content-between align-items-center">
				{/* Right Section - Content */}
				<div className="d-flex gap-2">
					<button className="btn btn-dark" type="button" onClick={() => setShowSidebar(true)}>
						<i className={`bi bi-list`}></i>
					</button>
					<strong>
						<Link className="navbar-brand text-light" href="/">
							{habit ? habit.name : "All Habits"}
						</Link>
					</strong>
				</div>
				{/* Left Section - Habits Buttons */}
				<div className="d-flex gap-2">
					{isLoggedIn && (
						<>
							<HabitsButtons
								icon="plus"
								modal={{
									id: "createHabitModal",
									title: "Create Habit",
									body: <CreateHabitForm onSubmit={createHabit} onModalClose={handleCloseModal} />,
									modalType: ModalTypeEnum.CreateHabitModal,
									showModal: showModal.createHabitModal,
									onModalOpen: handleOpenModal,
									onModalClose: handleCloseModal,
								}}
								onClick={createHabit}
							/>
							{habit && (
								<HabitsButtons
									icon="gear"
									modal={{
										id: "updateHabitModal",
										title: "Update Habit",
										body: <UpdateHabitForm habit={habit} onSubmit={updateHabit} onModalClose={handleCloseModal} />,
										modalType: ModalTypeEnum.UpdateHabitModal,
										showModal: showModal.updateHabitModal,
										onModalOpen: handleOpenModal,
										onModalClose: handleCloseModal,
									}}
									onClick={updateHabit}
								/>
							)}
							{habit && (
								<HabitsButtons
									icon="x"
									modal={{
										id: "deleteHabitModal",
										title: "Delete Habit",
										body: <DeleteHabitForm habit={habit} onSubmit={deleteHabit} onModalClose={handleCloseModal} />,
										modalType: ModalTypeEnum.DeleteHabitModal,
										showModal: showModal.deleteHabitModal,
										onModalOpen: handleOpenModal,
										onModalClose: handleCloseModal,
									}}
									onClick={deleteHabit}
								/>
							)}
						</>
					)}
				</div>
			</div>
		</nav>
	);
};

export default HabitsNavbar;
