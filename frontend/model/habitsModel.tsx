import IHabit from "../shared/interfaces/IHabit";

export class HabitsModel {
	public static validateHabit(habit: IHabit): string[] {
		const errors: string[] = [];

		if (!habit.name || habit.name.trim().length === 0) {
			errors.push("Habit name is required.");
		}

		if (!Array.isArray(habit.days) || habit.days.length === 0) {
			errors.push("At least one day should be selected.");
		}

		if (habit.daysTarget <= 0) {
			errors.push("Days target must be a positive number.");
		}

		return errors;
	}

	public static createHabit(habit: IHabit): IHabit {
		// Additional business logic, e.g., generating IDs, etc.
		return habit;
	}

	public static updateHabit(
		existingHabit: IHabit,
		updates: Partial<IHabit>,
	): IHabit {
		// Update the habit fields while maintaining existing ones
		return { ...existingHabit, ...updates };
	}

	public static deleteHabit(habitId: string): string {
		// Logic for deleting a habit
		return habitId;
	}
}
