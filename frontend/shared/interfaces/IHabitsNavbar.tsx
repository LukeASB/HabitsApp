import IHabit from "./IHabit";

export default interface IHabitsNavbar {
    habit: IHabit | null;
    habitOps: {
        createHabit: (habit: IHabit) => Promise<void>;
        updateHabit: (habit: IHabit) => Promise<void>;
        deleteHabit: (habit: IHabit | null) => Promise<void>;
    }
}

