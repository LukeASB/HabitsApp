import { useState } from 'react';
import { ReactNode } from "react";
// import Banner from "./sidebar";
import Navbar from "./navbar";
import Link from "next/link";
import Footer from "./footer";
import Sidebar from './sidebar';
// import Sidebar from "./sidebar";

interface IMainLayout {
    children?: ReactNode
}

const MainLayout: React.FC<IMainLayout> = ( { children } ) => {
    const [isCollapsed, setIsCollapsed] = useState(false);

const toggleSidebar = () => {
  setIsCollapsed(!isCollapsed);
};

    return (
        <>
        <Navbar />
        <main>
        <div className={`d-flex ${isCollapsed ? 'sidebar-collapsed' : ''}`} style={{ height: '100vh' }}>
            <Sidebar toggleSidebar={toggleSidebar} isCollapsed={isCollapsed}/>
            {/* Main Content */}
            <div className="flex-grow-1 p-3">
                <h1>Main Content</h1>
                <p>This is the main content area.</p>
            </div>
        </div>
        </main>
        <Footer />
        </>
    );
};

export default MainLayout;