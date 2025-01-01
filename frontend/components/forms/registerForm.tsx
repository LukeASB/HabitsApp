import { useRouter } from "next/router";
import IRegisterUserFormData from "../../shared/interfaces/IRegisterUserFormData";
import { useState } from "react";
import { AuthModel } from "../../model/authModel";
import { AuthService } from "../../services/authService";
import ILoginUser from "../../shared/interfaces/ILoginUser";
import GenericErrorForm from "./genericErrorForm";
import { ModalTypeEnum } from "../../shared/enum/modalTypeEnum";
import Modal from "../modal/modal";
import IUserAuthModalTypes from "../../shared/interfaces/IUserAuthModalTypes";

const RegisterForm: React.FC = () => {
	const router = useRouter();
	const form: IRegisterUserFormData = { firstName: "", lastName: "", emailAddress: "", password: "" };
	const [formData, setFormData] = useState<IRegisterUserFormData>(form);
	const [errors, setErrors] = useState<IRegisterUserFormData>({ firstName: "", lastName: "", emailAddress: "", password: "" });
        const [showModal, setShowModal] = useState<IUserAuthModalTypes>({ RegisterErrorModal: false, LoginErrorModal: false });
    
        const handleOpenModal = (modalType: ModalTypeEnum) => {
            if (modalType === ModalTypeEnum.RegisterErrorModal) return setShowModal({ RegisterErrorModal: true, LoginErrorModal: false });
            if (modalType === ModalTypeEnum.LoginErrorModal) return setShowModal({ RegisterErrorModal: false, LoginErrorModal: true });
        };
    
        const handleCloseModal = (modalType: ModalTypeEnum) => {
            if (modalType === ModalTypeEnum.RegisterErrorModal) return setShowModal({ RegisterErrorModal: false, LoginErrorModal: false });
            if (modalType === ModalTypeEnum.LoginErrorModal) return setShowModal({ RegisterErrorModal: false, LoginErrorModal: false });
        };
    
        const onGenericErrorModalSubmit = () => {
            setShowModal({ RegisterErrorModal: false, LoginErrorModal: false });
        }

	const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		const { name, value } = e.target;
		setFormData((prevData) => ({
			...prevData,
			[name]: value,
		}));
	};

	const validateForm = () => {
		let isValid = true;
		let newErrors: IRegisterUserFormData = { firstName: "", lastName: "", emailAddress: "", password: "" };
		const userErrors = AuthModel.processRegisterUser({ ...formData });

		userErrors.forEach((error) => {
			if (error.firstName) {
				isValid = false;
				return (newErrors.firstName = error.firstName);
			}

			if (error.lastName) {
				isValid = false;
				return (newErrors.lastName = error.lastName);
			}

			if (error.emailAddress) {
				isValid = false;
				return (newErrors.emailAddress = error.emailAddress);
			}

			if (error.password) {
				isValid = false;
				return (newErrors.password = error.password);
			}
		});

		!isValid ? setErrors({ ...newErrors }) : setErrors({ firstName: "", lastName: "", emailAddress: "", password: "" });

		return isValid;
	};

	const onSubmit = async (registerUser: IRegisterUserFormData) => {
		const registeredUser = await AuthService.register(registerUser);
		if (!registeredUser.Success) {
            handleOpenModal(ModalTypeEnum.RegisterErrorModal);
            return;
        }

		const loginUser: ILoginUser = {
			emailAddress: registerUser.emailAddress,
			password: registerUser.password,
		};

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
		setFormData({ firstName: "", lastName: "", emailAddress: "", password: "" });
	};

	return (
		<div id="register" className="register">
            <Modal
                id="registerErrorModal"
                title="Register"
                body={<GenericErrorForm error="Failed To Register User" modalType={ModalTypeEnum.RegisterErrorModal} onSubmit={onGenericErrorModalSubmit} onModalClose={handleCloseModal}/>}
                showModal={showModal.RegisterErrorModal}
                modalType={ModalTypeEnum.RegisterErrorModal}
                onModalOpen={handleOpenModal}
                onModalClose={handleCloseModal}
                />
            <Modal
                id="loginErrorModal"
                title="Login"
                body={<GenericErrorForm error="Invalid Login Details" modalType={ModalTypeEnum.LoginErrorModal} onSubmit={onGenericErrorModalSubmit} onModalClose={handleCloseModal}/>}
                showModal={showModal.LoginErrorModal}
                modalType={ModalTypeEnum.LoginErrorModal}
                onModalOpen={handleOpenModal}
                onModalClose={handleCloseModal}
                />
			<div className="container mt-5">
				<div className="row justify-content-center">
					<div className="col-md-6">
						<h2 className="text-center mb-4">Register</h2>
						<form>
							{/* First Name */}
							<div className="mb-3">
								<label htmlFor="firstName" className="form-label">
									First Name
								</label>
								<input
									type="text"
									className="form-control"
									id="firstName"
									name="firstName"
									value={formData.firstName}
									onChange={handleChange}
									placeholder="Enter your first name"
									required
								/>
								<div className="error text-danger">{errors.firstName}</div>
							</div>
							{/* Last Name */}
							<div className="mb-3">
								<label htmlFor="lastName" className="form-label">
									Last Name
								</label>
								<input
									type="text"
									className="form-control"
									id="lastName"
									name="lastName"
									value={formData.lastName}
									onChange={handleChange}
									placeholder="Enter your last name"
									required
								/>
								<div className="error text-danger">{errors.lastName}</div>
							</div>

							{/* Email Address */}
							<div className="mb-3">
								<label htmlFor="email" className="form-label">
									Email Address
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
							{/* Password */}
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
								Register
							</button>
						</form>
					</div>
				</div>
			</div>
		</div>
	);
};

export default RegisterForm;
