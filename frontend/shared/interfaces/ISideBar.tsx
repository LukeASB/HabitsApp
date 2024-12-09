import IHabit from "./IHabit";

export default interface ISideBar {
    habitsMenu: IHabit[];
    toggleSidebar: () => void;
    isCollapsed: boolean;
    updateMain: (habit: IHabit | null) => void;
}