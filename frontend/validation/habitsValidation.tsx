import IHabit from "../shared/interfaces/IHabit";
import IHabitFormError from "../shared/interfaces/IHabitFormError";

export class HabitsValidation {
    private static maxCharacterLength: number = 255;
    private static matchUpperLowerCaseLettersOnly = /[a-zA-Z ]/g;

    public static validateHabit(habit: Partial<IHabit>): Partial<IHabitFormError>[] {
        const errors: Partial<IHabitFormError>[] = [];

        const name = HabitsValidation.validateHabitName(habit.name)
        name && errors.push({ name: name });

        const days = HabitsValidation.validateHabitDays(habit.days);
        days && errors.push({ days: days });

        const daysTarget = HabitsValidation.validateHabitDays(habit.daysTarget);
        daysTarget && errors.push({ daysTarget: daysTarget })

        return errors;
    }

    private static validateHabitName(name?: string): string {
        if (name?.trim().length === 0) return "Habit name is required.";

        if (name && name.trim().length >= HabitsValidation.maxCharacterLength) return `Habit Name exceeds max character length of ${HabitsValidation.maxCharacterLength}`;

        if (name && name.matchAll(HabitsValidation.matchUpperLowerCaseLettersOnly)) return "Habit name is invalid.";

        return "";
    }

    private static validateHabitDays = (days?: number): string => HabitsValidation.validateDay(days);

    private static validateHabitDaysTarget = (days?: number): string => HabitsValidation.validateDay(days);

    private static validateDay(day?: number): string {
        if (day && day < 0) return "Habit Days cannot be less than 0";
    
        if (day && day >= 9999) return "Habit Days cannot be more than 9999 days";

        return "";
    }
}