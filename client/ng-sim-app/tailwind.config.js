/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ["./src/**/*.{html,ts}", "./projects/**/*.{html,ts}"],
	theme: {
		extend: {
			fontFamily: {
				poppins: "Poppins, 'sans-serif'",
			},
			colors: {
				dark: ["#0C0B10", "#222128", "#2E2D33"],
				light: ["#fff", "#F5F5F5", "#97979B"],
				primary: { neutral: "#5FC0B6" },
			},
		},
	},
	plugins: [],
};
