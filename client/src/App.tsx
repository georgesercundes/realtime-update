import React from "react";

const socket = new WebSocket("ws://127.0.0.1:8080/websocket");

function App() {
  const [message, setMessage] = React.useState("");
  const [inputValue, setInputValue] = React.useState("");

  React.useEffect(() => {
    socket.onopen = () => {
      setMessage("Connected");
    };

    socket.onmessage = (e) => {
      setMessage("Message from the server: " + JSON.parse(e.data).message);
    };

  }, []);
  
  const handleClick = () => {
    socket.send(JSON.stringify({ message: inputValue }));

      return () => {
        socket.close();
      };
  }

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(e.target.value)
  }

  return (
    <div className="App">
      <input type="text" value={inputValue} onChange={handleInputChange} />
      <button onClick={handleClick}>Send</button>
      <div>{message}</div>
    </div>
  );
}

export default App;
