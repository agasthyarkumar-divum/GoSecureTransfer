import { useState } from "react";
import Login from "./components/Login";
import Register from "./components/Register";
import Dashboard from "./components/Dashboard";

function App() {
  const [token, setToken] = useState(localStorage.getItem("token"));
  const [mode, setMode] = useState("login");

  if (!token) {
    return (
      <div>
        {mode === "login" ? (
          <>
            <Login setToken={setToken} />
            <p onClick={() => setMode("register")}>New user? Register</p>
          </>
        ) : (
          <>
            <Register />
            <p onClick={() => setMode("login")}>Already have account?</p>
          </>
        )}
      </div>
    );
  }

  return <Dashboard token={token} setToken={setToken} />;
}

export default App;