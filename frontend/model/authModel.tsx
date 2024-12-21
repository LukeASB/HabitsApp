import IRegisterUser from "../shared/interfaces/IRegisterUser";
export class AuthModel {
    private static readonly minPassword = 8;
    private static readonly maxPassword = 72;
    private static EmailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
    private static PasswordRegex = /^[a-zA-Z0-9!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]{8,20}$/;

    public static validateRegister(registerUser: IRegisterUser): string[] {
        if (AuthModel.EmailRegex.test(registerUser.EmailAddress) === false) return ["Invalid email address"];

        if (registerUser.Password.length < AuthModel.minPassword) return ["Password must be at least 8 characters"];
        if (registerUser.Password.length > AuthModel.maxPassword) return ["Password must be less than 72 characters"];
        if (AuthModel.PasswordRegex.test(registerUser.EmailAddress) === false) return ["Invalid password"];

        return [];
    }

	public static parseJWT(token: string) {
		try {
			const base64Url = token.split(".")[1];
			const base64 = base64Url.replace(/-/g, "+").replace(/_/g, "/");
			return JSON.parse(atob(base64));
		} catch (err) {
			console.error("Invalid JWT:", err);
			return null;
		}
	}
}
