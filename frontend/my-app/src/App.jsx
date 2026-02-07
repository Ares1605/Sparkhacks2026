import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import Sync from "./pages/sync";
import Home from "./pages/home"
import Pref from "./pages/pref"
import './App.css'


function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/sync" element={<Sync />} />
        <Route path="/pref" element={<Pref />} />
      </Routes>
    </Router>
  );
}

export default App
