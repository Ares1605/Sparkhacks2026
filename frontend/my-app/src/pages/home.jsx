import '../App.css'
import { Link } from "react-router-dom";

function Home() {
  return (
    <div className="layout">
      <div className="main">
        <div className="mainBox">
          <h1>Main Placeholder</h1>
          <p>yap yap yap yap yap</p>
        </div>

        <div className="subBox">
          <h2>Bottom Box</h2>
          <p>Additional content</p>
        </div>
      </div>

      <div className="sidebar">
        <h2>Sidebar</h2>
        <p>Navigation</p>
        <Link to="/sync" style={{ color: '#60a5fa', textDecoration: 'none' }}> Sync </Link>
      </div>
    </div>
  );
}

export default Home;