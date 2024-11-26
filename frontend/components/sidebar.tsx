import { useState } from 'react';
import Link from 'next/link';

interface ISideBar {
  toggleSidebar: () => void;
  isCollapsed: boolean;
}

const Sidebar: React.FC<ISideBar> = ({ toggleSidebar, isCollapsed }) => {

  return (
    <div
    className={`bg-primary text-white p-3 ${isCollapsed ? 'collapsed-width' : 'full-width'}`}
    style={{
    width: isCollapsed ? '60px' : '250px',
    transition: 'width 0.3s',
    }}
    >
    <div className="d-flex justify-content-between align-items-center">
    <h5 className={`${isCollapsed ? 'd-none' : ''}`}>Habits</h5>
    <button
        className="btn btn-sm btn-light"
        onClick={toggleSidebar}
        aria-label="Toggle Sidebar"
    >
        {isCollapsed ? '→' : '←'}
    </button>
    </div>
    <ul className="nav flex-column mt-3">
    <li className="nav-item">
        <Link href="/" className="nav-link text-white">
        <i className="bi bi-house" /> <span className={`${isCollapsed ? 'd-none' : ''}`}>Home</span>
        </Link>
    </li>
    <li className="nav-item">
        <Link href="/about" className="nav-link text-white">
        <i className="bi bi-info-circle" /> <span className={`${isCollapsed ? 'd-none' : ''}`}>About</span>
        </Link>
    </li>
    <li className="nav-item">
        <Link href="/contact" className="nav-link text-white">
        <i className="bi bi-envelope" /> <span className={`${isCollapsed ? 'd-none' : ''}`}>Contact</span>
        </Link>
    </li>
    </ul>

    </div>
  );
};

export default Sidebar;
