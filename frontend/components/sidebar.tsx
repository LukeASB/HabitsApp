import { useState } from "react";
import IHabit from "../shared/interfaces/IHabit";
import ISideBar from "../shared/interfaces/ISideBar";

const Sidebar: React.FC<ISideBar> = ({ habitsMenu, showSidebar, setShowSidebar, currentSelectedHabit, updateMain }) => {
	const handleHabitClick = (habit: IHabit | null) => {
		updateMain(habit, currentSelectedHabit, true);
		setShowSidebar(false);
	};

	return (
		<div id="sidebar">
			<div
				className={`offcanvas offcanvas-start bg-dark text-white ${showSidebar ? "show d-block" : ""}`}
				id="offcanvasSidebar"
				aria-labelledby="offcanvasSidebarLabel"
				tabIndex={-1}
				style={{ visibility: showSidebar ? "visible" : "hidden" }}
			>
				<div className="offcanvas-header">
					<h5 className="offcanvas-title" id="offcanvasSidebarLabel">
						Habits
					</h5>
					<button type="button" className="btn-close text-reset" onClick={() => setShowSidebar(false)} aria-label="Close"></button>
				</div>

				<div className="offcanvas-body">
					<ul className="nav flex-column">
						<li key={`all_habits`} className="nav-item">
							<button className="btn btn-link nav-link text-white" onClick={() => handleHabitClick(null)}>
								<i className="habit" /> All Habits
							</button>
						</li>
						{habitsMenu?.map((habit: IHabit, i) => {
							return (
								<li key={`${habit.name}_${i}`} className="nav-item">
									<button className="btn btn-link nav-link text-white" onClick={() => handleHabitClick(habit)}>
										<i className={`${habit.name}_${i}`} /> {habit.name}
									</button>
								</li>
							);
						})}
					</ul>
				</div>
			</div>
		</div>
	);
};

export default Sidebar;
