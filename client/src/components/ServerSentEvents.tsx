import axios from "axios";
import React from "react";

function ServerSentEvents() {
  const [message, setMessage] = React.useState("");
  const [inputValue, setInputValue] = React.useState("");

  React.useEffect(() => {
    const eventSource = new EventSource("https://127.0.0.1:5000/sse");
    eventSource.onmessage = (e) => {
      setMessage("Message from the server: " + e.data);
    };

    return () => {
      eventSource.close();
    };
  }, []);

  const handleClick = () => {
    axios.post(
      "https://127.0.0.1:5000/log",
      JSON.stringify({ message: inputValue })
    );
  };

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(e.target.value);
  };

  return (
    <div className="ServerSentEvents">
      <h1 style={{ color: "#DDE31C" }}>Server Sent Events</h1>
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

export default ServerSentEvents;
