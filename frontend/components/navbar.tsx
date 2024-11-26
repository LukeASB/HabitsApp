import Link from 'next/link';

const Navbar: React.FC = () => {
    return (
        <nav className="navbar navbar-expand-lg navbar-light bg-primary">
            <div className="container content">
            <strong>
                <Link className="navbar-brand text-light" href="/">
                    Habits Apps
                </Link>
            </strong>
            </div>
        <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
            <span className="navbar-toggler-icon"></span>
          </button>
        <div className="collapse navbar-collapse" id="navbarNav">
            <ul className="navbar-nav">
              <li className="nav-item active">
                <Link className="navbar-brand text-light" href="/">Home</Link>
                <Link className="navbar-brand text-light" href="/register">Register</Link>
                <Link className="navbar-brand text-light" href="/login">Login</Link>
              </li>
            </ul>
          </div>
            {/* <ul>
                <li><Link href="/">Home</Link></li>
                <li></li>
                <li></li>
            </ul>*/}
        </nav>
    );
};

export default Navbar;