import { useState, useEffect } from "react";

function App() {
  const [token, setToken] = useState("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [files, setFiles] = useState([]);
  const [status, setStatus] = useState("");

  // 🔐 Login
  const handleLogin = async () => {
    try {
      const res = await fetch("http://localhost:8080/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ username, password }),
      });

      const data = await res.json();

      if (data.token) {
        setToken(data.token);
        setStatus("✅ Login successful");
        fetchFiles(data.token);
      } else {
        setStatus("❌ Login failed");
      }
    } catch (err) {
      setStatus("❌ Error logging in");
    }
  };

  // 📂 Fetch files (authorized)
  const fetchFiles = async (authToken = token) => {
    if (!authToken) return;

    try {
      const res = await fetch("http://localhost:8080/files", {
        headers: {
          Authorization: `Bearer ${authToken}`,
        },
      });

      const data = await res.json();
      setFiles(data);
    } catch (err) {
      console.error("Error fetching files:", err);
    }
  };

  // 📤 Upload file
  const handleUpload = async (e) => {
    const file = e.target.files[0];
    if (!file) return;

    const formData = new FormData();
    formData.append("file", file);

    try {
      const res = await fetch("http://localhost:8080/upload", {
        method: "POST",
        headers: {
          Authorization: `Bearer ${token}`,
        },
        body: formData,
      });

      const text = await res.text();
      setStatus(text);

      fetchFiles(); // refresh list
    } catch (err) {
      setStatus("❌ Upload failed");
    }
  };

  // 📥 Download file
  const handleDownload = (filename) => {
    fetch(`http://localhost:8080/download?file=${filename}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
      .then((res) => res.blob())
      .then((blob) => {
        const url = window.URL.createObjectURL(blob);
        const a = document.createElement("a");
        a.href = url;
        a.download = filename;
        a.click();
      })
      .catch(() => setStatus("❌ Download failed"));
  };

  // 🔄 Auto-refresh file list after login
  useEffect(() => {
    if (token) {
      fetchFiles();
    }
  }, [token]);

  return (
    <div style={{ padding: "20px", fontFamily: "Arial" }}>
      <h1>🔐 Secure File Vault</h1>

      {/* 🔐 Login Section */}
      {!token && (
        <div style={{ marginBottom: "20px" }}>
          <h2>Login</h2>
          <input
            placeholder="Username"
            onChange={(e) => setUsername(e.target.value)}
            style={{ marginRight: "10px" }}
          />
          <input
            type="password"
            placeholder="Password"
            onChange={(e) => setPassword(e.target.value)}
            style={{ marginRight: "10px" }}
          />
          <button onClick={handleLogin}>Login</button>
        </div>
      )}

      {/* 📤 Upload Section */}
      {token && (
        <>
          <div style={{ marginBottom: "20px" }}>
            <h2>Upload File</h2>
            <input type="file" onChange={handleUpload} />
          </div>

          {/* 📂 File List */}
          <div>
            <h2>Your Files</h2>
            {files.length === 0 ? (
              <p>No files uploaded</p>
            ) : (
              <ul>
                {files.map((file, index) => (
                  <li key={index}>
                    {file}
                    <button
                      onClick={() => handleDownload(file)}
                      style={{ marginLeft: "10px" }}
                    >
                      Download
                    </button>
                  </li>
                ))}
              </ul>
            )}
          </div>
        </>
      )}

      {/* 📢 Status */}
      <p style={{ marginTop: "20px" }}>{status}</p>
    </div>
  );
}

export default App;