export default interface IUpdateHabitFormData {
    name: string;
    daysTarget: number;
    [key: string]: string | number;
};