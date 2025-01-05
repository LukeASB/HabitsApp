import { useState, useEffect } from "react";
import Calendar from "../calendar";
import Sidebar from "../sidebar";
import IHabit from "../../shared/interfaces/IHabit";
import { mockhabits } from "../../data/mock_habits";
import HabitsNavbar from "../habitsNavbar";
import { HabitsService } from "../../services/habitsService";

const Home: React.FC = () => {
	const [hasHabitsBeenUpdated, setHasHabitsBeenUpdated] = useState<boolean>(true);
	const [habitNavbar, setHabitNavbar] = useState<IHabit | null>(null);
	const [habitsMenu, setHabitsMenu] = useState<IHabit[]>([]);
	const [currentSelectedHabit, setCurrentSelectedHabit] = useState<IHabit | null>(null);

	useEffect(() => {
		if (!hasHabitsBeenUpdated) return;
		if (process.env.ENVIRONMENT === "DEV") {
			setHabitsMenu(mockhabits);
			setHasHabitsBeenUpdated(false);
			return;
		}

		const retrieveHabits = async () => setHabitsMenu(await HabitsService.retrieveHabits());

		retrieveHabits();
		setHasHabitsBeenUpdated(false);
	}, [hasHabitsBeenUpdated]);

	// Sidebar State
	const [showSidebar, setShowSidebar] = useState(false);

	// Calendar State
	const [completionDates, setCompletionDates] = useState<string[]>([]);
	const [completionDatesCounter, setCompletionDatesCompletionDatesCounter] = useState(0);

	const createHabit = async (currentSelectedHabit: IHabit) => {
		await HabitsService.createHabit(currentSelectedHabit);
		setHasHabitsBeenUpdated(true);
		setHabitNavbar(currentSelectedHabit);
		setCompletionDates(currentSelectedHabit.completionDates);
		setCompletionDatesCompletionDatesCounter(currentSelectedHabit.completionDates.length);
		setCurrentSelectedHabit(currentSelectedHabit);
	};

	const updateHabit = async (currentSelectedHabit: IHabit) => {
		await HabitsService.updateHabit(currentSelectedHabit);
		setHasHabitsBeenUpdated(true);
		setHabitNavbar(currentSelectedHabit);
		setCompletionDates(currentSelectedHabit.completionDates);
		setCompletionDatesCompletionDatesCounter(currentSelectedHabit.completionDates.length);
		setCurrentSelectedHabit(currentSelectedHabit);
	};

	const deleteHabit = async (currentSelectedHabit: IHabit | null) => {
		if (!currentSelectedHabit) return;
		await HabitsService.deleteHabit(currentSelectedHabit.habitId);
		updateMain(null, null, true);
	};

	const updateMain = async (habit: IHabit | null, currentSelectedHabit: IHabit | null, habitsUpdated = false) => {
		if (currentSelectedHabit && habitsUpdated) {
			await HabitsService.updateHabit(currentSelectedHabit);
			setHasHabitsBeenUpdated(true);
		}

		if (!currentSelectedHabit && habitsUpdated) {
			// Come from "All Habits" page. Update all the habits that have been changed.
			for (const habit of habitsMenu) await HabitsService.updateHabit(habit);
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
			<Sidebar habitsMenu={habitsMenu} showSidebar={showSidebar} setShowSidebar={setShowSidebar} currentSelectedHabit={currentSelectedHabit} updateMain={updateMain} />
			<HabitsNavbar showSidebar={showSidebar} setShowSidebar={setShowSidebar} habit={habitNavbar} habitOps={{ createHabit, updateHabit, deleteHabit }} />
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
		</div>
	);
};

export default Home;
