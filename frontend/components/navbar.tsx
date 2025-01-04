import { useState, useEffect } from "react";
import Link from "next/link";
import { HabitsService } from "../services/habitsService";
import { AuthService } from "../services/authService";

const Navbar: React.FC = () => {
	const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);

	useEffect(() => (sessionStorage.getItem("access-token") ? setIsLoggedIn(true) : setIsLoggedIn(false)), []);

	return (
		<nav id="navbar" className="navbar navbar-expand-lg navbar-light bg-dark">
			<div className="container-fluid">
				<strong>
					<span className="navbar-brand text-light">
						<Link className="navbar-brand text-light" href="/">
							Habits Apps
						</Link>
					</span>
				</strong>
			</div>
			<div className="navbar" id="navbarNav">
				<ul className="navbar-nav">
					<li className="nav-item">
						{!isLoggedIn && (
							<>
								<Link className="navbar-brand text-light" href="/register">
									Register
								</Link>
								<Link className="navbar-brand text-light" href="/login">
									Login
								</Link>
							</>
						)}
                        <span className="navbar-brand text-light" style={{ cursor: 'pointer' }} onClick={async () => {
                            const logout = await AuthService.logout();
                            if (logout?.Success) return setIsLoggedIn(false);
                            
                            alert("Failed to logout.");
                            
                        }}>Logout</span>
					</li>
				</ul>
			</div>
		</nav>
	);
};

export default Navbar;
