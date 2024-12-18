import { useState, useEffect } from "react";
import Link from "next/link";
import IHabit from "../shared/interfaces/IHabit";
import HabitsButtons from "./habitButtons";
import CreateHabitForm from "./forms/createHabitForm";
import UpdateHabitForm from "./forms/updateHabitForm";
import DeleteHabitForm from "./forms/deleteHabitForm";
import IHabitsNavbar from "../shared/interfaces/IHabitsNavbar";
import { HabitsService } from "../services/habitsService";

/**
 * @returns
 */
const HabitsNavbar: React.FC<IHabitsNavbar> = ({ habit, updateMain }) => {
	const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);
	const test = "debug";
	useEffect(() => (sessionStorage.getItem("access-token") || test ? setIsLoggedIn(true) : setIsLoggedIn(false)), []);

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
        const data = await HabitsService.deleteHabit(habit.id);
		updateMain(null, true);
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
										body: <CreateHabitForm onSubmit={createHabit} />,
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
													habit={{
														id: habit.id,
														createdAt: Date.now(),
														name: `Updated Habit_${habit.id}`,
														days: 3,
														daysTarget: 30,
														numberOfDays: 1,
														completionDates: ["2024-12-01", "2024-12-19", "2024-12-14"],
													}}
													onSubmit={updateHabit}
												/>
											),
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
													modalId={"deleteHabitModal"}
													onSubmit={deleteHabit}
												/>
											),
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
