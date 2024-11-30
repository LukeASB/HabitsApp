import { useState, useEffect } from 'react';

interface ISideBar {
  toggleSidebar: () => void;
  isCollapsed: boolean;
  updateMain: (header: string, content: string) => void;
}

// Move to data layer
interface IHabit {
  id: string;
  createdAt: number;
  name: string;
  days: number;
  daysTarget: number;
  numberOfDays: number;
}

// Move to mock_data
const habits: IHabit[] = [{ id: "1", createdAt: Date.now(), name: "habit 1", days: 0, daysTarget: 30, numberOfDays: 1 }, { id: "1", createdAt: Date.now(), name: "habit 2", days: 0, daysTarget: 30, numberOfDays: 1 }]

const Sidebar: React.FC<ISideBar> = ({ toggleSidebar, isCollapsed, updateMain }) => {
  const [habitsMenu, setHabitsMenu] = useState<IHabit[]>([]);
  const [renderHabitsMenu, setRenderHabitsMenu] = useState(!isCollapsed);

  // let retry = 0;

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
        setHabitsMenu(habits);
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
    
  }, []);

  useEffect(() => {
    if (isCollapsed) return setRenderHabitsMenu(false);

    const timer = setTimeout(() => setRenderHabitsMenu(true), 100);
    return () => clearTimeout(timer);
  }, [isCollapsed]);

  return (
    <div className={`bg-primary text-white p-3 ${isCollapsed ? 'collapsed-width' : 'full-width'}`} style={{width: isCollapsed ? '60px' : '250px', transition: 'width 0.s'}}>
      <div className="d-flex justify-content-between align-items-center">
        <h5 className={`${isCollapsed ? 'd-none' : ''}`}>Habits</h5>
          <button className="btn btn-sm btn-light" onClick={toggleSidebar} aria-label="Toggle Sidebar">{isCollapsed ? '→' : '←'}</button>
      </div>
      <ul className="nav flex-column mt-3">
        {renderHabitsMenu && habitsMenu.map((habit: IHabit, i) => {
          return (
            <li key={`${habit.name}_${i}`} className={`${isCollapsed ? 'd-none' : 'nav-item '}`}>
              <button className="btn btn-link nav-link text-white" onClick={() => updateMain(habit.name, habit.name)}><i className="bi bi-info-circle" /> {habit.name}</button>
            </li>
            );
          })
        }
      </ul>
    </div>
  );
};

export default Sidebar;
