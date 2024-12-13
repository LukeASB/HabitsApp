import { AuthModel } from "../model/authModel";

export class AuthService {
	public static async login() {}

	public static async register() {}

	public static async refresh(callback: Function) {
		try {
			const shortlivedJWTAccessToken = sessionStorage.getItem("access-token");
			const userData = shortlivedJWTAccessToken ? AuthModel.parseJWT(shortlivedJWTAccessToken) : null;
			const response = await fetch(`/api/${process.env.API_URL}/refresh`, {
				method: "POST",
				body: JSON.stringify({ EmailAddress: userData.Email }),
			});

			if (!response.ok) window.location.href = "/login";
			const data = await response.json();
			sessionStorage.setItem("access-token", data.Token);
			callback();
		} catch (err) {
			window.location.href = "/login";
		}
	}
}
