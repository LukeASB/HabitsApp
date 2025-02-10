export default interface ICreateHabitFormData {
    name: string;
    daysTarget: number;
    [key: string]: string | number;
};