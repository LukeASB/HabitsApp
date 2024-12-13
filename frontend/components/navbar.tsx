import { useState, useEffect } from "react";
import Link from "next/link";
import IHabit from "../shared/interfaces/IHabit";

/**
 * @returns
 */
const Navbar: React.FC = () => {
	const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);

	useEffect(() => (sessionStorage.getItem("access-token") ? setIsLoggedIn(true) : setIsLoggedIn(false)), []);

	return (
		<nav className="navbar navbar-expand-lg navbar-light bg-dark">
			<div className="container content">
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
					</li>
				</ul>
			</div>
		</nav>
	);
};

export default Navbar;
