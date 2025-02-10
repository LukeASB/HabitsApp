import App from "../components/app";
import { useEffect, useState } from "react";
import { useRouter } from "next/router";

const Index = () => {
	const router = useRouter();
	const [loading, setLoading] = useState(true); // Loading state while checking token

	useEffect(() => {
        if (!sessionStorage.getItem("access-token")) {
            router.push("/login");
            return;
        }

		setLoading(false);
	}, [router]);

	if (loading) {
		return (
			<div className="d-flex justify-content-center align-items-center vh-100">
				<div className="spinner-border" role="status">
					<span className="visually-hidden">Loading...</span>
				</div>
			</div>
		);
	}

	return <App page="home" />;
};

export default Index;
