import type { Config } from 'tailwindcss';

export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],

	theme: {
		extend: {
			colors:{
				header: '#1E293B',
				background: '#0F172A',
				delete: '#E82222'
			}
		}
	},

	plugins: []
} satisfies Config;
