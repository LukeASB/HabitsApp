import { ReactNode } from "react";
import Navbar from "./navbar";
import Footer from "./footer";

interface IMainLayout {
    children?: ReactNode
}

const MainLayout: React.FC<IMainLayout> = ( { children } ) => {
    return (
        <div className="mainLayout">
            <Navbar />
            <main>
                { children}
            </main>
            <Footer />
        </div>
    );
};

export default MainLayout;