import Navbar from "./navbar";
import Footer from "./footer";
import IMainLayout from "../shared/interfaces/IMainLayout";

const MainLayout: React.FC<IMainLayout> = ({ children }) => {
	return (
		<div className="mainLayout">
			<Navbar />
			<main>{children}</main>
			<Footer />
		</div>
	);
};

export default MainLayout;
