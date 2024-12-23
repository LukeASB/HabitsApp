const RegisterForm: React.FC = () => {
	return (
		<div className="register">
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
									placeholder="Enter your first name"
									required
								/>
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
									placeholder="Enter your last name"
									required
								/>
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
									placeholder="Enter your email"
									required
								/>
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
									placeholder="Enter your password"
									required
								/>
							</div>

							{/* Submit Button */}
							<button type="submit" className="btn btn-primary w-100">
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
