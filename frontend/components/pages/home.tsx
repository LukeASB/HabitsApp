import { useState, useEffect } from "react";
import Calendar from "../calendar";
import Sidebar from "../sidebar";
import IHabit from "../../shared/interfaces/IHabit";
import { mockhabits } from "../../data/mock_habits";
import HabitsNavbar from "../habitsNavbar";
import { HabitsService } from "../../services/habitsService";
import { useRouter } from "next/router";
import { ModalTypeEnum } from "../../shared/enum/modalTypeEnum";
import IGenericModalTypes from "../../shared/interfaces/IGenericModalTypes";
import Modal from "../modal/modal";
import GenericErrorForm from "../forms/genericErrorForm";

const Home: React.FC = () => {
	const router = useRouter();
	const [hasHabitsBeenUpdated, setHasHabitsBeenUpdated] = useState<boolean>(true);
	const [habitNavbar, setHabitNavbar] = useState<IHabit | null>(null);
	const [habitsMenu, setHabitsMenu] = useState<IHabit[]>([]);
	const [currentSelectedHabit, setCurrentSelectedHabit] = useState<IHabit | null>(null);
	const [showNoHabitsMessage, setShowNoHabitsMessage] = useState<boolean>(false);
    const [showModal, setShowModal] = useState<IGenericModalTypes>({ GenericErrorModal: false });

    const handleOpenModal = (modalType: ModalTypeEnum) => {
        if (modalType === ModalTypeEnum.GenericErrorModal) return setShowModal({ GenericErrorModal: true });
    };

    const handleCloseModal = (modalType: ModalTypeEnum) => {
        if (modalType === ModalTypeEnum.GenericErrorModal) return setShowModal({ GenericErrorModal: false });
    };

    const onGenericErrorModalSubmit = () => {
        setShowModal({ GenericErrorModal: false });
    };

	useEffect(() => {
		if (sessionStorage.getItem("access-token") === null) {
			router.push("/login");
			return;
		}

		if (!hasHabitsBeenUpdated) return;
		if (process.env.ENVIRONMENT === "DEV") {
			setHabitsMenu(mockhabits);
			setHasHabitsBeenUpdated(false);
			return;
		}

		const retrieveHabits = async () => {
			const usersHabits = await HabitsService.retrieveHabits(router);

			if (usersHabits.length === 0) {
                setShowNoHabitsMessage(true);
            } else {
                setShowNoHabitsMessage(false);
            }

			setHabitsMenu(usersHabits);
		};

		retrieveHabits();
		setHasHabitsBeenUpdated(false);
	}, [router, hasHabitsBeenUpdated]);

	// Sidebar State
	const [showSidebar, setShowSidebar] = useState(false);

	// Calendar State
	const [completionDates, setCompletionDates] = useState<string[]>([]);
	const [completionDatesCounter, setCompletionDatesCompletionDatesCounter] = useState(0);

	const createHabit = async (currentSelectedHabit: IHabit) => {
		const habit = await HabitsService.createHabit(currentSelectedHabit, router);

		if (!habit) return;

		setHasHabitsBeenUpdated(true);
		setHabitNavbar(habit);
		setCompletionDates(habit.completionDates);
		setCompletionDatesCompletionDatesCounter(habit.completionDates.length);
		setCurrentSelectedHabit(habit);
	};

	const updateHabit = async (currentSelectedHabit: IHabit) => {
		await HabitsService.updateHabit(currentSelectedHabit, router);
		setHasHabitsBeenUpdated(true);
		setHabitNavbar(currentSelectedHabit);
		setCompletionDates(currentSelectedHabit.completionDates);
		setCompletionDatesCompletionDatesCounter(currentSelectedHabit.completionDates.length);
		setCurrentSelectedHabit(currentSelectedHabit);
	};

	const deleteHabit = async (currentSelectedHabit: IHabit | null) => {
		if (!currentSelectedHabit) return;
		await HabitsService.deleteHabit(currentSelectedHabit.habitId, router);
		updateMain(null, null, true);
	};

	const updateMain = async (habit: IHabit | null, currentSelectedHabit: IHabit | null, habitsUpdated = false) => {
		if (currentSelectedHabit && habitsUpdated) {
			const data = await HabitsService.updateHabit(currentSelectedHabit, router);
            if (!data) {
                handleOpenModal(ModalTypeEnum.GenericErrorModal);
                return;
            }
			setHasHabitsBeenUpdated(true);
		}

		if (!currentSelectedHabit && habitsMenu && habitsUpdated) {
			// Come from "All Habits" page. Update all the habits that have been changed.
			if (habitsMenu?.length > 1) {
                const data = await HabitsService.updateAllHabits(habitsMenu, router);
                if (!data) {
                    handleOpenModal(ModalTypeEnum.GenericErrorModal);
                    return;
                }
            }
			setHasHabitsBeenUpdated(true);
		}

		if (!habit) {
			setHabitNavbar(null);
			setCompletionDates([]);
			setCompletionDatesCompletionDatesCounter(0);
			if (habitsUpdated) setHasHabitsBeenUpdated(true);
			setCurrentSelectedHabit(null);
			return;
		}

		setHabitNavbar(habit);
		setCompletionDates(habit.completionDates);
		setCompletionDatesCompletionDatesCounter(habit.completionDates.length);
		setCurrentSelectedHabit(habit);
	};

	return (
		<div id="home" className="home">
            <Modal
				id="genericErrorModal"
				title="Error"
				body={<GenericErrorForm error="An error has occured." modalType={ModalTypeEnum.GenericErrorModal} onSubmit={onGenericErrorModalSubmit} onModalClose={handleCloseModal} />}
				showModal={showModal.GenericErrorModal}
				modalType={ModalTypeEnum.GenericErrorModal}
				onModalOpen={handleOpenModal}
				onModalClose={handleCloseModal}
			/>
			<Sidebar habitsMenu={habitsMenu} showSidebar={showSidebar} setShowSidebar={setShowSidebar} currentSelectedHabit={currentSelectedHabit} updateMain={updateMain} />
			<HabitsNavbar setShowSidebar={setShowSidebar} habit={habitNavbar} habitOps={{ createHabit, updateHabit, deleteHabit }} />
			{currentSelectedHabit && (
				<Calendar
					currentSelectedHabit={currentSelectedHabit}
					completionDatesCounter={completionDatesCounter}
					setCompletionDatesCompletionDatesCounter={setCompletionDatesCompletionDatesCounter}
					setCompletionDates={setCompletionDates}
					completionDates={completionDates}
				/>
			)}
			{!currentSelectedHabit &&
				habitsMenu?.map((habit, i) => (
					<div key={`calendar_${i}`}>
						<Calendar
							currentSelectedHabit={habit}
							completionDatesCounter={habit.completionDates ? habit.completionDates.length : 0}
							setCompletionDatesCompletionDatesCounter={setCompletionDatesCompletionDatesCounter}
							setCompletionDates={setCompletionDates}
							completionDates={habit.completionDates}
						/>
					</div>
				))}
			{showNoHabitsMessage && (
				<div className="d-flex justify-content-center align-items-center vh-100">
					<div className="no-habits text-center fw-bold">
						<p>No Habits here...</p>
						<p>Let&apos;s create some new habits by clicking the + icon in top right to track and achieve our goals together!</p>
					</div>
				</div>
			)}
		</div>
	);
};

export default Home;
