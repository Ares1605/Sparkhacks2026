import LinkPill from './LinkPill';
import '../App.css'

export default function Sidebar() {
  return (
    <div className="sidebar">

      <h2 style = {{ fontWeight: "800", fontFamily: 'Brush Script MT', fontSize: "3.5rem", marginBottom: 0}}>PreCog</h2>

      <p>Navigation:</p>

      <LinkPill to="/">Home</LinkPill>
      <br />
      <LinkPill to="/sync">Sync</LinkPill>
      <br />
      <LinkPill to="/pref">Preference</LinkPill>

      <img 
        src="https://www.sparkhacks.org/sparkhacks-logo.svg"
        alt="PreCog Logo"
        style={{ width: "80%", margin: "0 auto", display: "block" }}
      />      
      <p style={{ marginTop: "0px", fontSize: "0.85rem", opacity: 0.8 }}>
        SparkHacks <br /> 2026 <br /> @UIC
      </p>

    </div>
  );
}
