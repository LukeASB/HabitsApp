import { Dispatch, SetStateAction } from "react";
import IHabit from "./IHabit";

export default interface IHabitsNavbar {
    showSidebar: boolean;
    setShowSidebar: Dispatch<SetStateAction<boolean>>;
    habit: IHabit | null;
    habitOps: {
        createHabit: (habit: IHabit) => Promise<void>;
        updateHabit: (habit: IHabit) => Promise<void>;
        deleteHabit: (habit: IHabit | null) => Promise<void>;
    }
}

