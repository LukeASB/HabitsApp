import { NextRouter } from "next/router";
import { mockhabits, createMockHabit, retrieveMockHabit, updateMockHabit, deleteMockHabit, updateAllMockHabits } from "../data/mock_habits";
import IHabit from "../shared/interfaces/IHabit";
import { AuthService } from "./authService";

export class HabitsService {
	public static async createHabit(habit: IHabit, router: NextRouter): Promise<IHabit | null> {
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

			if (response.status === 401) await AuthService.refresh(HabitsService.createHabit, router, habit);
			if (response.status === 403) await AuthService.refresh(HabitsService.createHabit, router, habit);

			if (!response.ok) throw new Error("Failed to fetch habits.");

            const newAccessToken = response.headers.get("Authorization");
			const newCSRFToken = response.headers.get("X-Csrf-Token");
			if (!newAccessToken || !newCSRFToken) throw new Error("No access token provided.");

			sessionStorage.setItem("access-token", newAccessToken);
			sessionStorage.setItem("csrf-token", newCSRFToken);

			const data: IHabit = await response.json();

			return data;
		} catch (ex) {
			console.log(ex);
			return null;
		}
	}

	public static async retrieveHabits(router: NextRouter): Promise<IHabit[]> {
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

			if (response.status === 401) await AuthService.refresh(HabitsService.retrieveHabits, router);
			if (!response.ok) throw new Error("Failed to fetch habits.");

            const newAccessToken = response.headers.get("Authorization");
			if (!newAccessToken) throw new Error("No access token provided.");

			sessionStorage.setItem("access-token", newAccessToken);

			const data: IHabit[] = await response.json();

			return data;
		} catch (ex) {
			console.log(ex);
			return [];
		}
	}

	public static async retrieveHabit(habitId: string, router: NextRouter): Promise<IHabit[]> {
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

			if (response.status === 401) await AuthService.refresh(HabitsService.retrieveHabit, router, habitId);
			if (!response.ok) throw new Error("Failed to fetch habit.");

            const newAccessToken = response.headers.get("Authorization");
			if (!newAccessToken) throw new Error("No access token provided.");

			sessionStorage.setItem("access-token", newAccessToken);

			const data: IHabit[] = await response.json();

			return data;
		} catch (ex) {
			console.log(ex);
			return [];
		}
	}

	public static async updateHabit(habit: IHabit, router: NextRouter): Promise<IHabit | null> {
		if (process.env.ENVIRONMENT === "DEV") return updateMockHabit(habit);

		try {
			const csrfToken = sessionStorage.getItem("csrf-token");
			const shortlivedJWTAccessToken = sessionStorage.getItem("access-token");

			const response = await fetch(`/api/${process.env.API_URL}/updatehabit`, {
				method: "PUT",
				headers: {
					"Content-Type": "application/json",
					"X-CSRF-Token": csrfToken || "",
					Authorization: shortlivedJWTAccessToken || "",
				},
				body: JSON.stringify(habit),
			});

			if (response.status === 401) await AuthService.refresh(HabitsService.updateHabit, router, habit);
			if (response.status === 403) await AuthService.refresh(HabitsService.updateHabit, router, habit);
			if (!response.ok) throw new Error("Failed to fetch habit.");

            const newAccessToken = response.headers.get("Authorization");
			const newCSRFToken = response.headers.get("X-Csrf-Token");
			if (!newAccessToken || !newCSRFToken) throw new Error("No access token provided.");

			sessionStorage.setItem("access-token", newAccessToken);
			sessionStorage.setItem("csrf-token", newCSRFToken);

			const data: IHabit = await response.json();

			return data;
		} catch (ex) {
			console.log(ex);
			return null;
		}
	}

	public static async updateAllHabits(habit: IHabit[], router: NextRouter): Promise<IHabit[] | null> {
		if (process.env.ENVIRONMENT === "DEV") return updateAllMockHabits(habit);

		try {
			const csrfToken = sessionStorage.getItem("csrf-token");
			const shortlivedJWTAccessToken = sessionStorage.getItem("access-token");

			const response = await fetch(`/api/${process.env.API_URL}/updatehabits`, {
				method: "PUT",
				headers: {
					"Content-Type": "application/json",
					"X-CSRF-Token": csrfToken || "",
					Authorization: shortlivedJWTAccessToken || "",
				},
				body: JSON.stringify(habit),
			});

			if (response.status === 401) await AuthService.refresh(HabitsService.updateAllHabits, router, habit);
			if (response.status === 403) await AuthService.refresh(HabitsService.updateAllHabits, router, habit);
			if (!response.ok) throw new Error("Failed to fetch habit.");

            const newAccessToken = response.headers.get("Authorization");
			const newCSRFToken = response.headers.get("X-Csrf-Token");
			if (!newAccessToken || !newCSRFToken) throw new Error("No access token provided.");

			sessionStorage.setItem("access-token", newAccessToken);
			sessionStorage.setItem("csrf-token", newCSRFToken);

			const data: IHabit[] = await response.json();

			return data;
		} catch (ex) {
			console.log(ex);
			return null;
		}
	}

	public static async deleteHabit(habitId: string, router: NextRouter): Promise<boolean> {
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

			if (response.status === 401) await AuthService.refresh(HabitsService.deleteHabit, router, habitId);
			if (response.status === 403) await AuthService.refresh(HabitsService.deleteHabit, router, habitId);
			if (!response.ok) throw new Error("Failed to delete habit.");

            const newAccessToken = response.headers.get("Authorization");
			const newCSRFToken = response.headers.get("X-Csrf-Token");
			if (!newAccessToken || !newCSRFToken) throw new Error("No access token provided.");

			sessionStorage.setItem("access-token", newAccessToken);
			sessionStorage.setItem("csrf-token", newCSRFToken);

			const data = await response.json();

			return data;
		} catch (ex) {
			console.log(ex);
			return false;
		}
	}
}
