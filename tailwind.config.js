/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./public/**/*.html", "./src/**/*.html"],
  theme: {
    extend: {},
    fontFamily: {
      serif: ["EB Garamond", "ui-serif", "Georgia", "serif"],
      mono: ["IBM Plex Mono", "ui-monospace", "SFMono-Regular"],
      display: ["Modern Antiqua", "serif"],
    },
    backgroundImage: {
      main: 'url("./background.jpg")',
    },
  },
  plugins: [require("@tailwindcss/typography")],
};
