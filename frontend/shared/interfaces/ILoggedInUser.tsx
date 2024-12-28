export default interface ILoggedInUser {
    Success: boolean,
    User: {
        FirstName: string;
        LastName: string;
        EmailAddress: string;
        CreatedAt: string;
    },
    AccessToken: string;
    LoggedInAt: string;
}