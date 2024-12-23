import { useState } from "react";
import ILoginUserFormData from "../../shared/interfaces/ILoginUserFormData";
import ILoginUserFormError from "../../shared/interfaces/IRegisterUserFormError";
import { AuthModel } from "../../model/authModel";

const LoginForm: React.FC = () => {
    const form: ILoginUserFormData = { emailAddress: "", password: "" };
    const [formData, setFormData] = useState<ILoginUserFormData>(form);
    const [errors, setErrors] = useState<ILoginUserFormError>({ emailAddress: "", password: "" });

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setFormData((prevData) => ({
            ...prevData,
            [name]: value,
        }));
    };

    const validateForm = () => {
        let isValid = true;
        let newErrors: ILoginUserFormError = { emailAddress: "", password: "" };
        const userErrors = AuthModel.processLoginUser({ ...formData });

        userErrors.forEach((error) => {
            if (error.emailAddress) {
                isValid = false;
                return (newErrors.emailAddress = error.emailAddress);
            }

            if (error.password) {
                isValid = false;
                return (newErrors.password = error.password);
            }
        });

        !isValid ? setErrors({ ...newErrors }) : setErrors({ emailAddress: "", password: "" });

        return isValid;
    };

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        if (!validateForm()) return;

        // onsubmit(); // call endpoint. If successful, store the access-token in session. Redirect User to Habits Page.
        setFormData({ emailAddress: "", password: "" });
    };


	return (
		<div className="login">
			<div className="container mt-5">
				<div className="row justify-content-center">
					<div className="col-md-6">
						<h2 className="text-center mb-4">Login</h2>
						<form>
							{/* Email Field */}
							<div className="mb-3">
								<label htmlFor="email" className="form-label">
									Email address
								</label>
								<input
									type="email"
									className="form-control"
									id="email"
                                    name="days"
                                    value={formData.emailAddress}
                                    onChange={handleChange}
									placeholder="Enter your email"
									required
								/>
                                <div className="error text-danger">{errors.emailAddress}</div>
							</div>

							{/* Password Field */}
							<div className="mb-3">
								<label htmlFor="password" className="form-label">
									Password
								</label>
								<input
									type="password"
									className="form-control"
									id="password"
									placeholder="Enter your password"
									required
								/>
                                <div className="error text-danger">{errors.password}</div>
							</div>

							{/* Submit Button */}
							<button type="submit" className="btn btn-primary w-100" onClick={handleSubmit}>
								Login
							</button>
						</form>
					</div>
				</div>
			</div>
		</div>
	);
};

export default LoginForm;
