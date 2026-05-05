import { useState } from "react";
import Login from "./components/Login";
import Register from "./components/Register";
import Dashboard from "./components/Dashboard";
import "./index.css";

function App() {
  const [token, setToken] = useState(localStorage.getItem("token"));
  const [mode, setMode] = useState("login");

  if (!token) {
    return (
      <>
        {mode === "login" ? (
          <Login setToken={setToken} setMode={setMode} />
        ) : (
          <Register setMode={setMode} />
        )}
      </>
    );
  }

  return <Dashboard token={token} setToken={setToken} />;
}

export default App;