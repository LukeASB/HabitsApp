import IHabit from "../shared/interfaces/IHabit";

export const mockhabits: IHabit[] = [{ id: "1", createdAt: Date.now(), name: "habit 1", days: 0, daysTarget: 30, numberOfDays: 1, completionDates: ["2024-12-01", "2024-12-02", "2024-12-03"]}, { id: "1", createdAt: Date.now(), name: "habit 2", days: 0, daysTarget: 30, numberOfDays: 1, completionDates: ["2024-12-20", "2024-12-02", "2024-12-03"],}]