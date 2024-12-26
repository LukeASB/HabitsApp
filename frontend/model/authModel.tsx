import ILoginUser from "../shared/interfaces/ILoginUser";
import IRegisterUser from "../shared/interfaces/IRegisterUser";
import ILoginUserFormError from "../shared/interfaces/IRegisterUserFormError";
import IRegisterUserFormError from "../shared/interfaces/IRegisterUserFormError";
import { UsersValidation } from "../validation/usersValidation";
export class AuthModel {
    public static processRegisterUser = (registerUser: IRegisterUser): Partial<IRegisterUserFormError>[] => UsersValidation.validateRegisterUser(registerUser);
    public static processLoginUser = (loginUser: ILoginUser): Partial<ILoginUserFormError>[] => UsersValidation.validateLoginUser(loginUser);

	public static parseJWT(token: string = "") {
		try {
            if (!token) throw new Error("No token provided");
			const base64Url = token.split(".")[1];
			const base64 = base64Url.replace(/-/g, "+").replace(/_/g, "/");
			return JSON.parse(atob(base64));
		} catch (err) {
			console.error("Invalid JWT:", err);
		}
	}
}
