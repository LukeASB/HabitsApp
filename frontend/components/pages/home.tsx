import { useState, useEffect } from 'react';
import Calendar from "../calendar";
import Sidebar from "../sidebar";
import IHabit from '../../shared/interfaces/IHabit';
import { mockhabits } from '../../data/mock_habits';
import HabitsNavbar from '../habitsNavbar';

const Home: React.FC = () => {
    const [habitNavbar, setHabitNavbar] = useState<IHabit | null>(null);
    const [habitsMenu, setHabitsMenu] = useState<IHabit[]>([]);
    const [currentSelectedHabit, setCurrentSelectedHabit] = useState<IHabit | null>(null);

    useEffect(() => {
        console.log(`"http://localhost:8080/dohabitsapp/v1/retrievehabits`)
        const getHabits = async () => {
        try {
            // This will be called in every protected handler, will want moving to a global helper.
            // const csrfToken = sessionStorage.getItem("csrf-token");
            // const shortlivedJWTAccessToken = sessionStorage.getItem("access-token");

            // const res = await fetch("http://localhost:8080/dohabitsapp/v1/retrievehabits", { method: "GET", headers: { "Content-Type": "application/json", "X-CSRF-Token": csrfToken || "", "Authorization": `Bearer ${shortlivedJWTAccessToken || ""}` }});

            // if (res.status === 401) {
            //   if (retry < 1) {
            //     refresh(getHabits);
            //     retry++;
            //     return;
            //   } else {
            //     retry = 0;
            //   }
            // }

            // if (!res.ok) throw new Error("Network response error");
            // const data = await res.json();
            // console.log(data);
            setHabitsMenu(mockhabits);
        } catch (err) {
            console.log("Sidebar - Error occured:", err);

            setHabitsMenu([])
        }
        }

        const parseJwt = (token: string) => {
            try {
                const base64Url = token.split(".")[1];
                const base64 = base64Url.replace(/-/g, "+").replace(/_/g, "/");
                return JSON.parse(atob(base64));
            } catch (err) {
                console.error("Invalid JWT:", err);
                return null;
            }
        };

        // This will be used in multiple places. So wants to go in a global area...
        const refresh = async (callback: Function) => {
            try {
                const shortlivedJWTAccessToken = sessionStorage.getItem("access-token");
                const userData = shortlivedJWTAccessToken ? parseJwt(shortlivedJWTAccessToken) : null
                const res = await fetch("http://localhost:8080/dohabitsapp/v1/refresh", { method: "POST", body: JSON.stringify({ EmailAddress: userData.Email })});
                if (!res.ok) throw new Error("Network response error");
                const data = await res.json();
                console.log(data);
                sessionStorage.setItem("access-token", data.Token)
                callback();
            } catch (err) {
                console.log("Sidebar - Error occured:", err);
                // Redirect to Login
            }
        }

        getHabits();
        
    }, []); // Dependency - would need to be called whenever a habits added/removed.

    const [isCollapsed, setIsCollapsed] = useState(false);
    const toggleSidebar = () => setIsCollapsed(!isCollapsed);

    const [completionDates, setCompletionDates] = useState<string[]>([]);
    const [completionDatesCounter, setCompletionDatesCompletionDatesCounter] = useState(0);

    const updateMain = (habit: IHabit | null, habitsUpdated = false) => {
        if (!habit) {
            setHabitNavbar(null);
            setCurrentSelectedHabit(null);
            setCompletionDates([]);
            setCompletionDatesCompletionDatesCounter(0);
            habitsUpdated && setHabitsMenu(mockhabits); // This would become a dependency on useEffect to recall /retrievehabits
            return;
        }

        setHabitNavbar(habit);
        setCurrentSelectedHabit(habit);
        setCompletionDates(habit.completionDates);
        setCompletionDatesCompletionDatesCounter(habit.completionDates.length);
        habitsUpdated && setHabitsMenu(mockhabits); // This would become a dependency on useEffect to recall /retrievehabits
    };

    return (
        <div className="home">
            <div className={`d-flex ${isCollapsed ? 'sidebar-collapsed' : ''}`} style={ currentSelectedHabit || habitsMenu.length <= 1 ? { height: '90vh' } : {}}>
                <Sidebar habitsMenu={habitsMenu} toggleSidebar={toggleSidebar} isCollapsed={isCollapsed} updateMain={updateMain} />
                <div className="flex-grow-1">
                    <HabitsNavbar habit={habitNavbar} updateMain={updateMain}/>
                    {currentSelectedHabit && <Calendar currentSelectedHabit={currentSelectedHabit} completionDatesCounter={completionDatesCounter} setCompletionDatesCompletionDatesCounter={setCompletionDatesCompletionDatesCounter} setCompletionDates={setCompletionDates} completionDates={completionDates}/>}
                    {!currentSelectedHabit && habitsMenu.map((habit, i) => <div key={`calendar_${i}`}><Calendar currentSelectedHabit={habit} completionDatesCounter={habit.completionDates.length} setCompletionDatesCompletionDatesCounter={setCompletionDatesCompletionDatesCounter} setCompletionDates={setCompletionDates} completionDates={habit.completionDates}/></div>)}        
                </div>
            </div>
        </div>
    );
};

export default Home;
