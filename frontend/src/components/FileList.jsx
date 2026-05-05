import { useEffect, useState } from "react";

function FileList({ token }) {
  const [files, setFiles] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8080/files", {
      headers: {
        Authorization: `Bearer ${token}`
      }
    })
      .then(res => res.json())
      .then(data => setFiles(Array.isArray(data) ? data : []));
  }, [token]);

  const download = (file) => {
    fetch(`http://localhost:8080/download?file=${file}`, {
      headers: {
        Authorization: `Bearer ${token}`
      }
    })
      .then(res => res.blob())
      .then(blob => {
        const url = window.URL.createObjectURL(blob);
        const a = document.createElement("a");
        a.href = url;
        a.download = file;
        a.click();
      });
  };

  return (
    <div>
      <h2>Your Files</h2>
      {files.length === 0 ? (
        <p>No files</p>
      ) : (
        files.map((f, i) => (
          <div key={i}>
            {f}
            <button onClick={() => download(f)}>Download</button>
          </div>
        ))
      )}
    </div>
  );
}

export default FileList;