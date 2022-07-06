import ws from "k6/ws";

import { check } from "k6";
import { randomString } from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';

export default function () {
  const url = "ws://127.0.0.1:8080/websocket";
  const response = ws.connect(url, null, function (socket) {
    socket.send(JSON.stringify({ message: randomString(10) }));
    socket.close();
  });

  check(response, { "status is 101": (r) => r && r.status === 101 });
}
