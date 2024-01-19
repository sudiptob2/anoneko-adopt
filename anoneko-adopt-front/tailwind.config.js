/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,svelte,js,ts}'],
	daisyui: {
		themes: ['bumblebee']
	},

	plugins: [require('daisyui')]
};
