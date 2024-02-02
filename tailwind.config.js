import plugin from 'tailwindcss'


/** @type {import('tailwindcss').Config} */
export default {
  content: ["./views/**/*.templ"],
  darkMode: 'media',
  theme: {
    extend: {
      colors: {
        'telekom': "#ea0a8e",
      }
    },
  },
  plugins: [
    // plugin(function ({ addVariant }) {
    //   throw new Error('test')
    //   addVariant('htmx-settling', ['&.htmx-settling', '.htmx-settling &'])
    //   addVariant('htmx-request', ['&:htmx-request', '.htmx-request &'])
    //   addVariant('htmx-swapping', ['&.htmx-swapping', '.htmx-swapping &'])
    //   addVariant('htmx-added', ['&.htmx-added', '.htmx-added &'])
    // }),
    require('tailwind-htmx'),

    require('flowbite/plugin')({
      charts: true,
    }),
  ],
}

