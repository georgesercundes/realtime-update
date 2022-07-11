import http from "k6/http";

import { check } from "k6";
import { randomString } from "https://jslib.k6.io/k6-utils/1.2.0/index.js";

export const options = {
  insecureSkipTLSVerify: true,
  noConnectionReuse: true,
  iterations: 1000,
  vus: 1
};

export default function () {
  const res = http.post(
    "https://127.0.0.1:5000/log",
    JSON.stringify({ message: randomString(10) })
  );
  check(res, { "status was 200": (r) => r.status == 200 });
}
