import { useState, useEffect } from "react";
import Calendar from "../calendar";
import Sidebar from "../sidebar";
import IHabit from "../../shared/interfaces/IHabit";
import { mockhabits } from "../../data/mock_habits";
import HabitsNavbar from "../habitsNavbar";
import { HabitsService } from "../../services/habitsService";

const Home: React.FC = () => {
	const [habitNavbar, setHabitNavbar] = useState<IHabit | null>(null);
	const [habitsMenu, setHabitsMenu] = useState<IHabit[]>([]);
	const [currentSelectedHabit, setCurrentSelectedHabit] =
		useState<IHabit | null>(null);

	useEffect(() => {
		const retrieveHabits = async () => {
			const data = await HabitsService.retrieveHabits();
			setHabitsMenu(data);
		};

		retrieveHabits();
	}, []); // Dependency - would need to be called whenever a habits added/removed.

	const [isCollapsed, setIsCollapsed] = useState(false);
	const toggleSidebar = () => setIsCollapsed(!isCollapsed);

	const [completionDates, setCompletionDates] = useState<string[]>([]);
	const [completionDatesCounter, setCompletionDatesCompletionDatesCounter] =
		useState(0);

	const updateMain = (habit: IHabit | null, habitsUpdated = false) => {
		if (!habit) {
			setHabitNavbar(null);
			setCurrentSelectedHabit(null);
			setCompletionDates([]);
			setCompletionDatesCompletionDatesCounter(0);
			habitsUpdated && setHabitsMenu(mockhabits); // This would become a dependency on useEffect to recall /retrievehabits
			return;
		}

		setHabitNavbar(habit);
		setCurrentSelectedHabit(habit);
		setCompletionDates(habit.completionDates);
		setCompletionDatesCompletionDatesCounter(habit.completionDates.length);
		habitsUpdated && setHabitsMenu(mockhabits); // This would become a dependency on useEffect to recall /retrievehabits
	};

	return (
		<div className="home">
			<div
				className={`d-flex ${isCollapsed ? "sidebar-collapsed" : ""}`}
				style={
					currentSelectedHabit || habitsMenu.length <= 1
						? { height: "90vh" }
						: {}
				}
			>
				<Sidebar
					habitsMenu={habitsMenu}
					toggleSidebar={toggleSidebar}
					isCollapsed={isCollapsed}
					updateMain={updateMain}
				/>
				<div className="flex-grow-1">
					<HabitsNavbar habit={habitNavbar} updateMain={updateMain} />
					{currentSelectedHabit && (
						<Calendar
							currentSelectedHabit={currentSelectedHabit}
							completionDatesCounter={completionDatesCounter}
							setCompletionDatesCompletionDatesCounter={
								setCompletionDatesCompletionDatesCounter
							}
							setCompletionDates={setCompletionDates}
							completionDates={completionDates}
						/>
					)}
					{!currentSelectedHabit &&
						habitsMenu.map((habit, i) => (
							<div key={`calendar_${i}`}>
								<Calendar
									currentSelectedHabit={habit}
									completionDatesCounter={
										habit.completionDates.length
									}
									setCompletionDatesCompletionDatesCounter={
										setCompletionDatesCompletionDatesCounter
									}
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
