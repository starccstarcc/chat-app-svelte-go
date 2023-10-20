import { PUBLIC_API_KEY } from '$env/static/public';
import axios from 'axios';

const BASE_URL = `https://api.weatherapi.com/v1/current.json?key=${PUBLIC_API_KEY}`;

type weatherApiOptions = {
	city: string;
};

function generateUrl({ city }: weatherApiOptions) {
	return `${BASE_URL}&q=${city}&aqi=no`;
}

const api = {
	get: async ({ city }: weatherApiOptions) => {
		try {
			const options = {
				method: 'GET',
				url: generateUrl({ city })
			};
			const response = await axios.request(options);
			return await response.data;
		} catch (error) {
			console.log(`ERROR API CALL : `, error);
			return {};
		}
	}
};

export { api, BASE_URL };
