import IHabit from "../shared/interfaces/IHabit";

export let mockhabits: IHabit[] = [{ id: "1", createdAt: Date.now(), name: "habit 1", days: 0, daysTarget: 30, numberOfDays: 1, completionDates: ["2024-12-01", "2024-12-02", "2024-12-03"]}, { id: "2", createdAt: Date.now(), name: "habit 2", days: 0, daysTarget: 30, numberOfDays: 1, completionDates: ["2024-12-20", "2024-12-02", "2024-12-03"],}]
export const createMockHabit = (habit: IHabit) => mockhabits.push(habit);
export const updateMockHabit = (habit: IHabit) => {
    for (let i = 0; i < mockhabits.length; i++) {
        if (mockhabits[i].id === habit.id) {
            mockhabits[i] = habit;
            break;
        }
    }
};
export const deleteMockHabit = (id: string) => mockhabits = mockhabits.filter(mockHabit => mockHabit.id !== id);