import IRegisterUser from "../shared/interfaces/IRegisterUser";

export const mockRegisterUsers: IRegisterUser[] = [{ emailAddress: "johndoe1@example.com", password: "secretPassword012!", firstName: "John", lastName: "Doe"}];

export const createUser = (user: IRegisterUser) => {
    mockRegisterUsers.push(user);
}