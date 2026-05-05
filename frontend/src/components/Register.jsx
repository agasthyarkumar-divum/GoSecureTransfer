import { useState } from "react";

function Register() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [confirm, setConfirm] = useState("");

  const register = async () => {
    if (password !== confirm) {
      alert("Passwords do not match");
      return;
    }

    const res = await fetch("http://localhost:8080/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({ username, password })
    });

    const text = await res.text();
    alert(text);
  };

  return (
    <div>
      <h2>Register</h2>
      <input placeholder="username" onChange={e => setUsername(e.target.value)} />
      <input type="password" placeholder="password" onChange={e => setPassword(e.target.value)} />
      <input type="password" placeholder="confirm password" onChange={e => setConfirm(e.target.value)} />
      <button onClick={register}>Register</button>
    </div>
  );
}

export default Register;