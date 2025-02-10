import { Dispatch, SetStateAction } from "react";
import IHabit from "./IHabit";

export default interface IHabitsNavbar {
    setShowSidebar: Dispatch<SetStateAction<boolean>>;
    habit: IHabit | null;
    habitOps: {
        createHabit: (habit: IHabit) => Promise<void>;
        updateHabit: (habit: IHabit) => Promise<void>;
        deleteHabit: (habit: IHabit | null) => Promise<void>;
    }
}

