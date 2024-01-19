import { getPets } from '$lib/services.js';
export const load = async () => {
	console.log('home page ran');
	try {
		return {
			contents: await getPets()
		};
	} catch (error) {
		console.error(error);
	}
};
