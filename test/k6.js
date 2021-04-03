import http from 'k6/http';
import { sleep, check } from 'k6';
import { Counter } from 'k6/metrics';
// A simple counter for http requests
export const requests = new Counter('http_reqs');
// you can specify stages of your test (ramp up/down patterns) through the options object
// target is the number of VUs you are aiming for
export const options = {
	vus: 50,
	stages: [ { target: 100, duration: '15s' } ],
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
	const res = http.post('http://0.0.0.0:8080/apis/v1/spotOrder', data, {
		headers: {
			'Content-Type': 'application/json',
			authorization:
				'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwibmFtZSI6ImhhcnZleSIsImVtYWlsIjoiIiwicGhvbmUiOiIiLCJleHBpcmVkX2F0IjoiMDAwMS0wMS0wMVQwMDowMDowMFoiLCJleHAiOjE2NTMxNjk1NDV9.Qd5blIl_RzdHE2km7_-2e0ZNX4DXY_Bvow3CSAmH5DM'
		}
	});
	//console.log(res.body)
	const checkRes = check(res, {
		'status is 200': (r) => r.status === 200
	});
}
