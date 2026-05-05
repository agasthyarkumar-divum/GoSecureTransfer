import Upload from "./Upload";
import FileList from "./FileList";

function Dashboard({ token, setToken }) {
  const logout = () => {
    localStorage.removeItem("token");
    setToken(null);
  };

  return (
    <div>
      <h1>Secure File Vault</h1>
      <button onClick={logout}>Logout</button>

      <Upload token={token} />
      <FileList token={token} />
    </div>
  );
}

export default Dashboard;