import { useState, useEffect } from 'react';
import Link from 'next/link';
import IHabit from '../shared/interfaces/IHabit';
import { mockhabits, createMockHabit, updateMockHabit, deleteMockHabit } from '../data/mock_habits';
import Modal from './modal/modal';
import HabitsButtons from './habitButtons';

interface IHabitsNavbar {
  habit: IHabit | null;
  updateMain: (habit: IHabit | null, habitsUpdated?: boolean) => void;
}

/**
 * @returns
 */
const HabitsNavbar: React.FC<IHabitsNavbar> = ({ habit, updateMain}) => {
    const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);
    const test = 'debug';
    useEffect(() => sessionStorage.getItem("access-token") || test ? setIsLoggedIn(true) : setIsLoggedIn(false), []);

    const createHabit = (habit: IHabit) => {
        createMockHabit(habit);
        updateMain(habit, true);
    };

    const updateHabit = (habit: IHabit) => {
      updateMockHabit(habit);
      updateMain(habit, true);
    };

    const deleteHabit = (habit: IHabit | null) => {
        if (!habit) return;
        habit.id && deleteMockHabit(habit.id);
        updateMain(null, true);
    };

    return (
        <nav className="navbar navbar-expand-lg navbar-light bg-primary">
            <div className="container content">
            <strong>
                <Link className="navbar-brand text-light" href="/">{habit ? habit.name : "All Habits"}</Link>
            </strong>
            </div>
        <div className="navbar" id="navbarNav">
            <ul className="navbar-nav">
              <li className="nav-item active">
                {isLoggedIn && (
                      <div className="d-flex gap-2">
                        <button
                          type="button"
                          className="btn btn-primary btn-custom robo popup-trigger popmake-680"
                          data-popup-id="680"
                          data-do-default="0"
                          data-bs-toggle="modal"
                          data-bs-target={/*`#${card.readMoreModal.id}`*/`#addHabit`}
                        >
                          Test
                        </button>
                        <Modal
                          id={`addHabit`}
                          title={"test modal 123"}
                          body={<div>"test body 123"</div>}
                        />
                      {/* Plus Icon Button */}
                      <HabitsButtons 
                        icon="plus"
                        habit={{ 
                          id: mockhabits.length > 0 ? String(parseInt(mockhabits[mockhabits.length-1]?.id)+1) : "1",
                          createdAt: Date.now(),
                          name: mockhabits.length > 0 ? `Added Habit_${String(parseInt(mockhabits[mockhabits.length-1].id)+1)}` : "Added Habit_1",
                          days: 0,
                          daysTarget: 30,
                          numberOfDays: 1,
                          completionDates: ["2024-12-20", "2024-12-19", "2024-12-17"]
                        }}
                        onClick={createHabit}/>
                      {/* Update Icon Button */}
                      {habit && <HabitsButtons
                        icon="gear"
                        habit={{ id: habit.id, createdAt: Date.now(), name: `Updated Habit_${habit.id}`, days: 0, daysTarget: 30, numberOfDays: 1, completionDates: ["2024-12-01", "2024-12-19", "2024-12-14"]}}
                        onClick={updateHabit}
                      />}
                      {/* X Icon Button */}
                      {habit && <HabitsButtons 
                        icon="x"
                        habit={habit}
                        onClick={deleteHabit}
                      />}
                    </div>
                )}
              </li>
            </ul>
          </div>
        </nav>
    );
};

export default HabitsNavbar;