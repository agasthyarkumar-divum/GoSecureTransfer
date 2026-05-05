import { useState } from "react";

function App() {
  const [status, setStatus] = useState("");

  const handleUpload = async (e) => {
    const file = e.target.files[0];
    if (!file) return;

    const formData = new FormData();
    formData.append("file", file);

    try {
      const res = await fetch("http://localhost:8080/upload", {
        method: "POST",
        body: formData,
      });

      const text = await res.text();
      setStatus(text);
    } catch (err) {
      setStatus("Upload failed");
    }
  };

  return (
    <div style={{ padding: "20px" }}>
      <h1>Secure File Vault</h1>

      <input type="file" onChange={handleUpload} />

      <p>{status}</p>
    </div>
  );
}

export default App;