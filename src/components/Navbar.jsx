import { Link, useResolvedPath, useLocation } from "react-router-dom";

export default function Navbar() {
    return (
        <nav className="navbar">
            <ul className="nav-links">
                <CustomLink to="/add-order">Add new orders</CustomLink>
                <CustomLink to="/orders">Orders</CustomLink>
                <CustomLink to="/profits">Monthly profits</CustomLink>
                <CustomLink to="/storage">Storage</CustomLink>
                <CustomLink to="/add-to-storage">Add to storage</CustomLink>
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
