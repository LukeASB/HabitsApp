export default interface IRegisterUserFormData {
    firstName: string;
    lastName: string;
    emailAddress: string;
    password: string;
    [key: string]: string | number;
};