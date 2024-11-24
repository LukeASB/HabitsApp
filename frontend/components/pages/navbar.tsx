import Link from 'next/link';

const Navbar: React.FC = () => {
    return (
        <nav>
            <ul>
                <li><Link href="/">Home</Link></li>
                <li><Link href="/login">Register</Link></li>
                <li><Link href="/login">Login</Link></li>
            </ul>
        </nav>
    );
};

export default Navbar;