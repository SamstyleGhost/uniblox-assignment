/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        background: "#EAE7DC",
        text: "#0B0308",
        primary: "#E85A4F",
        secondary: "#8E8D8A",
        accent: "#D8C3A5"
      },
      maxWidth: {
        page: "1800px",
      }
    },
  },
  plugins: [],
}

