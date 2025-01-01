import ICreateHabitFormError from "../shared/interfaces/ICreateHabitFormError";
import IHabit from "../shared/interfaces/IHabit";
import IUpdateHabitFormError from "../shared/interfaces/IUpdateHabitFormError";
import { HabitsValidation } from "../validation/habitsValidation";

export class HabitsModel {
	public static processCreateHabit = (habit: Partial<IHabit>): Partial<ICreateHabitFormError>[] => HabitsValidation.validateCreateHabit(habit);
	public static processUpdateHabit = (habit: Partial<IHabit>): Partial<IUpdateHabitFormError>[] => HabitsValidation.validateUpdateHabit(habit);
}
