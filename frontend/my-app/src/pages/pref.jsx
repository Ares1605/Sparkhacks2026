// src/pages/NewPage.jsx

import '../App.css';
import Sidebar from "../components/SideBar";

import { useState } from "react";

export default function Pref() {

    const [preferences, setPreferences] = useState(""); 
    const [savedPreferences, setSavedPreferences] = useState("");

  return (
    <div className="layout">

      <div className="main">
        <div className="mainBox">

          <h1 className="pref-title">Preference Page</h1>

          <textarea 
            className="pref-textbox"
            placeholder="Tell us your preferences..."

            value={preferences}
            onChange={(e) => setPreferences(e.target.value)}
          />

            <button
            className="pref-update-button"
            onClick={() => setSavedPreferences(preferences)}
            >
            Update
            </button>

            <div className="pref-preview-box">
            {savedPreferences || "Your saved preferences will appear here..."}
            </div>


        </div>
      </div>
      <Sidebar/>
    </div>
  );
}


