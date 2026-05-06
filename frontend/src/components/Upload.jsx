import { useRef, useState } from "react";
import { apiCall, getApiUrl } from "../utils/api";
import { FolderIcon } from "./Icons";

function Upload({ token }) {
  const fileInputRef = useRef(null);
  const [uploading, setUploading] = useState(false);
  const [progress, setProgress] = useState(0);
  const [error, setError] = useState("");
  const [dragOver, setDragOver] = useState(false);

  const handleUpload = async (file) => {
    if (!file) return;

    setUploading(true);
    setProgress(0);
    setError("");

    try {
      const formData = new FormData();
      formData.append("file", file);

      console.log("📤 Starting upload for file:", file.name);
      console.log("🔗 Uploading to:", getApiUrl ? getApiUrl() + "/upload" : "/upload");

      const res = await apiCall("/upload", {
        method: "POST",
        headers: {
          Authorization: `Bearer ${token}`
        },
        body: formData
      });

      console.log("📊 Upload response status:", res.status);
      console.log("✅ Upload response:", res);

      if (res.ok) {
        setProgress(100);
        setTimeout(() => {
          setProgress(0);
          if (fileInputRef.current) fileInputRef.current.value = "";
        }, 1000);
        window.location.reload();
      } else {
        setError("Upload failed. Please try again.");
      }
    } catch (err) {
      console.error("❌ Upload error:", err);
      setError("Connection error. Please try again.");
    } finally {
      setUploading(false);
    }
  };

  const handleFileChange = (e) => {
    const file = e.target.files?.[0];
    if (file) handleUpload(file);
  };

  const handleDrop = (e) => {
    e.preventDefault();
    setDragOver(false);
    const file = e.dataTransfer.files?.[0];
    if (file) handleUpload(file);
  };

  const handleDragOver = (e) => {
    e.preventDefault();
    setDragOver(true);
  };

  const handleDragLeave = () => {
    setDragOver(false);
  };

  return (
    <div className="upload-section">
      {error && <div className="error-message">{error}</div>}

      <div className="file-input-wrapper">
        <label
          className={`upload-label ${dragOver ? "drag-over" : ""}`}
          onDrop={handleDrop}
          onDragOver={handleDragOver}
          onDragLeave={handleDragLeave}
        >
          <div className="upload-label-content">
            <div style={{ display: "flex", justifyContent: "center", marginBottom: "12px" }}>
              <FolderIcon size={40} color="#0066cc" />
            </div>
            <p>
              {uploading ? "Uploading..." : "Click or drag files here"}
            </p>
            <p className="subtitle">
              {uploading ? "Please wait..." : "Max 100MB per file"}
            </p>
          </div>
          <input
            ref={fileInputRef}
            type="file"
            onChange={handleFileChange}
            disabled={uploading}
          />
        </label>
      </div>

      {uploading && progress > 0 && (
        <div className="upload-progress">
          <div className="progress-bar">
            <div className="progress-fill" style={{ width: `${progress}%` }} />
          </div>
          <div className="upload-status">{progress}% uploaded</div>
        </div>
      )}
    </div>
  );
}

export default Upload;