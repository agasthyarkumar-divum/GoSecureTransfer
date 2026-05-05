function Upload({ token }) {
  const upload = async (e) => {
    const file = e.target.files[0];

    const formData = new FormData();
    formData.append("file", file);

    await fetch("http://localhost:8080/upload", {
      method: "POST",
      headers: {
        Authorization: `Bearer ${token}`
      },
      body: formData
    });

    alert("Uploaded");
  };

  return <input type="file" onChange={upload} />;
}

export default Upload;