import IHabit from "./IHabit";

export default interface IHabitsNavbar {
    habit: IHabit | null;
    updateMain: (habit: IHabit | null, habitsUpdated?: boolean) => void;
}