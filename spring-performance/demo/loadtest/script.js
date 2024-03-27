import http from "k6/http";
import { check } from "k6";
import { randomString } from "https://jslib.k6.io/k6-utils/1.2.0/index.js";
// import { htmlReport } from "https://raw.githubusercontent.com/benc-uk/k6-reporter/main/dist/bundle.js";

// export function handleSummary(data) {
//   return {
//     "summary.html": htmlReport(data),
//   };
// }

export const options = {
  scenarios: {
    contacts: {
      executor: "constant-arrival-rate",
      duration: "2m",
      rate: 2000,
      timeUnit: "1s",
      preAllocatedVUs: 3000,
    },
  },
};
// export const options = {
//   scenarios: {
//     contacts: {
//       executor: "ramping-arrival-rate",
//       timeUnit: "1s",
//       preAllocatedVUs: 100,
//       maxVUs: 10000,
//       stages: [
//         { target: 10000, duration: "10s" },
//         { target: 15000, duration: "10s" },
//         { target: 20000, duration: "40s" },
//         { target: 0, duration: "10s" },
//       ],
//     },
//   },
// thresholds: {
//   http_req_duration: ["p(95)<60000"], //units in miliseconds 60000ms = 1m
//   http_req_failed: ["rate<0.01"], // http errors should be less than 1%
//   checks: ["rate>0.99"],
// },
// };
export default function () {
  const chars = `abcdefghijklxyzvbnm`;
  const k = randomString(3, chars);
  const v = randomString(3, chars);
  let res = http.post(`http://localhost:9191/cache?key=${k}&value=${v}`, {
    // headers: { "Content-Type": "application/json" },
  });
  check(res, {
    "is status 200": (r) => r.status === 200,
  });
  // console.log(res.json());
}
