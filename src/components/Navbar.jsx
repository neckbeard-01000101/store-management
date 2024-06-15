import { Link, useResolvedPath, useLocation } from "react-router-dom";

export default function Navbar() {
    return (
        <nav className="navbar">
            <ul className="nav-links">
                <CustomLink to="/form">Form</CustomLink>
                <CustomLink to="/data">Data</CustomLink>
            </ul>
        </nav>
    );
}
export function CustomLink({ to, children, ...props }) {
    const resolvedPath = useResolvedPath(to);
    const currentPath = useLocation().pathname;
    const isActive = currentPath.startsWith(resolvedPath.pathname);

    return (
        <li className={`nav-link ${isActive ? "active" : ""}`}>
            <Link to={to} {...props}>
                {children}
            </Link>
        </li>
    );
}
