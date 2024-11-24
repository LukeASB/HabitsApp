import { ReactNode } from "react";
import Banner from "./banner";
import Navbar from "./pages/navbar";
import Link from "next/link";

interface IMainLayout {
    children?: ReactNode
}

const MainLayout: React.FC<IMainLayout> = ( { children } ) => {
    return (
        <>
        <header>
            <Banner><h1><Link href="/">Do Habits App</Link></h1></Banner>
        </header>
        <Navbar />
        <main>
            { children }
        </main>
        </>
    );
};

export default MainLayout;