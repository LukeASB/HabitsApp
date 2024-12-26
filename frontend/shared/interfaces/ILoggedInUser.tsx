export default interface ILoggedInUser {
    Success: boolean,
    User: {
        UserID: string;
        FirstName: string;
        LastName: string;
        EmailAddress: string;
        CreatedAt: string;
    },
    AccessToken: string;
    LoggedInAt: string;
}