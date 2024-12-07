import { useState, useEffect } from 'react';
import Link from 'next/link';
import IHabit from '../shared/interfaces/IHabit';

interface IHabitsNavbar {
  habit: IHabit | null;
}

/**
 * To Do:
 * - Tidy
 * - Hide Register/Login if user has session.
 * @returns
 */
const HabitsNavbar: React.FC<IHabitsNavbar> = ({ habit }) => {
    const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);

    useEffect(() => sessionStorage.getItem("access-token") ? setIsLoggedIn(true) : setIsLoggedIn(false), [])

    return (
        <nav className="navbar navbar-expand-lg navbar-light bg-primary">
            <div className="container content">
            <strong>
                <Link className="navbar-brand text-light" href="/">
                    {habit ? habit.name : "All Habits"}
                </Link>
            </strong>
            </div>
        <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
            <span className="navbar-toggler-icon"></span>
          </button>
        <div className="collapse navbar-collapse" id="navbarNav">
            <ul className="navbar-nav">
              <li className="nav-item active">
                {!isLoggedIn && (
                  <>
                    <Link className="navbar-brand text-light" href="/register">Register</Link>
                    <Link className="navbar-brand text-light" href="/login">Login</Link>
                  </>
                )}
              </li>
            </ul>
          </div>
        </nav>
    );
};

export default HabitsNavbar;