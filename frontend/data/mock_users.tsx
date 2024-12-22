import IRegisterUser from "../shared/interfaces/IRegisterUser";

export const mockRegisterUsers: IRegisterUser[] = [{ EmailAddress: "johndoe1@example.com", Password: "secretPassword012!", FirstName: "John", LastName: "Doe"}];

export const createUser = (user: IRegisterUser) => {
    mockRegisterUsers.push(user);
}