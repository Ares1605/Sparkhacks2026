import { Link } from "react-router-dom"
import '../App.css'


export default function Sidebar() {
  return (
    <div className="sidebar">

      <h2 style = {{ fontWeight: "800", fontSize: "1.4rem"}}>PreCog</h2>

      <p>Navigation</p>

      <Link to="/" style={{ color: '#60a5fa', textDecoration: 'none' }}>Home</Link>
      <br />

      <Link to="/sync" style={{ color: '#60a5fa', textDecoration: 'none' }}>Sync</Link>
      <br />

      <Link to="/pref" style={{ color: '#60a5fa', textDecoration: 'none' }}>Pref</Link>

      <img 
        src="https://www.sparkhacks.org/sparkhacks-logo.svg"
        alt="PreCog Logo"
        style={{ width: "80%", margin: "0 auto", display: "block" }}
      />      
      <p style={{ marginTop: "12px", fontSize: "0.85rem", opacity: 0.8 }}>
        SparkHacks <br /> 2026 <br /> @UIC
      </p>

    </div>
  );
}
