import { useEffect, useState } from "react";
import { apiCall } from "../utils/api";
import { RefreshIcon, FileIcon, DownloadIcon } from "./Icons";

function FileList({ token }) {
  const [files, setFiles] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  const loadFiles = async () => {
    setLoading(true);
    setError("");
    try {
      const res = await apiCall("/files", {
        headers: {
          Authorization: `Bearer ${token}`
        }
      });
      const data = await res.json();
      setFiles(Array.isArray(data) ? data : []);
    } catch (err) {
      setError("Failed to load files");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadFiles();
  }, [token]);

  const download = async (file) => {
    try {
      const res = await apiCall(`/download?file=${file}`, {
        headers: {
          Authorization: `Bearer ${token}`
        }
      });
      
      if (!res.ok) {
        alert("Download failed");
        return;
      }

      const blob = await res.blob();
      const url = window.URL.createObjectURL(blob);
      const a = document.createElement("a");
      a.href = url;
      a.download = file;
      a.click();
      window.URL.revokeObjectURL(url);
    } catch (err) {
      alert("Download error");
    }
  };

  if (loading) {
    return (
      <div className="file-list-container">
        <div className="loading">
          <div className="spinner"></div>
        </div>
      </div>
    );
  }

  return (
    <div className="file-list-container">
      {error && <div className="error-message">{error}</div>}

      <div className="files-header">
        <h3>Your Secure Files</h3>
        <button className="refresh-btn" onClick={loadFiles} style={{ display: "flex", alignItems: "center", gap: "4px" }}>
          <RefreshIcon size={16} color="#0066cc" />
          Refresh
        </button>
      </div>

      {files.length === 0 ? (
        <div className="empty-state">
          <p>No files yet. Upload one to get started!</p>
        </div>
      ) : (
        <div className="files-list">
          {files.map((file, i) => (
            <div key={i} className="file-item">
              <div className="file-name" style={{ display: "flex", alignItems: "center", gap: "8px" }}>
                <FileIcon size={20} color="#0066cc" />
                <div className="file-text">
                  <p>{file}</p>
                  <p className="date">{new Date().toLocaleDateString()}</p>
                </div>
              </div>
              <div className="file-actions">
                <button
                  className="download-btn"
                  onClick={() => download(file)}
                  title="Download file"
                  style={{ display: "flex", alignItems: "center", gap: "4px" }}
                >
                  <DownloadIcon size={16} color="#0066cc" />
                  Download
                </button>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

export default FileList;