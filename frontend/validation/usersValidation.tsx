
import ILoginUser from "../shared/interfaces/ILoginUser";
import IRegisterUser from "../shared/interfaces/IRegisterUser";
import ILoginUserFormError from "../shared/interfaces/IRegisterUserFormError";
import IRegisterUserFormError from "../shared/interfaces/IRegisterUserFormError";

export class UsersValidation {
    private static readonly maxName = 50;
    private static readonly minPassword = 8;
    private static readonly maxPassword = 72;
    private static readonly matchUpperLowerCaseLettersOnly = /^[a-zA-Z ]*$/;
    private static readonly EmailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
    private static readonly PasswordRegex = /^[a-zA-Z0-9!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]{8,20}$/;

    public static validateRegisterUser(registerUser: Partial<IRegisterUser>): Partial<IRegisterUserFormError>[] {
        const errors: Partial<IRegisterUserFormError>[] = [];

        const firstName = UsersValidation.validateName(registerUser.firstName);
        firstName && errors.push({ firstName: firstName });

        const lastName = UsersValidation.validateName(registerUser.lastName);
        lastName && errors.push({ lastName: lastName });

        const email = UsersValidation.validateUserEmail(registerUser.emailAddress);
        email && errors.push({ emailAddress: email });

        const password = UsersValidation.validateUserPassword(registerUser.password);
        password && errors.push({ password: password });
        
        return errors;
    }

    public static validateLoginUser(loginUser: Partial<ILoginUser>): Partial<ILoginUserFormError>[] {
        const errors: Partial<ILoginUserFormError>[] = [];

        const email = UsersValidation.validateUserEmail(loginUser.emailAddress);
        email && errors.push({ emailAddress: email });

        const password = UsersValidation.validateUserPassword(loginUser.password);
        password && errors.push({ password: password });

        return errors;
    }

    private static validateName(name?: string): string {
        if (name?.trim().length === 0) return "Required.";
        if (name && name.trim().length >= UsersValidation.maxName) return `Name exceeds max character length of ${UsersValidation.maxName}`;
        if (name && !UsersValidation.matchUpperLowerCaseLettersOnly.test(name)) return "Invalid.";
        return "";
    }

    private static validateUserEmail(email?: string): string {
        if (email?.trim().length === 0) return "Email address is required";
        if (email && UsersValidation.EmailRegex.test(email) === false) return "Invalid email address";
        return "";
    }

    private static validateUserPassword(password?: string): string {
        if (password?.trim().length === 0) return "Required.";
        if (password && password.length < UsersValidation.minPassword) return "Password must be at least 8 characters";
        if (password && password.length > UsersValidation.maxPassword) return "Password must be less than 72 characters";
        if (password && UsersValidation.PasswordRegex.test(password) === false) return "Invalid.";
        return "";
    }

}