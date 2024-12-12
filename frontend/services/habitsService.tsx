import { mockhabits } from "../data/mock_habits";
import { HabitsModel } from "../model/habitsModel";
import IHabit from "../shared/interfaces/IHabit";
import { AuthService } from "./authService";

export class HabitsService {
	public static async createHabit(habit: IHabit) {
		const validationErrors = HabitsModel.validateHabit(habit);
		if (validationErrors.length > 0) {
			throw new Error(validationErrors.join(", "));
		}

		// Call the backend API to persist the habit
		const response = await fetch("/api/habits/createHabit", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify(habit),
		});

		if (!response.ok) {
			throw new Error("Failed to create habit.");
		}

		return await response.json();
	}

	public static async retrieveHabits(): Promise<IHabit[]> {
		console.log(process.env.ENVIRONMENT);
		if (process.env.ENVIRONMENT === "DEV") {
			return mockhabits;
		}

		try {
			const csrfToken = sessionStorage.getItem("csrf-token");
			const shortlivedJWTAccessToken =
				sessionStorage.getItem("access-token");

			const response = await fetch(
				`/api/${process.env.API_URL}/retrievehabits`,
				{
					method: "GET",
					headers: {
						"Content-Type": "application/json",
						"X-CSRF-Token": csrfToken || "",
						Authorization: `Bearer ${shortlivedJWTAccessToken || ""}`,
					},
				},
			);

			if (response.status === 403) {
				await AuthService.refresh(HabitsService.retrieveHabits);
			}

			if (!response.ok) {
				throw new Error("Failed to fetch habits.");
			}

			const data: IHabit[] = await response.json();

			return data;
		} catch (ex) {
			return [];
		}
	}

	public static async retrieveHabit(habitId: string) {
		const response = await fetch(`/api/habits/retrieveHabit/${habitId}`);
		if (!response.ok) {
			throw new Error("Failed to fetch habit.");
		}
		return await response.json();
	}

	public static async updateHabit(habitId: string, updates: Partial<IHabit>) {
		const response = await fetch(`/api/habits/updateHabit/${habitId}`, {
			method: "PUT",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify(updates),
		});

		if (!response.ok) {
			throw new Error("Failed to update habit.");
		}

		return await response.json();
	}

	public static async deleteHabit(habitId: string) {
		const response = await fetch(`/api/habits/deleteHabit/${habitId}`, {
			method: "DELETE",
		});

		if (!response.ok) {
			throw new Error("Failed to delete habit.");
		}

		return await response.json();
	}
}
