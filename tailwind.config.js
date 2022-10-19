/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/*.html"],
  plugins: [
    // require('@tailwindcss/forms'),
    require("daisyui")
  ],
  daisyui: {
    themes: false,
  },
}
