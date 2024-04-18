/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./view/**/*.html"],
  theme: {
    extend: {
      colors: {
        "main-red": "#C8102E",
      },
      'backgroundImage':{
        'pattern': "url('/static/img/pattern.png')",
      }
    },
  },
  plugins: [],
}

