import presetQuick from "franken-ui/shadcn-ui/preset-quick";

/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: 'class',
  content: ["./ui/**/*.{templ,html,js}", "./ui/html/**"],
  theme: {
    extend: {},
  },
  presets: [presetQuick({ theme: "slate" })],
}
