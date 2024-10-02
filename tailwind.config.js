/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./tailwind-test/**/*.{html,js}"],
  theme: {
    extend: {},
    corePlugins: {
      textOpacity: false,
    backgroundOpacity: false,
    borderOpacity: false,
    divideOpacity: false,
    placeholderOpacity: false,
    ringOpacity: false,
    }
  },
  plugins: [],
}

