import { useState } from 'react';
import Calendar from "../calendar";
import Sidebar from "../sidebar";
import IHabit from '../../shared/interfaces/IHabit';

const Home: React.FC = () => {
    const [isCollapsed, setIsCollapsed] = useState(false);
    const toggleSidebar = () => setIsCollapsed(!isCollapsed);

    const [mainHeader, setMainHeader] = useState("Main Content");
    const [mainContent, setMainContent] = useState("This is main content area...");
    const [completionDates, setCompletionDates] = useState<string[]>([]);


    const updateMain = (habit: IHabit) => {
        setMainHeader(habit.name);
        setMainContent(habit.name);
        setCompletionDates(habit.completionDates);
        
        // Reset the selected Calendar Dates
        setSelectedDays({});
    };

    // State for Calendar
    const [selectedDays, setSelectedDays] = useState<{ [key: string]: number[] }>({});
    const [counter, setCounter] = useState(0);

    // Handler to update selected days
    const handleDaySelection = (year: number, month: number, day: number) => {
        const dateKey = `${year}-${month}`;
        const selectedForMonth = selectedDays[dateKey] || [];

        if (selectedForMonth.includes(day)) {
            // Remove the day
            const updatedDays = selectedForMonth.filter((selectedDay) => selectedDay !== day);
            setSelectedDays({ ...selectedDays, [dateKey]: updatedDays });
            setCounter(counter - 1);
        } else {
            // Add the day
            const updatedDays = [...selectedForMonth, day];
            setSelectedDays({ ...selectedDays, [dateKey]: updatedDays });
            setCounter(counter + 1);
        }
    };

    return (
        <div className="home">
            <div className={`d-flex ${isCollapsed ? 'sidebar-collapsed' : ''}`} style={{ height: '100vh' }}>
                <Sidebar toggleSidebar={toggleSidebar} isCollapsed={isCollapsed} updateMain={updateMain} />
                {/* Main Content */}
                <div className="flex-grow-1 p-3">
                    <h1>{mainHeader}</h1>
                    <p>{mainContent}</p>
                    <Calendar
                selectedDays={selectedDays}
                counter={0}
                handleDaySelection={handleDaySelection}
                completionDates={completionDates}
            />
                </div>
            </div>
        </div>
    );
};

export default Home;
