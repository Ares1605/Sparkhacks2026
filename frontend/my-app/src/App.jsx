import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import Sync from "./pages/sync";
import Home from "./pages/home"
import './App.css'

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/sync" element={<Sync />} />
      </Routes>
    </Router>
  );
}

export default App
