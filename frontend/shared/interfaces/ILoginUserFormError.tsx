export default interface IRegisterUserFormError {
    firstName: string;
    lastName: string;
    emailAddress: string;
    password: string;
    [key: string]: string
}