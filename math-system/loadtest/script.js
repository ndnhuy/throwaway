import http from "k6/http";
import { check } from "k6";
// import { htmlReport } from "https://raw.githubusercontent.com/benc-uk/k6-reporter/main/dist/bundle.js";

// export function handleSummary(data) {
//   return {
//     "summary.html": htmlReport(data),
//   };
// }

// export const options = {
//   scenarios: {
//     contacts: {
//       executor: "constant-arrival-rate",
//       duration: "1m",
//       rate: 200,
//       timeUnit: "1s",
//       preAllocatedVUs: 100,
//     },
//   },
// };
export const options = {
  scenarios: {
    contacts: {
      executor: "ramping-arrival-rate",
      timeUnit: "1s",
      preAllocatedVUs: 100,
      maxVUs: 100,
      stages: [
        { target: 50, duration: "10s" },
        { target: 80, duration: "10s" },
        { target: 100, duration: "40s" },
        // { target: 0, duration: "10s" },
      ],
    },
  },
  // thresholds: {
  //   http_req_duration: ["p(95)<60000"], //units in miliseconds 60000ms = 1m
  //   http_req_failed: ["rate<0.01"], // http errors should be less than 1%
  //   checks: ["rate>0.99"],
  // },
};
export default function () {
  let res = http.post(
    "http://localhost:8989/add",
    JSON.stringify({ a: 1, b: 2 }),
    {
      headers: { "Content-Type": "application/json" },
    }
  );
  check(res, {
    "is status 200": (r) => r.status === 200,
  });
  // console.log(res.json());
}
