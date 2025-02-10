import React from "react";
import Home from "./pages/home";
import Login from "./pages/login";
import Register from "./pages/register";
import MainLayout from "./mainLayout";

interface IApp {
	page?: string;
}

const App: React.FC<IApp> = ({ page = "" }) => {
	const renderPage = () => {
		if (page === "home") return <Home />;
		if (page === "register") return <Register />;
		if (page === "login") return <Login />;

		return <h3>No component for navigation value. {page} not found</h3>;
	};

	return <MainLayout>{renderPage()}</MainLayout>;
};

export default App;
