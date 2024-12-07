import { useState, useEffect } from 'react';
import Link from 'next/link';
import IHabit from '../shared/interfaces/IHabit';

/**
 * @returns
 */
const Navbar: React.FC = () => {
    const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);
    const test = 'debug';

    useEffect(() => sessionStorage.getItem("access-token") || test ? setIsLoggedIn(true) : setIsLoggedIn(false), [])

    const IconButtons = () => {
      return (
        <div className="d-flex gap-2">
          {/* Plus Icon Button */}
          <button className="btn btn-dark">
            <i className="bi bi-plus"></i>
          </button>
    
          {/* X Icon Button */}
          <button className="btn btn-dark">
            <i className="bi bi-x"></i>
          </button>
        </div>
      );
    };

    return (
        <nav className="navbar navbar-expand-lg navbar-light bg-dark">
            <div className="container content">
            <strong>
                <span className="navbar-brand text-light">
                  <Link className="navbar-brand text-light" href="/">Habits Apps</Link>
                </span>
            </strong>
            </div>
        <div className="navbar" id="navbarNav">
            <ul className="navbar-nav">
              <li className="nav-item">

                {isLoggedIn && IconButtons()}
              </li>
            </ul>
          </div>
        </nav>
    );
};

export default Navbar;