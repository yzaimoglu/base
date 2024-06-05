import presetQuick from "franken-ui/shadcn-ui/preset-quick";

/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: 'class',
  content: ["./http/**/*.{go,html,js}", "./ui/html/**"],
  theme: {
    extend: {},
  },
  important: true,
  presets: [presetQuick({ theme: "slate" })],
}
