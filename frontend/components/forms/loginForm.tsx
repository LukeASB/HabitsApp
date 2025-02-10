import { useState } from "react";
import ILoginUserFormData from "../../shared/interfaces/ILoginUserFormData";
import ILoginUserFormError from "../../shared/interfaces/ILoginUserFormError";
import { AuthModel } from "../../model/authModel";
import ILoginUser from "../../shared/interfaces/ILoginUser";
import { AuthService } from "../../services/authService";
import { useRouter } from "next/router";
import { ModalTypeEnum } from "../../shared/enum/modalTypeEnum";
import Modal from "../modal/modal";
import GenericErrorForm from "./genericErrorForm";
import IUserAuthModalTypes from "../../shared/interfaces/IUserAuthModalTypes";

const LoginForm: React.FC = () => {
	const router = useRouter();
	const form: ILoginUserFormData = { emailAddress: "", password: "" };
	const [formData, setFormData] = useState<ILoginUserFormData>(form);
	const [errors, setErrors] = useState<ILoginUserFormError>({ emailAddress: "", password: "" });
	const [showModal, setShowModal] = useState<IUserAuthModalTypes>({ RegisterErrorModal: false, LoginErrorModal: false });

	const handleOpenModal = (modalType: ModalTypeEnum) => {
		if (modalType === ModalTypeEnum.LoginErrorModal) return setShowModal({ RegisterErrorModal: false, LoginErrorModal: true });
	};

	const handleCloseModal = (modalType: ModalTypeEnum) => {
		if (modalType === ModalTypeEnum.LoginErrorModal) return setShowModal({ RegisterErrorModal: false, LoginErrorModal: false });
	};

	const onGenericErrorModalSubmit = () => {
		setShowModal({ RegisterErrorModal: false, LoginErrorModal: false });
	};

	const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		const { name, value } = e.target;
		setFormData((prevData) => ({
			...prevData,
			[name]: value,
		}));
	};

	const validateForm = () => {
		let isValid = true;
		const newErrors: ILoginUserFormError = { emailAddress: "", password: "" };
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

		if (!isValid) {
            setErrors({ ...newErrors });
        } else {
            setErrors({ emailAddress: "", password: "" });
        }

		return isValid;
	};

	const onSubmit = async (loginUser: ILoginUser) => {
		const loggedInUser = await AuthService.login(loginUser);
		if (!loggedInUser.Success) {
			handleOpenModal(ModalTypeEnum.LoginErrorModal);
			return;
		}
		router.push("/");
	};

	const handleSubmit = (e: React.FormEvent) => {
		e.preventDefault();
		if (!validateForm()) return;

		onSubmit({ ...formData });
		setFormData({ emailAddress: "", password: "" });
	};

	return (
		<div id="login" className="login">
			<Modal
				id="loginErrorModal"
				title="Login"
				body={<GenericErrorForm error="Invalid Login Details" modalType={ModalTypeEnum.LoginErrorModal} onSubmit={onGenericErrorModalSubmit} onModalClose={handleCloseModal} />}
				showModal={showModal.LoginErrorModal}
				modalType={ModalTypeEnum.LoginErrorModal}
				onModalOpen={handleOpenModal}
				onModalClose={handleCloseModal}
			/>
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
									name="emailAddress"
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
									name="password"
									value={formData.password}
									onChange={handleChange}
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
