import { Link } from "react-router-dom"
import '../App.css'

export default function Sidebar() {
    return (
      <div className="sidebar">
        <h2 style={{ textAlign: "center", width: "100%"}}>PreCog</h2>
        <p>Navigation</p>
        <Link to="/" style={{ color: '#60a5fa', textDecoration: 'none' }}> Home </Link>
        <br/>
        <Link to="/sync" style={{ color: '#60a5fa', textDecoration: 'none' }}> Sync </Link>
        <br/>
        <Link to="/pref" style={{ color: '#60a5fa', textDecoration: 'none' }}> Pref </Link>
      </div>
    );
}