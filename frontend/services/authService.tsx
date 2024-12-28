import { mockLoggedInUser, mockRegisteredUser } from "../data/mock_users";
import { AuthModel } from "../model/authModel";
import ILoggedInUser from "../shared/interfaces/ILoggedInUser";
import ILoggedOutUser from "../shared/interfaces/ILoggedOutUser";
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
            if (!newAccessToken) throw new Error("No access token provided.");
            
            sessionStorage.setItem("access-token", newAccessToken);

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

        return { Success: false }; // on false, return an error modal
    }

    public static async logout(emailAddress: string): Promise<Partial<ILoggedOutUser>> {
        if (process.env.ENVIRONMENT === "DEV") {
            sessionStorage.removeItem("access-token");
            sessionStorage.removeItem("csrf-token");
            sessionStorage.removeItem("user-data");
            return { Success: true };
        }

        try {
            const response = await fetch(`/api/${process.env.API_URL}/logout`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ EmailAddress: emailAddress }),
            });

            if (!response.ok) throw new Error("Failed to logout.");

            const loggedOutUser: ILoggedInUser = await response.json();

            return loggedOutUser;
        } catch (ex) {
            console.log(ex);
        }

        return { Success: false };
    }

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
