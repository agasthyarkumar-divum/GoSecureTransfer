import { useEffect, useState } from "react";
import { apiCall } from "../utils/api";
import { RefreshIcon, FileIcon, DownloadIcon } from "./Icons";

function FileList({ token }) {
  const [files, setFiles] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [deleting, setDeleting] = useState(null);

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

  const deleteFile = async (file) => {
    if (!window.confirm(`Are you sure you want to delete "${file}"?`)) {
      return;
    }

    setDeleting(file);
    try {
      console.log("🗑️ Deleting file:", file);
      const res = await apiCall(`/delete?file=${file}`, {
        method: "DELETE",
        headers: {
          Authorization: `Bearer ${token}`
        }
      });

      console.log("Delete response:", res.status);

      if (res.ok) {
        console.log("✅ File deleted successfully");
        setFiles(files.filter(f => f !== file));
      } else {
        const errorData = await res.json().catch(() => ({}));
        setError("Failed to delete file: " + (errorData.message || "Unknown error"));
      }
    } catch (err) {
      console.error("❌ Delete error:", err);
      setError("Delete error: " + err.message);
    } finally {
      setDeleting(null);
    }
  };

  const deleteAllFiles = async () => {
    if (!window.confirm("Are you sure you want to delete ALL files? This action cannot be undone.")) {
      return;
    }

    setDeleting("all");
    try {
      console.log("🗑️ Deleting all files");
      const res = await apiCall("/delete-all", {
        method: "DELETE",
        headers: {
          Authorization: `Bearer ${token}`
        }
      });

      console.log("Delete all response:", res.status);

      if (res.ok) {
        const data = await res.json();
        console.log("✅ All files deleted. Deleted count:", data.deleted);
        setFiles([]);
      } else {
        const errorData = await res.json().catch(() => ({}));
        setError("Failed to delete all files: " + (errorData.message || "Unknown error"));
      }
    } catch (err) {
      console.error("❌ Delete all error:", err);
      setError("Delete error: " + err.message);
    } finally {
      setDeleting(null);
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
        <div className="header-buttons" style={{ display: "flex", gap: "8px" }}>
          <button className="refresh-btn" onClick={loadFiles} style={{ display: "flex", alignItems: "center", gap: "4px" }}>
            <RefreshIcon size={16} color="#0066cc" />
            Refresh
          </button>
          {files.length > 0 && (
            <button 
              className="delete-all-btn" 
              onClick={deleteAllFiles}
              disabled={deleting === "all"}
              style={{ 
                display: "flex", 
                alignItems: "center", 
                gap: "4px",
                backgroundColor: "#dc3545",
                color: "white",
                border: "none",
                padding: "8px 12px",
                borderRadius: "4px",
                cursor: deleting === "all" ? "not-allowed" : "pointer",
                opacity: deleting === "all" ? 0.6 : 1
              }}
            >
              🗑️ Delete All
            </button>
          )}
        </div>
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
                  disabled={deleting !== null}
                  style={{ display: "flex", alignItems: "center", gap: "4px" }}
                >
                  <DownloadIcon size={16} color="#0066cc" />
                  Download
                </button>
                <button
                  className="delete-btn"
                  onClick={() => deleteFile(file)}
                  disabled={deleting !== null}
                  title="Delete file"
                  style={{
                    display: "flex",
                    alignItems: "center",
                    gap: "4px",
                    backgroundColor: "#dc3545",
                    color: "white",
                    border: "none",
                    padding: "6px 10px",
                    borderRadius: "4px",
                    cursor: deleting !== null ? "not-allowed" : "pointer",
                    opacity: deleting !== null ? 0.6 : 1
                  }}
                >
                  {deleting === file ? "🗑️ Deleting..." : "🗑️ Delete"}
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