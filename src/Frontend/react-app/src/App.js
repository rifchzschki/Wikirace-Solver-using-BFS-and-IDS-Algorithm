import { useEffect, useState } from "react";
// import { BrowserRouter as Router, Route, Switch } from "react-router-dom";

function App() {
  const [data, setData] = useState(null);

  useEffect(() => {
    fetch("http://localhost:8080/api/data")
      .then((response) => response.json())
      .then((data) => {
        // Update the state with the data received from the backend
        console.log(data);
        setData(data);
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  }, []);
  return <>Data: {JSON.stringify(data)}</>;
}

export default App;
