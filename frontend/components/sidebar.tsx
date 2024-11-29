import { useState } from 'react';
import Link from 'next/link';

interface ISideBar {
  toggleSidebar: () => void;
  isCollapsed: boolean;
  updateMain: (header: string, content: string) => void;
}

const Sidebar: React.FC<ISideBar> = ({ toggleSidebar, isCollapsed, updateMain }) => {

  return (
    <div className={`bg-primary text-white p-3 ${isCollapsed ? 'collapsed-width' : 'full-width'}`} style={{width: isCollapsed ? '60px' : '250px', transition: 'width 0.3s'}}>
      <div className="d-flex justify-content-between align-items-center">
        <h5 className={`${isCollapsed ? 'd-none' : ''}`}>Habits</h5>
          <button className="btn btn-sm btn-light" onClick={toggleSidebar} aria-label="Toggle Sidebar">{isCollapsed ? '→' : '←'}</button>
      </div>
      <ul className="nav flex-column mt-3">
        <li className="nav-item">
          <button className="btn btn-link nav-link text-white" onClick={() => updateMain("Home", "Home from sidebar.tsx")}><i className="bi bi-info-circle" /> Home</button>
        </li>
        <li className="nav-item">
          <button className="btn btn-link nav-link text-white" onClick={() => updateMain("About", "About from sidebar.tsx")}><i className="bi bi-info-circle" /> About</button>
        </li>
        <li className="nav-item">
          <button className="btn btn-link nav-link text-white" onClick={() => updateMain("Contact", "Contact from sidebar.tsx")}><i className="bi bi-envelope" /> Contact</button>
          </li>
      </ul>
    </div>
  );
};

export default Sidebar;
