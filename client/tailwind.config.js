/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      fontFamily: {
        agu: ["Agu Display", "serif"],
        skyer: ["Skyer", "serif"]
      },
      colors: {
        text: "#050315",
        background: "#fbfbfe",
        primary: "#2f27ce",
        secondary: "#dedcff",
        accent: "#3de456",
      },
      maxWidth: {
        page: "1800px",
      },
    },
  },
  plugins: [],
};
