const BASE_URL = 'http://localhost:3000';

import axios from 'axios';

import { PUBLIC_API_KEY } from '$env/static/public';

console.log('PUBLIC_API_KEY', PUBLIC_API_KEY);

const options = {
	method: 'GET',
	url: 'https://weatherapi-com.p.rapidapi.com/current.json',
	params: { q: '53.1,-0.13' },
	headers: {
		// need key
		'X-RapidAPI-Key': PUBLIC_API_KEY,
		'X-RapidAPI-Host': 'weatherapi-com.p.rapidapi.com'
	}
};

try {
	const response = await axios.request(options);
	console.log(response.data);
} catch (error) {
	console.error(error);
}

const api = {
	get: async () => {
		try {
			const response = await axios.request(options);
			return await response.data;
		} catch (error) {
			return [];
		}
	}
};

export { api, BASE_URL };
