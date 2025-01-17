export default interface IHabit {
    habitId: string;
    createdAt: number;
    name: string;
    days: number;
    daysTarget: number;
    numberOfDays: number;
    completionDates: string[];
}
  