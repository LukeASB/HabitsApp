import IHabit from "../shared/interfaces/IHabit";
import IHabitFormError from "../shared/interfaces/IHabitFormError";
import { HabitsValidation } from "../validation/habitsValidation";

export class HabitsModel {
	public static processHabit(habit: Partial<IHabit>): Partial<IHabitFormError>[] {
        const validationErrors = HabitsValidation.validateHabit(habit);
		return validationErrors;
	}
}
