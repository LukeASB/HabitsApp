import ILoggedInUser from "../shared/interfaces/ILoggedInUser";
import IRegisteredUser from "../shared/interfaces/IRegisteredUser";
import IRegisterUser from "../shared/interfaces/IRegisterUser";

export const mockRegisterUsers: IRegisterUser[] = [{ emailAddress: "johndoe1@example.com", password: "1secret?Password", firstName: "John", lastName: "Doe"}];

export const createUser = (user: IRegisterUser) => {
    mockRegisterUsers.push(user);
};

export const mockRegisteredUser: IRegisteredUser = {
    Success: true,
    User: {
        FirstName: "TestUser123",
        LastName: "TestUser123",
        EmailAddress: "test222@example.com",
        CreatedAt: "2024-12-25T07:24:29.9453151Z"
    }
};

export const mockLoggedInUser: ILoggedInUser = {
    Success: true,
    User: {
        FirstName: "John",
        LastName: "Doe",
        EmailAddress: "johndoe1@example.com",
        CreatedAt: "2024-10-10T09:00:00Z"
    },
    AccessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG5kb2UxQGV4YW1wbGUuY29tIiwiZXhwIjoxNzM1MTEwOTc5fQ.SBGPK0ogWTe_T6HrVJhhK33qOjJCd6H43Q0GJWcLmfU",
    LoggedInAt: "2024-12-25T07:11:19.9480848Z"
};