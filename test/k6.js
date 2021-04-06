import http from 'k6/http';
import { sleep, check } from 'k6';
import { Counter } from 'k6/metrics';
// A simple counter for http requests
export const requests = new Counter('http_reqs');
// you can specify stages of your test (ramp up/down patterns) through the options object
// target is the number of VUs you are aiming for
export const options = {
	vus: 150,
	stages: [ { target: 200, duration: '30s' } ],
	//thresholds: {
	//	requests: [ 'count < 100' ]
	//}
};

function getCard() {
	return Math.floor(Math.random() * 3 + 1);
}

function getTradeSide() {
	return Math.floor(Math.random() * 2 + 1);
}

function getAmount() {
	return Math.floor(Math.random() * 9 + 1);
}

export default function() {
	// our HTTP request, note that we are saving the response to res, which can be accessed later
	let reqBody = {
		card_id: getCard(),
		trade_side: getTradeSide(),
		expected_amount: getAmount(),
		card_quantity: 1
	};
	let data = JSON.stringify(reqBody);
	// let url = 'http://localhost/apis/v1/spotOrder'
	let url = 'http://0.0.0.0:9090/apis/v1/spotOrder'
	const res = http.post(url, data, {
		headers: {
			'Content-Type': 'application/json',
			authorization:
				'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwibmFtZSI6ImhhcnZleSIsImVtYWlsIjoiIiwicGhvbmUiOiIiLCJleHBpcmVkX2F0IjoiMDAwMS0wMS0wMVQwMDowMDowMFoiLCJleHAiOjE2NTM1MjQwMjJ9.b0bHWRWgGQKITOg0X3Vjuqq3Fh3OmOG9ornf5z_CT3M'
		}
	});
	//console.log(res.body)
	const checkRes = check(res, {
		'status is 200': (r) => r.status === 200
	});
}
