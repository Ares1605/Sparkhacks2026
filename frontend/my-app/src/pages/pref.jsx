// src/pages/NewPage.jsx

import Sidebar from "../components/SideBar";

export default function NewPage() {
  return (
    <div className="layout">
        <div className="main">
            <div className="mainBox">
                <h1>New Page</h1>
                <p>This is a basic new page screen.</p>
            </div>
            <Sidebar/>
        </div>
    </div>
  );
}
