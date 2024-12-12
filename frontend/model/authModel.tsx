export class AuthModel {
	public static validateUser(habit: string): string[] {
		const errors: string[] = [];

		return errors;
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
