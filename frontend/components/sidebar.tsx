import { useState, useEffect } from 'react';
import IHabit from '../shared/interfaces/IHabit';


interface ISideBar {
  habitsMenu: IHabit[];
  toggleSidebar: () => void;
  isCollapsed: boolean;
  updateMain: (habit: IHabit) => void;
}

// Move to mock_data
const Sidebar: React.FC<ISideBar> = ({ habitsMenu, toggleSidebar, isCollapsed, updateMain }) => {
  const [renderHabitsMenu, setRenderHabitsMenu] = useState(!isCollapsed);

  useEffect(() => {
    if (isCollapsed) return setRenderHabitsMenu(false);

    const timer = setTimeout(() => setRenderHabitsMenu(true), 100);
    return () => clearTimeout(timer);
  }, [isCollapsed]);

  return (
    <div className={`bg-primary text-white p-3 ${isCollapsed ? 'collapsed-width' : 'full-width'}`} style={{width: isCollapsed ? '60px' : '250px', transition: 'width 0.s'}}>
      <div className="d-flex justify-content-between align-items-center">
        <h5 className={`${isCollapsed ? 'd-none' : ''}`}>Habits</h5>
          <button className="btn btn-sm btn-light" onClick={toggleSidebar} aria-label="Toggle Sidebar">{isCollapsed ? '→' : '←'}</button>
      </div>
      <ul className="nav flex-column mt-3">
        {renderHabitsMenu && habitsMenu.map((habit: IHabit, i) => {
          return (
            <li key={`${habit.name}_${i}`} className={`${isCollapsed ? 'd-none' : 'nav-item '}`}>
              <button className="btn btn-link nav-link text-white" onClick={() => updateMain(habit)}><i className="bi bi-info-circle" /> {habit.name}</button>
            </li>
            );
          })
        }
      </ul>
    </div>
  );
};

export default Sidebar;
