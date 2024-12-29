import IHabit from "./IHabit";

export default interface ISideBar {
    habitsMenu: IHabit[];
    toggleSidebar: () => void;
    isCollapsed: boolean;
    currentSelectedHabit: IHabit | null; 
    updateMain: (habit: IHabit | null, currentSelectedHabit: IHabit | null, habitsUpdated?: boolean) => Promise<void>;
}