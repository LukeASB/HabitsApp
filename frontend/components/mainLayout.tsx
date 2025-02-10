import Navbar from "./navbar";
import Footer from "./footer";
import IMainLayout from "../shared/interfaces/IMainLayout";
import { useEffect, useState } from "react";
import UnsupportedScreenMessage from "./unsupportedScreenMessage";

const MainLayout: React.FC<IMainLayout> = ({ children }) => {
	const [isUnsupported, setIsUnsupported] = useState<boolean>(false);

	useEffect(() => {
		const checkScreenSize = () => {
			setIsUnsupported(window.innerWidth < 300);
		};

		checkScreenSize();

		window.addEventListener("resize", checkScreenSize);
	}, []);

	if (isUnsupported) return <UnsupportedScreenMessage isUnsupported={isUnsupported} />;

	return (
		<div id="mainLayout" className="mainLayout">
			<Navbar />
			<main>{children}</main>
			<Footer />
		</div>
	);
};

export default MainLayout;
