import { Link } from "react-router-dom"
import '../app.css'

export default function Sidebar() {
    return (
      <div className="sidebar">
        <h2>Sidebar</h2>
        <p>Navigation</p>
        <Link to="/" style={{ color: '#60a5fa', textDecoration: 'none' }}> Home </Link>
        <br/>
        <Link to="/sync" style={{ color: '#60a5fa', textDecoration: 'none' }}> Sync </Link>
      </div>
    );
}