import { HabitsModel } from "../model/habitsModel";
import IHabit from "../shared/interfaces/IHabit";

export class HabitsService {
    public static async createHabit(habit: IHabit) {
    const validationErrors = HabitsModel.validateHabit(habit);
    if (validationErrors.length > 0) {
      throw new Error(validationErrors.join(', '));
    }

    // Call the backend API to persist the habit
    const response = await fetch('/api/habits/createHabit', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(habit),
    });

    if (!response.ok) {
      throw new Error('Failed to create habit.');
    }

    return await response.json();
  }

  public static async retrieveHabits() {
    const response = await fetch('/api/habits/retrieveHabits');
    if (!response.ok) {
      throw new Error('Failed to fetch habits.');
    }
    return await response.json();
  }

  public static async retrieveHabit(habitId: string) {
    const response = await fetch(`/api/habits/retrieveHabit/${habitId}`);
    if (!response.ok) {
      throw new Error('Failed to fetch habit.');
    }
    return await response.json();
  }

  public static async updateHabit(habitId: string, updates: Partial<IHabit>) {
    const response = await fetch(`/api/habits/updateHabit/${habitId}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(updates),
    });

    if (!response.ok) {
      throw new Error('Failed to update habit.');
    }

    return await response.json();
  }

  public static async deleteHabit(habitId: string) {
    const response = await fetch(`/api/habits/deleteHabit/${habitId}`, {
      method: 'DELETE',
    });

    if (!response.ok) {
      throw new Error('Failed to delete habit.');
    }

    return await response.json();
  }
}
