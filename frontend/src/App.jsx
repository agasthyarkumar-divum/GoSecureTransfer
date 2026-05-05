import { useState, useEffect } from "react";

function App() {
  const [token, setToken] = useState(localStorage.getItem("token") || null);
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [files, setFiles] = useState([]); // ALWAYS array
  const [status, setStatus] = useState("");

  // 🔐 LOGIN
  const handleLogin = async () => {
    try {
      const res = await fetch("http://localhost:8080/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({ username, password })
      });

      const data = await res.json();

      if (data.token) {
        setToken(data.token);
        localStorage.setItem("token", data.token);
        setStatus("✅ Login successful");
      } else {
        setStatus("❌ Login failed");
      }
    } catch (err) {
      console.error(err);
      setStatus("❌ Login error");
    }
  };

  // 🚪 LOGOUT
  const handleLogout = () => {
    setToken(null);
    setFiles([]);
    localStorage.removeItem("token");
    setStatus("Logged out");
  };

  // 📂 FETCH FILES (SAFE)
  const fetchFiles = async (authToken = token) => {
    if (!authToken) return;

    try {
      const res = await fetch("http://localhost:8080/files", {
        headers: {
          Authorization: `Bearer ${authToken}`
        }
      });

      const data = await res.json();

      // ✅ ALWAYS ensure array
      if (Array.isArray(data)) {
        setFiles(data);
      } else {
        setFiles([]);
      }
    } catch (err) {
      console.error("Fetch error:", err);
      setFiles([]);
    }
  };

  // 🔄 AUTO FETCH AFTER LOGIN
  useEffect(() => {
    if (token) {
      fetchFiles(token);
    }
  }, [token]);

  // 📤 UPLOAD
  const handleUpload = async (e) => {
    const file = e.target.files[0];
    if (!file) return;

    const formData = new FormData();
    formData.append("file", file);

    try {
      await fetch("http://localhost:8080/upload", {
        method: "POST",
        headers: {
          Authorization: `Bearer ${token}`
        },
        body: formData
      });

      setStatus("✅ File uploaded");
      fetchFiles(); // refresh list
    } catch (err) {
      console.error(err);
      setStatus("❌ Upload failed");
    }
  };

  // 📥 DOWNLOAD
  const handleDownload = (filename) => {
    fetch(`http://localhost:8080/download?file=${filename}`, {
      headers: {
        Authorization: `Bearer ${token}`
      }
    })
      .then(res => res.blob())
      .then(blob => {
        const url = window.URL.createObjectURL(blob);
        const a = document.createElement("a");
        a.href = url;
        a.download = filename;
        a.click();
      })
      .catch(err => {
        console.error(err);
        setStatus("❌ Download failed");
      });
  };

  return (
    <div style={{ padding: "20px", fontFamily: "Arial" }}>
      <h1>🔐 Secure File Vault</h1>

      {/* 🔐 LOGIN */}
      {!token && (
        <div>
          <h2>Login</h2>
          <input
            placeholder="Username"
            onChange={(e) => setUsername(e.target.value)}
          />
          <input
            type="password"
            placeholder="Password"
            onChange={(e) => setPassword(e.target.value)}
          />
          <button onClick={handleLogin}>Login</button>
        </div>
      )}

      {/* 🚀 MAIN APP */}
      {token && (
        <div>
          <button onClick={handleLogout}>Logout</button>

          <h2>Upload File</h2>
          <input type="file" onChange={handleUpload} />

          <h2>Your Files</h2>

          {/* ✅ SAFE RENDER */}
          {!files || files.length === 0 ? (
            <p>📂 No files yet — upload something!</p>
          ) : (
            files.map((file, i) => (
              <div key={i}>
                {file}
                <button
                  onClick={() => handleDownload(file)}
                  style={{ marginLeft: "10px" }}
                >
                  Download
                </button>
              </div>
            ))
          )}
        </div>
      )}

      <p>{status}</p>
    </div>
  );
}

export default App;