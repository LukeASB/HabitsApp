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

	const [isCollapsed, setIsCollapsed] = useState(false);
	const toggleSidebar = () => setIsCollapsed(!isCollapsed);

	const [completionDates, setCompletionDates] = useState<string[]>([]);
	const [completionDatesCounter, setCompletionDatesCompletionDatesCounter] = useState(0);

    const createHabit = async (habit: IHabit) => {
        await HabitsService.createHabit(habit);
        setHasHabitsBeenUpdated(true);
        setHabitNavbar(habit);
		setCompletionDates(habit.completionDates);
		setCompletionDatesCompletionDatesCounter(habit.completionDates.length);
        setCurrentSelectedHabit(habit);
    };

    const updateHabit = async (habit: IHabit) => {
        const data = await HabitsService.updateHabit(habit);
        setHasHabitsBeenUpdated(true);
        setHabitNavbar(habit);
		setCompletionDates(habit.completionDates);
		setCompletionDatesCompletionDatesCounter(habit.completionDates.length);
        setCurrentSelectedHabit(habit);
    };

    const deleteHabit = async (habit: IHabit | null) => {
        if (!habit) return;
        await HabitsService.deleteHabit(habit.habitId);
        updateMain(null, null, true);
    };

	const updateMain = async (habit: IHabit | null, currentSelectedHabit: IHabit | null, habitsUpdated = false) => {
		if (!habit) {
			setHabitNavbar(null);
			setCompletionDates([]);
			setCompletionDatesCompletionDatesCounter(0);
            if (habitsUpdated) {
                setHasHabitsBeenUpdated(true);
            }
            setCurrentSelectedHabit(null);
			return;
		}

        if (currentSelectedHabit && habitsUpdated) {
            await HabitsService.updateHabit(currentSelectedHabit);
            setHasHabitsBeenUpdated(true);
        }

        if (!currentSelectedHabit && habitsUpdated) {
            // Come from "All Habits" page. Update all the habits that have been changed.
            for (const habit of habitsMenu) {
                await HabitsService.updateHabit(habit);
            }
            setHasHabitsBeenUpdated(true);
        }

        setHabitNavbar(habit);
		setCompletionDates(habit.completionDates);
		setCompletionDatesCompletionDatesCounter(habit.completionDates.length);
        setCurrentSelectedHabit(habit);
	};

	return (
		<div className="home">
			<div
				className={`d-flex ${isCollapsed ? "sidebar-collapsed" : ""}`}
				style={currentSelectedHabit || habitsMenu?.length <= 1 ? { height: "90vh" } : {}}
			>
				<Sidebar
					habitsMenu={habitsMenu}
					toggleSidebar={toggleSidebar}
					isCollapsed={isCollapsed}
                    currentSelectedHabit={currentSelectedHabit}
					updateMain={updateMain}
				/>
				<div className="flex-grow-1">
					<HabitsNavbar habit={habitNavbar} habitOps={{createHabit, updateHabit, deleteHabit}} />
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
			</div>
		</div>
	);
};

export default Home;
