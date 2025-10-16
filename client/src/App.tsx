// import { react } from 'react'
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Dashboard from './pages/Dashboard';
import Home from './pages/Home';
import Login from "./pages/Login";

import ProtectedRoute from "./components/protected_routes";

function App() {

  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<Login />} />

          {/* --------------------------secured routes------------------------------- */}

          {/* <Route path="/dashboard" element={<Dashboard />} /> */}

          <Route
            path="/dashboard"
            element={
              <ProtectedRoute>
                <Dashboard />
              </ProtectedRoute>
            }/>

            {/* ----------------------------------------------------------------------------------- */}
      </Routes>
      </BrowserRouter>
    </>
  )
}

export default App
