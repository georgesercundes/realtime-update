import React from "react";

const socket = new WebSocket("ws://127.0.0.1:8080/websocket");

function Websocket() {
  const [message, setMessage] = React.useState("");
  const [inputValue, setInputValue] = React.useState("");

  React.useEffect(() => {
    socket.onmessage = (e) => {
      setMessage("Message from the server: " + JSON.parse(e.data).message);
    };
  }, [message]);

  const handleClick = () => {
    socket.send(JSON.stringify({ message: inputValue }));

    return () => {
      socket.close();
    };
  };

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(e.target.value);
  };

  return (
    <div className="Websocket">
      <h1 style={{ color: "#DDE31C" }}>Websocket</h1>
      <input type="text" value={inputValue} onChange={handleInputChange} />
      <button
        onClick={handleClick}
        style={{
          background: "#DDE31C",
          margin: "8px",
          padding: "4px",
          border: 0,
          borderRadius: "4px",
        }}
      >
        Send
      </button>
      <h3 style={{ color: "#7ECBE3" }}>{message}</h3>
    </div>
  );
}

export default Websocket;
