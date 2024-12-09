import { MouseEventHandler } from "react";
import IHabit from "../shared/interfaces/IHabit";

interface IHabitsButton {
    icon: string;
    habit: IHabit;
    onClick: (habit: IHabit) => void;
}

const HabitsButtons: React.FC<IHabitsButton> = ({ icon, habit, onClick}) => {
    return (
        <button className="btn btn-dark" onClick={() => onClick(habit)}>
            <i className={`bi bi-${icon}`}></i>
        </button>
    );
}

export default HabitsButtons;