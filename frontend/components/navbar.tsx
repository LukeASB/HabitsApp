import { useState, useEffect } from "react";
import Link from "next/link";
import { AuthService } from "../services/authService";
import { useRouter } from "next/router";

const Navbar: React.FC = () => {
	const router = useRouter();
	const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);

	useEffect(() => (sessionStorage.getItem("access-token") ? setIsLoggedIn(true) : setIsLoggedIn(false)), []);

	return (
		<nav id="navbar" className="navbar navbar-expand-lg navbar-dark bg-dark">
			<div className="container-fluid">
				<Link className="navbar-brand" href="/">
					<strong>Habits Apps</strong>
				</Link>
				<button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
					<span className="navbar-toggler-icon"></span>
				</button>
				<div className="collapse navbar-collapse" id="navbarNav">
					<ul className="navbar-nav ms-auto">
						{!isLoggedIn && (
							<>
								<li className="nav-item">
									<Link className="nav-link" href="/register">
										Register
									</Link>
								</li>
								<li className="nav-item">
									<Link className="nav-link" href="/login">
										Login
									</Link>
								</li>
							</>
						)}
						{isLoggedIn && (
							<li className="nav-item">
								<span
									className="nav-link"
									style={{ cursor: "pointer" }}
									onClick={async () => {
										await AuthService.logout();
										router.push("/login");
										return setIsLoggedIn(false);
									}}
								>
									Logout
								</span>
							</li>
						)}
					</ul>
				</div>
			</div>
		</nav>
	);
};

export default Navbar;
