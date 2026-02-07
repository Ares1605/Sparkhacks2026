import { Link } from "react-router-dom"
import '../App.css'
import logo from "frontend/my-app/src/assets/sparkhacks-logo.svg";

import { Link } from "react-router-dom"
import '../App.css'

export default function Sidebar() {
  return (
    <div className="sidebar">

      <img 
        src="https://www.sparkhacks.org/sparkhacks-logo.svg"
        alt="PreCog Logo"
        style={{ width: "80%", margin: "0 auto", display: "block" }}
      />

      <h2>PreCog</h2>

      <p>Navigation</p>

      <Link to="/" style={{ color: '#60a5fa', textDecoration: 'none' }}>Home</Link>
      <br />

      <Link to="/sync" style={{ color: '#60a5fa', textDecoration: 'none' }}>Sync</Link>
      <br />

      <Link to="/pref" style={{ color: '#60a5fa', textDecoration: 'none' }}>Pref</Link>

    </div>
  );
}
