/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: ["class"],
  content: [
    './pages/**/*.{ts,tsx}',
    './components/**/*.{ts,tsx}',
    './app/**/*.{ts,tsx}',
    './src/**/*.{ts,tsx}',
  ],
  prefix: "",
  theme: {
  	container: {
  		center: 'true',
  		padding: '2rem',
  		screens: {
  			'2xl': '1400px'
  		},
		extend: {
		colors: {
			  // ... existing colors
			  'primary-light': 'hsl(var(--primary-light))',
			  'primary-200': 'hsl(var(--primary-200))',
			  'primary-800': 'hsl(var(--primary-800))',
			  
			  'secondary-light': 'hsl(var(--secondary-light))',
			  'secondary-200': 'hsl(var(--secondary-200))',
			  'secondary-800': 'hsl(var(--secondary-800))',
			  
			  success: {
				light: 'hsl(var(--success-light))',
				DEFAULT: 'hsl(var(--success))',
				dark: 'hsl(var(--success-dark))',
				200: 'hsl(var(--success-200))',
			  },
			  
			  error: {
				light: 'hsl(var(--error-light))',
				DEFAULT: 'hsl(var(--error))',
				dark: 'hsl(var(--error-dark))',
			  },
			  
			  orange: {
				light: 'hsl(var(--orange-light))',
				DEFAULT: 'hsl(var(--orange))',
				dark: 'hsl(var(--orange-dark))',
			  },
			  
			  warning: {
				light: 'hsl(var(--warning-light))',
				DEFAULT: 'hsl(var(--warning))',
				dark: 'hsl(var(--warning-dark))',
			  },
			  
			  grey: {
				50: 'hsl(var(--grey-50))',
				100: 'hsl(var(--grey-100))',
				200: 'hsl(var(--grey-200))',
				300: 'hsl(var(--grey-300))',
				500: 'hsl(var(--grey-500))',
				600: 'hsl(var(--grey-600))',
				700: 'hsl(var(--grey-700))',
				900: 'hsl(var(--grey-900))',
			  },
			  
			  // Dark theme specific colors
			  'dark-paper': 'hsl(var(--paper))',
			  'dark-level-1': 'hsl(var(--level-1))',
			  'dark-level-2': 'hsl(var(--level-2))',
			  'dark-text': {
				title: 'hsl(var(--text-title))',
				primary: 'hsl(var(--text-primary))',
				secondary: 'hsl(var(--text-secondary))',
			  },
			},
		  },
  	},
  	extend: {
  		colors: {
  			border: 'hsl(var(--border))',
  			input: 'hsl(var(--input))',
  			ring: 'hsl(var(--ring))',
  			background: 'hsl(var(--background))',
  			foreground: 'hsl(var(--foreground))',
  			primary: {
  				DEFAULT: 'hsl(var(--primary))',
  				foreground: 'hsl(var(--primary-foreground))'
  			},
  			secondary: {
  				DEFAULT: 'hsl(var(--secondary))',
  				foreground: 'hsl(var(--secondary-foreground))'
  			},
  			destructive: {
  				DEFAULT: 'hsl(var(--destructive))',
  				foreground: 'hsl(var(--destructive-foreground))'
  			},
  			muted: {
  				DEFAULT: 'hsl(var(--muted))',
  				foreground: 'hsl(var(--muted-foreground))'
  			},
  			accent: {
  				DEFAULT: 'hsl(var(--accent))',
  				foreground: 'hsl(var(--accent-foreground))'
  			},
  			popover: {
  				DEFAULT: 'hsl(var(--popover))',
  				foreground: 'hsl(var(--popover-foreground))'
  			},
  			card: {
  				DEFAULT: 'hsl(var(--card))',
  				foreground: 'hsl(var(--card-foreground))'
  			},
  			chart: {
  				'1': 'hsl(var(--chart-1))',
  				'2': 'hsl(var(--chart-2))',
  				'3': 'hsl(var(--chart-3))',
  				'4': 'hsl(var(--chart-4))',
  				'5': 'hsl(var(--chart-5))'
  			}
  		},
  		borderRadius: {
  			lg: 'var(--radius)',
  			md: 'calc(var(--radius) - 2px)',
  			sm: 'calc(var(--radius) - 4px)'
  		},
  		keyframes: {
  			'accordion-down': {
  				from: {
  					height: '0'
  				},
  				to: {
  					height: 'var(--radix-accordion-content-height)'
  				}
  			},
  			'accordion-up': {
  				from: {
  					height: 'var(--radix-accordion-content-height)'
  				},
  				to: {
  					height: '0'
  				}
  			}
  		},
  		animation: {
  			'accordion-down': 'accordion-down 0.2s ease-out',
  			'accordion-up': 'accordion-up 0.2s ease-out'
  		}
  	}
  },
  plugins: [require("tailwindcss-animate")],
}