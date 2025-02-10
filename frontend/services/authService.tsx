import { NextRouter } from "next/router";
import { mockLoggedInUser, mockRegisteredUser } from "../data/mock_users";
import { AuthModel } from "../model/authModel";
import IHabit from "../shared/interfaces/IHabit";
import ILoggedInUser from "../shared/interfaces/ILoggedInUser";
import ILoginUser from "../shared/interfaces/ILoginUser";
import IRegisteredUser from "../shared/interfaces/IRegisteredUser";
import IRegisterUser from "../shared/interfaces/IRegisterUser";

export class AuthService {
	public static async login(loginUser: ILoginUser): Promise<Partial<ILoggedInUser>> {
		if (process.env.ENVIRONMENT === "DEV") {
			sessionStorage.setItem("access-token", mockLoggedInUser.AccessToken);
			return mockLoggedInUser;
		}

		try {
			const response = await fetch(`/api/${process.env.API_URL}/login`, {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify(loginUser),
			});

			if (!response.ok) throw new Error("Failed to login.");

			const newAccessToken = response.headers.get("Authorization");
			const newCSRFToken = response.headers.get("X-Csrf-Token");
			if (!newAccessToken || !newCSRFToken) throw new Error("No access token provided.");

			sessionStorage.setItem("access-token", newAccessToken);
			sessionStorage.setItem("csrf-token", newCSRFToken);

			const loggedInUser: ILoggedInUser = await response.json();
			sessionStorage.setItem("user-data", JSON.stringify({ emailAddress: loggedInUser?.User?.EmailAddress, firstName: loggedInUser?.User?.FirstName, lastName: loggedInUser?.User?.LastName }));

			return loggedInUser;
		} catch (ex) {
			console.log(ex);
		}

		return { Success: false };
	}

	public static async register(registerUser: IRegisterUser): Promise<Partial<IRegisteredUser>> {
		if (process.env.ENVIRONMENT === "DEV") return mockRegisteredUser;

		try {
			const response = await fetch(`/api/${process.env.API_URL}/register`, {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify(registerUser),
			});

			const registeredUser: IRegisteredUser = await response.json();

			return registeredUser;
		} catch (ex) {
			console.log(ex);
		}

		return { Success: false };
	}

	public static async logout(router: NextRouter) {
		if (process.env.ENVIRONMENT === "DEV") {
			sessionStorage.removeItem("access-token");
			sessionStorage.removeItem("user-data");
		}

		try {
			const shortlivedJWTAccessToken = sessionStorage.getItem("access-token");

			const response = await fetch(`/api/${process.env.API_URL}/logout`, {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
					Authorization: shortlivedJWTAccessToken || "",
				},
			});

			if (response.status === 401) await AuthService.refresh(AuthService.logout, router);

			if (!response.ok) throw new Error("Failed to logout.");
		} catch (ex) {
			console.log(ex);
		} finally {
			sessionStorage.removeItem("access-token");
			sessionStorage.removeItem("csrf-token");
			sessionStorage.removeItem("user-data");
		}
	}

	public static async refresh(callback: Function, router: NextRouter, callbackBody?: string | IHabit | IHabit[]) {
		try {
			const shortlivedJWTAccessToken = sessionStorage.getItem("access-token");
			const userAccessToken = shortlivedJWTAccessToken ? AuthModel.parseJWT(shortlivedJWTAccessToken) : null;
			const userData = !userAccessToken ? sessionStorage.getItem("user-data") : null;

			if (!userAccessToken) {
				console.log("User Doesn't Have An Access Token. Attempt to get Email Address from 'user-data' Session Storage");
				if (!userData) {
					router.push("/login");
					console.log("User Doesn't Have An Email Address. Redirecting to Login Page");
					return;
				}
				return;
			}

			const response = await fetch(`/api/${process.env.API_URL}/refresh`, {
				method: "POST",
				body: JSON.stringify({ EmailAddress: userAccessToken.username }),
			});

			if (!response.ok) {
				router.push("/login");
				return;
			}
			const newAccessToken = response.headers.get("Authorization");
			const newCSRFToken = response.headers.get("X-Csrf-Token");
			if (!newAccessToken || !newCSRFToken) throw new Error("No access token provided.");

			sessionStorage.setItem("access-token", newAccessToken);
			sessionStorage.setItem("csrf-token", newCSRFToken);

			if (callbackBody && typeof callbackBody === "string") return callback(callbackBody);
			if (callbackBody && typeof callbackBody === "object") return callback(callbackBody);
			if (callbackBody && Array.isArray(callbackBody)) return callback(callbackBody);
			callback();
		} catch (err) {
			console.log(err);
		}

		router.push("/login");
	}
}
