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
    require('flowbite/plugin')({
      charts: true,
    })
  ],
}

