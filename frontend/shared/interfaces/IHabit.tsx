export default interface IHabit {
    id: string;
    createdAt: number;
    name: string;
    days: number;
    daysTarget: number;
    numberOfDays: number;
    completionDates: string[];
}
  