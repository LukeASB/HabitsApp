export default interface IRegisteredUser {
    Success: boolean,
    User: {
        FirstName: string;
        LastName: string;
        EmailAddress: string;
        CreatedAt: string;
    }
};