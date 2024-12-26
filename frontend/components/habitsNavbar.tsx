import { useState, useEffect } from "react";
import Link from "next/link";
import IHabit from "../shared/interfaces/IHabit";
import HabitsButtons from "./habitButtons";
import CreateHabitForm from "./forms/createHabitForm";
import UpdateHabitForm from "./forms/updateHabitForm";
import DeleteHabitForm from "./forms/deleteHabitForm";
import IHabitsNavbar from "../shared/interfaces/IHabitsNavbar";
import { HabitsService } from "../services/habitsService";
import IModalTypes from "../shared/interfaces/IModalTypes";
import { ModalTypeEnum } from "../shared/enum/modalTypeEnum";
import { useRouter } from "next/router";
import { AuthModel } from "../model/authModel";


const HabitsNavbar: React.FC<IHabitsNavbar> = ({ habit, updateMain }) => {
    const router = useRouter();
	const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);
    const [showModal, setShowModal] = useState<IModalTypes>({ createHabitModal: false, updateHabitModal: false, deleteHabitModal: false });

	const test = "debug";
	useEffect(() => {
        // sessionStorage.getItem("access-token") || test ? setIsLoggedIn(true) : setIsLoggedIn(false);

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

	const createHabit = async (habit: IHabit) => {
		await HabitsService.createHabit(habit);
		updateMain(habit, true);
	};

	const updateHabit = async (habit: IHabit) => {
		const data = await HabitsService.updateHabit(habit);
		updateMain(data, true);
	};

	const deleteHabit = async (habit: IHabit | null) => {
		if (!habit) return;
        await HabitsService.deleteHabit(habit.id);
		updateMain(null, true);
	};

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
			<div className="container content">
				<strong>
					<Link className="navbar-brand text-light" href="/">
						{habit ? habit.name : "All Habits"}
					</Link>
				</strong>
			</div>
			<div className="navbar" id="navbarNav">
				<ul className="navbar-nav">
					<li className="nav-item active">
						{isLoggedIn && (
							<div className="d-flex gap-2">
								{/* Plus Icon Button */}
								<HabitsButtons
									icon="plus"
									modal={{
										id: "createHabitModal",
										title: "Create Habit",
										body: <CreateHabitForm onSubmit={createHabit} onModalClose={handleCloseModal} />,
                                        modalType: ModalTypeEnum.CreateHabitModal,
                                        showModal: showModal.createHabitModal,
                                        onModalOpen: handleOpenModal,
                                        onModalClose: handleCloseModal
									}}
									onClick={createHabit}
								/>
								{/* Update Icon Button */}
								{habit && (
									<HabitsButtons
										icon="gear"
										modal={{
											id: "updateHabitModal",
											title: "Update Habit",
											body: (
												<UpdateHabitForm
                                                    habit={habit}
													onSubmit={updateHabit}
                                                    onModalClose={handleCloseModal}
												/>
											),
                                            modalType: ModalTypeEnum.UpdateHabitModal,
                                            showModal: showModal.updateHabitModal,
                                            onModalOpen: handleOpenModal,
                                            onModalClose: handleCloseModal
										}}
										onClick={updateHabit}
									/>
								)}
								{/* X Icon Button */}
								{habit && (
									<HabitsButtons
										icon="x"
										modal={{
											id: "deleteHabitModal",
											title: "Delete Habit",
											body: (
												<DeleteHabitForm
													habit={habit}
													onSubmit={deleteHabit}
                                                    onModalClose={handleCloseModal}
												/>
											),
                                            modalType: ModalTypeEnum.DeleteHabitModal,
                                            showModal: showModal.deleteHabitModal,
                                            onModalOpen: handleOpenModal,
                                            onModalClose: handleCloseModal
										}}
										onClick={deleteHabit}
									/>
								)}
							</div>
						)}
					</li>
				</ul>
			</div>
		</nav>
	);
};

export default HabitsNavbar;
