import Websocket from "./components/Websocket";
import Sse from "./components/ServerSentEvents";
import './index.css'

function App() {
  return (
    <div
      className="App"
      style={{
        display: "flex",  
        justifyContent: "space-around",
        backgroundColor: "#0A0806",
        height: "100vh",
      }}
    >
      <Websocket />
      <Sse />
    </div>
  );
}

export default App;
