import { ErrorsEnum } from "../shared/enum/errorsEnum";
import ICreateHabitFormError from "../shared/interfaces/ICreateHabitFormError";
import IHabit from "../shared/interfaces/IHabit";
import IUpdateHabitFormError from "../shared/interfaces/IUpdateHabitFormError";

export class HabitsValidation {
	private static readonly minHabitsDays: number = 0;
	private static readonly maxHabitsDays: number = 9999;
	private static readonly maxCharacterLength: number = 255;
	private static readonly matchLettersNumbersAndColonOnly = /^[a-zA-Z0-9: ]*$/;

	public static validateCreateHabit(habit: Partial<IHabit>): Partial<ICreateHabitFormError>[] {
		const errors: Partial<ICreateHabitFormError>[] = [];

		const name = HabitsValidation.validateHabitName(habit.name);
		name && errors.push({ name: name });

		const daysTarget = HabitsValidation.validateHabitDays(habit.daysTarget);
		daysTarget && errors.push({ daysTarget: daysTarget });

		return errors;
	}

	public static validateUpdateHabit(habit: Partial<IHabit>): Partial<IUpdateHabitFormError>[] {
		const errors: Partial<IUpdateHabitFormError>[] = [];

		const name = HabitsValidation.validateHabitName(habit.name);
		name && errors.push({ name: name });

		const daysTarget = HabitsValidation.validateHabitDays(habit.daysTarget);
		daysTarget && errors.push({ daysTarget: daysTarget });

		return errors;
	}

	private static validateHabitName(name?: string): string {
		if (name?.trim().length === 0) return ErrorsEnum.Required;
		if (name && name.trim().length >= HabitsValidation.maxCharacterLength) return ErrorsEnum.HabitNameMax.replace("{0}", `${HabitsValidation.maxCharacterLength}`);
		if (name && !HabitsValidation.matchLettersNumbersAndColonOnly.test(name)) return ErrorsEnum.Invalid;

		return "";
	}

	private static validateHabitDays = (days?: number): string => HabitsValidation.validateDay(days);

	private static validateDay(day?: number): string {
		if (day && day < 0) return ErrorsEnum.HabitDaysMin.replace("{0}", `${HabitsValidation.minHabitsDays}`);
		if (day && day >= 9999) return ErrorsEnum.HabitDaysMax.replace("{0}", `${HabitsValidation.maxHabitsDays}`);

		return "";
	}
}
