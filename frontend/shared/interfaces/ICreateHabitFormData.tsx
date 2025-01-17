export default interface ICreateHabitFormData {
    name: string;
    days: number;
    daysTarget: number;
    [key: string]: string | number;
};