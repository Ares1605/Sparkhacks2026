import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import Sync from "./pages/sync";
import Home from "./pages/home"
import Pref from "./pages/pref"
import API from "./api/api"

import './App.css'
import { useMemo } from "react";


function App() {
  const api = useMemo(() => new API("http://localhost:8080"), [])

  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home api={api} />} />
        <Route path="/sync" element={<Sync api={api} />} />
        <Route path="/pref" element={<Pref api={api} />} />
      </Routes>
    </Router>
  );
}

export default App
