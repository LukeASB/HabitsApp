import { useState } from 'react';
import { ReactNode } from "react";
// import Banner from "./sidebar";
import Navbar from "./navbar";
import Link from "next/link";
import Footer from "./footer";
import Sidebar from './sidebar';
import Calendar from './calendar';
// import Sidebar from "./sidebar";

interface IMainLayout {
    children?: ReactNode
}

const MainLayout: React.FC<IMainLayout> = ( { children } ) => {
    const [isCollapsed, setIsCollapsed] = useState(false);
    const toggleSidebar = () => setIsCollapsed(!isCollapsed);

    const [mainHeader, setMainHeader] = useState("Main Content")
    const [mainContent, setMainContent] = useState("This is main content area...")

    const updateMain = (header: string, content: string) => {
        setMainHeader(header);
        setMainContent(content);
    }

    return (
        <>
        <Navbar />
        <main>
        <div className={`d-flex ${isCollapsed ? 'sidebar-collapsed' : ''}`} style={{ height: '100vh' }}>
            <Sidebar toggleSidebar={toggleSidebar} isCollapsed={isCollapsed} updateMain={updateMain}/>
            {/* Main Content */}
            <div className="flex-grow-1 p-3">
                <h1>{mainHeader}</h1>
                <p>{mainContent}</p>
                <Calendar />
            </div>
        </div>
        </main>
        <Footer />
        </>
    );
};

export default MainLayout;