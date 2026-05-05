import Upload from "./Upload";
import FileList from "./FileList";
import { LockIcon, UploadIcon, FolderIcon } from "./Icons";

function Dashboard({ token, setToken }) {
  const logout = () => {
    localStorage.removeItem("token");
    setToken(null);
  };

  return (
    <div className="dashboard-container">
      <div className="dashboard-header">
        <h1 style={{ display: "flex", alignItems: "center", gap: "8px" }}>
          <LockIcon size={28} color="#0066cc" />
          SecureVault
        </h1>
        <button className="logout-btn" onClick={logout}>
          Sign Out
        </button>
      </div>

      <div className="dashboard-content">
        <div className="dashboard-section">
          <h2 style={{ display: "flex", alignItems: "center", gap: "8px" }}>
            <UploadIcon size={20} color="#0066cc" />
            Upload Files
          </h2>
          <Upload token={token} />
        </div>

        <div className="dashboard-section">
          <h2 style={{ display: "flex", alignItems: "center", gap: "8px" }}>
            <FolderIcon size={20} color="#0066cc" />
            Your Files
          </h2>
          <FileList token={token} />
        </div>
      </div>
    </div>
  );
}

export default Dashboard;