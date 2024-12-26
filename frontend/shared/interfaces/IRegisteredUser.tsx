export default interface IRegisteredUser {
    Success: boolean,
    User: {
        UserID: string;
        FirstName: string;
        LastName: string;
        EmailAddress: string;
        CreatedAt: string;
    }
};