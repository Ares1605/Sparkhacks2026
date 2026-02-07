import { Link } from "react-router-dom";
import './LinkPill.css';

export default function LinkPill({to, children}) {
    return (
        <Link to={to} className="link-pill">
            {children}
        </Link>
    );
}