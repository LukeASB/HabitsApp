import { useState, useEffect } from 'react';
import IHabit from '../shared/interfaces/IHabit';
import ISideBar from '../shared/interfaces/ISideBar';

const Sidebar: React.FC<ISideBar> = ({ habitsMenu, toggleSidebar, isCollapsed, updateMain }) => {
    const [renderHabitsMenu, setRenderHabitsMenu] = useState(!isCollapsed);

    useEffect(() => {
        if (isCollapsed) return setRenderHabitsMenu(false);

        const timer = setTimeout(() => setRenderHabitsMenu(true), 100);
        return () => clearTimeout(timer);
    }, [isCollapsed]);

    return (
        <div className={`bg-dark text-white p-3 ${isCollapsed ? 'collapsed-width' : 'full-width'}`} style={{width: isCollapsed ? '60px' : '250px', transition: 'width 0.s'}}>
            <div className="d-flex justify-content-between align-items-center">
                <h5 className={`${isCollapsed ? 'd-none' : ''}`}>Habits</h5>
                <button className="btn btn-sm btn-light" onClick={toggleSidebar} aria-label="Toggle Sidebar">{isCollapsed ? '→' : '←'}</button>
            </div>
            <ul className="nav flex-column mt-3">
                <li key={`all_habits`} className={`${isCollapsed ? 'd-none' : 'nav-item '}`}>
                    <button className="btn btn-link nav-link text-white" onClick={() => updateMain(null)}><i className="habit" />All Habits</button>
                </li>
                {renderHabitsMenu && habitsMenu.map((habit: IHabit, i) => {
                    return (
                        <li key={`${habit.name}_${i}`} className={`${isCollapsed ? 'd-none' : 'nav-item '}`}>
                            <button className="btn btn-link nav-link text-white" onClick={() => updateMain(habit)}><i className={`${habit.name}_${i}`} /> {habit.name}</button>
                        </li>
                    );
                })}
            </ul>
        </div>
        );
    };

export default Sidebar;
