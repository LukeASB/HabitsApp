import { mockhabits, createMockHabit, retrieveMockHabit, updateMockHabit, deleteMockHabit } from "../data/mock_habits";
import { HabitsModel } from "../model/habitsModel";
import IHabit from "../shared/interfaces/IHabit";
import { HabitsValidation } from "../validation/habitsValidation";
import { AuthService } from "./authService";

export class HabitsService {
	public static async createHabit(habit: IHabit) {
		if (process.env.ENVIRONMENT === "DEV") return createMockHabit(habit);

		try {
			const csrfToken = sessionStorage.getItem("csrf-token");
			const shortlivedJWTAccessToken = sessionStorage.getItem("access-token");

			const response = await fetch(`/api/${process.env.API_URL}/createhabit`, {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
					"X-CSRF-Token": csrfToken || "",
					Authorization: shortlivedJWTAccessToken || "",
				},
				body: JSON.stringify(habit),
			});

			if (response.status === 403) await AuthService.refresh(HabitsService.createHabit);
			if (!response.ok) throw new Error("Failed to fetch habits.");

		} catch (ex) {
			console.log(ex);
		}
	}

	public static async retrieveHabits(): Promise<IHabit[]> {
		if (process.env.ENVIRONMENT === "DEV") return mockhabits;

		try {
			const csrfToken = sessionStorage.getItem("csrf-token");
			const shortlivedJWTAccessToken = sessionStorage.getItem("access-token");

			const response = await fetch(`/api/${process.env.API_URL}/retrievehabits`, {
				method: "GET",
				headers: {
					"Content-Type": "application/json",
					"X-CSRF-Token": csrfToken || "",
					Authorization: shortlivedJWTAccessToken || "",
				},
			});

			if (response.status === 403) await AuthService.refresh(HabitsService.retrieveHabits);
			if (!response.ok) throw new Error("Failed to fetch habits.");

			const data: IHabit[] = await response.json();

			return data;
		} catch (ex) {
			console.log(ex);
			return [];
		}
	}

	public static async retrieveHabit(habitId: string): Promise<IHabit[]> {
		if (process.env.ENVIRONMENT === "DEV") return retrieveMockHabit(habitId);

		try {
			const csrfToken = sessionStorage.getItem("csrf-token");
			const shortlivedJWTAccessToken = sessionStorage.getItem("access-token");

			const response = await fetch(`/api/${process.env.API_URL}/retrievehabit?habitId=${habitId}`, {
				method: "GET",
				headers: {
					"Content-Type": "application/json",
					"X-CSRF-Token": csrfToken || "",
					Authorization: shortlivedJWTAccessToken || "",
				},
			});

			if (response.status === 403) await AuthService.refresh(HabitsService.retrieveHabit);
			if (!response.ok) throw new Error("Failed to fetch habit.");

			const data: IHabit[] = await response.json();

			return data;
		} catch (ex) {
			console.log(ex);
			return [];
		}
	}

	public static async updateHabit(habit: IHabit): Promise<IHabit | null> {
		if (process.env.ENVIRONMENT === "DEV") return updateMockHabit(habit);

		try {
			const csrfToken = sessionStorage.getItem("csrf-token");
			const shortlivedJWTAccessToken = sessionStorage.getItem("access-token");

			const response = await fetch(`/api/${process.env.API_URL}/updatehabit?habitId=${habit.habitId}`, {
				method: "PUT",
				headers: {
					"Content-Type": "application/json",
					"X-CSRF-Token": csrfToken || "",
					Authorization: shortlivedJWTAccessToken || "",
				},
                body: JSON.stringify(habit),
			});

			if (response.status === 403) await AuthService.refresh(HabitsService.updateHabit);
			if (!response.ok) throw new Error("Failed to fetch habit.");

			const data: IHabit = await response.json();

			return data;
		} catch (ex) {
			console.log(ex);
			return null;
		}
	}

	public static async deleteHabit(habitId: string): Promise<boolean> {
		if (process.env.ENVIRONMENT === "DEV") return deleteMockHabit(habitId);

		try {
			const csrfToken = sessionStorage.getItem("csrf-token");
			const shortlivedJWTAccessToken = sessionStorage.getItem("access-token");

			const response = await fetch(`/api/${process.env.API_URL}/deletehabit?habitId=${habitId}`, {
				method: "DELETE",
				headers: {
					"Content-Type": "application/json",
					"X-CSRF-Token": csrfToken || "",
					Authorization: shortlivedJWTAccessToken || "",
				},
			});

			if (response.status === 403) await AuthService.refresh(HabitsService.deleteHabit);
			if (!response.ok) throw new Error("Failed to delete habit.");

			const data = await response.json();

			return data;
		} catch (ex) {
			console.log(ex);
			return false;
		}
	}
}
