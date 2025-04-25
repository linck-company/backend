// import http from "k6/http";
// import { check } from "k6";

// export const options = {
// 	vus: 1,
// 	duration: "10s",
// };

// export default () => {
// 	const url = "http://139.59.108.119:8081/srm/linckapiv1/auth/login"; // Replace with your actual URL
// 	const payload = JSON.stringify({
// 		username: "karthik",
// 		password: "cartake",
// 	});

// 	const params = {
// 		headers: {
// 			"Content-Type": "application/json",
// 		},
// 	};

// 	const response = http.post(url, payload, params);

// 	// Check if the response status is 200
// 	check(response, {
// 		"is status 200": (r) => r.status === 200,
// 	});
// };

import http from "k6/http";
import { check, sleep } from "k6";

// export const options = {
// 	vus: 1,
// 	duration: "10s",
// };

export let options = {
	stages: [
		{ duration: "30s", target: 150 }, // Ramp up to 10 users in 30 seconds
		{ duration: "1m", target: 1000 }, // Stay at 10 users for 1 minute
		{ duration: "30s", target: 1000 }, // Stay at 10 users for 1 minute
		// { duration: "30s", target: 20 }, // Ramp up to 20 users in 30 seconds
		// { duration: "1m", target: 20 }, // Stay at 20 users for 1 minute
		// { duration: "30s", target: 0 }, // Ramp down to 10 users in 30 seconds
		// { duration: "30s", target: 0 }, // Ramp down to 0 users in 30 seconds
	],
	// vus: 500,
	// duration: "20s",
	// thresholds: {
	//   'http_req_duration': ['p(95)<200'], // 95% of requests must complete in under 200ms
	//   'http_req_failed': ['rate<0.01'],   // Failure rate must be less than 1%
	// },
};

export default () => {
	// const url = "http://139.59.108.119:8081/srm/linckapiv1/auth/login"; // Replace with your actual URL
	const url =
		"http://localhost:8080/srm/linckapiv1/auth/login"; // Replace with your actual URL
	const payload = JSON.stringify({
		username: "karthik",
		password: "cartake",
	});

	const params = {
		headers: {
			"Content-Type": "application/json",
			"Authorization": "temp_token"
		},
	};

	const response = http.post(url, payload, params);
	// const response = http.get(url);

	// Check if the response status code is 200
	// const checkStatus = check(response, {
	// 	"is status 200": (r) => r.status === 200,
	// });
	// Check if the response body matches the expected structure
	// const checkBody = check(response, {
	// 	// "response body contains jwt_token": (r) => {
	// 	// 	const responseBody = JSON.parse(r.body);
	// 	// 	return (
	// 	// 		responseBody.status_code === 200
	// 	// 		//  &&
	// 	// 		// responseBody.jwt_token === "Login successful"
	// 	// 	);
	// 	// },
	// 	console.log(`Response body: ${r.body}`)
	// 	"response status is 200": (r) => r.status === 200,
	// });
	const checkBody = check(response, {
		"response status is 200": (r) => {
			const isSuccess = r.status === 200;
			if(!isSuccess) {
				console.log(`Response body: ${r.body}`);
			}

			return isSuccess;
		},
	});

	// const checkBody = check(response, {
	// 	"response body contains success status": (r) => {
	// 		const responseBody = JSON.parse(r.body);
	// 		return responseBody.status === "success";
	// 	},
	// });

	// Log the results of the checks
	// if (!checkStatus) {
	// 	console.log("Status check failed");
	// }
	// if (!checkBody) {
	// 	console.log("Body check failed");
	// }
};
