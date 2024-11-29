import { ReactNode } from "react";
import Navbar from "./navbar";
import Footer from "./footer";

// import Sidebar from "./sidebar";

interface IMainLayout {
    children?: ReactNode
}

const MainLayout: React.FC<IMainLayout> = ( { children } ) => {
    return (
        <>
        <Navbar />
        <main>
            { children}
        </main>
        <Footer />
        </>
    );
};

export default MainLayout;