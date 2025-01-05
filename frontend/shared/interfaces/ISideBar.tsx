import { Dispatch, SetStateAction } from "react";
import IHabit from "./IHabit";

export default interface ISideBar {
    habitsMenu: IHabit[];
    showSidebar: boolean;
    setShowSidebar: Dispatch<SetStateAction<boolean>>;
    currentSelectedHabit: IHabit | null; 
    updateMain: (habit: IHabit | null, currentSelectedHabit: IHabit | null, habitsUpdated?: boolean) => Promise<void>;
}