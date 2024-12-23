export default interface ILoginUser {
    emailAddress: string; /*`json:"EmailAddress"`*/
	password:     string; /*`json:"Password"`*/
    [key: string]: string;
}