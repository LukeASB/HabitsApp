import IHabit from "../shared/interfaces/IHabit";

export let mockhabits: IHabit[] = [{ habitId: "1", createdAt: Date.now(), name: "habit 1", days: 3, daysTarget: 30, numberOfDays: 1, completionDates: ["2024-12-01", "2024-12-02", "2024-12-03"]}, { habitId: "2", createdAt: Date.now(), name: "habit 2", days: 3, daysTarget: 30, numberOfDays: 1, completionDates: ["2024-12-20", "2024-12-02", "2024-12-03"]}]
export const createMockHabit = (habit: IHabit) => {
    habit.habitId = mockhabits.length > 0 ? String(parseInt(mockhabits[mockhabits.length - 1]?.habitId) + 1) : "1",
	habit.createdAt = Date.now(),
    mockhabits.push(habit);

    return habit;
}
export const retrieveMockHabit = (habitId: string) => mockhabits.filter(mockhabit => mockhabit.habitId === habitId);
export const updateMockHabit = (habit: IHabit) => {
    for (let i = 0; i < mockhabits.length; i++) {
        if (mockhabits[i].habitId === habit.habitId) {
            mockhabits[i] = habit;
            return mockhabits[i];
        }
    }

    return habit;
};

export const updateAllMockHabits = (habits: IHabit[]) => {
    for (let i = 0; i < habits.length; i++) {
        for (let j = 0; j < mockhabits.length; j++) {
            if (habits[i].habitId === mockhabits[j].habitId) {
                mockhabits[j] = habits[i];
            }
        }
    }

    return habits;
};

export const deleteMockHabit = (habitId: string): boolean => {
    mockhabits = mockhabits.filter(mockHabit => mockHabit.habitId !== habitId);
    return true;
}