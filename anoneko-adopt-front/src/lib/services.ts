import { error } from '@sveltejs/kit';
import Axios from 'axios';

const axios = Axios.create({
	baseURL: 'http://localhost:8000/api/'
});

export const getPets = async () => {
	try {
		const result = await axios('pets/');
		return result.data.mysfits;
	} catch (e) {
		throw error(500, {
			message: 'an error occurred, refresh the page and try again'
		});
	}
};
