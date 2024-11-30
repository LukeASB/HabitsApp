import { useState } from 'react';
import Calendar from "../calendar";
import Sidebar from "../sidebar";

const Home: React.FC = () => {
    const [isCollapsed, setIsCollapsed] = useState(false);
    const toggleSidebar = () => setIsCollapsed(!isCollapsed);

    const [mainHeader, setMainHeader] = useState("Main Content")
    const [mainContent, setMainContent] = useState("This is main content area...")

    const updateMain = (header: string, content: string) => {
        setMainHeader(header);
        setMainContent(content);
    }

    return (
    <div className="home">
        <div className={`d-flex ${isCollapsed ? 'sidebar-collapsed' : ''}`} style={{ height: '100vh' }}>
            <Sidebar toggleSidebar={toggleSidebar} isCollapsed={isCollapsed} updateMain={updateMain}/>
            {/* Main Content */}
            <div className="flex-grow-1 p-3">
                <h1>{mainHeader}</h1>
                <p>{mainContent}</p>
                <Calendar /> {/* Calendar needs to be controlled by state */}
            </div>
        </div>
    </div>);
}

export default Home;