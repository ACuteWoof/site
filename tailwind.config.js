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
      main: 'url("https://lewoof.xyz/api/background")',
    },
    extend: {
      colors: {
        gruvboxDark: {
          // Official Gruvbox HARD theme
          // Background colors (from darkest to lightest)
          bg0_h: "#1d2021",   // Hard background
          bg0: "#282828",     // Default background
          bg0_s: "#32302f",   // Soft background
          bg1: "#3c3836",
          bg2: "#504945",
          bg3: "#665c54",
          bg4: "#7c6f64",

          // Foreground colors (from lightest to darkest)
          fg0: "#fbf1c7",     // Brightest foreground
          fg1: "#ebdbb2",     // Default foreground
          fg2: "#d5c4a1",
          fg3: "#bdae93",
          fg4: "#a89984",

          // Aliases for backwards compatibility
          bg: "#1d2021",      // bg0_h
          fg: "#ebdbb2",      // fg1

          // Accent colors
          red: "#cc241d",
          red2: "#fb4934",    // Bright red
          green: "#98971a",
          green2: "#b8bb26",  // Bright green
          yellow: "#d79921",
          yellow2: "#fabd2f", // Bright yellow
          blue: "#458588",
          blue2: "#83a598",   // Bright blue
          purple: "#b16286",
          purple2: "#d3869b", // Bright purple
          aqua: "#689d6a",
          aqua2: "#8ec07c",   // Bright aqua
          orange: "#d65d0e",
          orange2: "#fe8019", // Bright orange

          // Gray scale
          gray: "#928374",
        },
        gruvbox: {
          bg: "#fbf1c7",
          bg0: "#fbf1c7",
          bgH: "#f9f5d7",
          bgS: "#f2e5bc",
          bg1: "#ebdbb2",
          bg2: "#d5c4a1",
          bg3: "#bdae93",
          bg4: "#a89984",

          fg: "#3c3836",
          fg0: "#282828",
          fg1: "#3c3836",
          fg2: "#504945",
          fg3: "#665c54",
          fg4: "#7c6f64",

          red: "#cc241d",
          red2: "#9d0006",
          green: "#98971a",
          green2: "#79740e",
          yellow: "#d79921",
          yellow2: "#b57614",
          blue: "#458588",
          blue2: "#076678",
          purple: "#b16286",
          purple2: "#8f3f71",
          aqua: "#689d6a",
          aqua2: "#427b58",
          orange: "#d65d0e",
          orange2: "#af3a03",
          gray: "#7c6f64",
          gray2: "#928374",
        },
      },
      typography: ({ theme }) => ({
        DEFAULT: {
          css: {
            "--tw-prose-body": theme("colors.gruvbox.fg"),
            "--tw-prose-headings": theme("colors.gruvbox.fg"),
            "--tw-prose-links": theme("colors.gruvbox.aqua"),
            "--tw-prose-bold": theme("colors.gruvbox.fg"),
            "--tw-prose-counters": theme("colors.gruvbox.aqua"),
            "--tw-prose-bullets": theme("colors.gruvbox.aqua"),
            "--tw-prose-hr": theme("colors.gruvbox.bg"),
            "--tw-prose-quotes": theme("colors.gruvbox.fg"),
            "--tw-prose-quote-borders": theme("colors.gruvbox.aqua"),
            "--tw-prose-captions": theme("colors.gruvbox.aqua"),
            "--tw-prose-code": theme("colors.gruvbox.yellow"),
            "--tw-prose-pre-code": theme("colors.gruvbox.fg"),
            "--tw-prose-pre-bg": theme("colors.gruvbox.bg"),
            "--tw-prose-th-borders": theme("colors.gruvbox.bg"),
            "--tw-prose-td-borders": theme("colors.gruvbox.bg"),
            // Dark mode invert colors (using Gruvbox HARD theme)
            "--tw-prose-invert-body": theme("colors.gruvboxDark.fg1"),
            "--tw-prose-invert-headings": theme("colors.gruvboxDark.fg1"),
            "--tw-prose-invert-links": theme("colors.gruvboxDark.aqua"),
            "--tw-prose-invert-bold": theme("colors.gruvboxDark.fg1"),
            "--tw-prose-invert-counters": theme("colors.gruvboxDark.aqua"),
            "--tw-prose-invert-bullets": theme("colors.gruvboxDark.aqua"),
            "--tw-prose-invert-hr": theme("colors.gruvboxDark.bg0_h"),
            "--tw-prose-invert-quotes": theme("colors.gruvboxDark.fg1"),
            "--tw-prose-invert-quote-borders": theme("colors.gruvboxDark.aqua"),
            "--tw-prose-invert-captions": theme("colors.gruvboxDark.aqua"),
            "--tw-prose-invert-code": theme("colors.gruvboxDark.yellow"),
            "--tw-prose-invert-pre-code": theme("colors.gruvboxDark.fg1"),
            "--tw-prose-invert-pre-bg": theme("colors.gruvboxDark.bg0"),
            "--tw-prose-invert-th-borders": theme("colors.gruvboxDark.bg1"),
            "--tw-prose-invert-td-borders": theme("colors.gruvboxDark.bg1"),
          },
        },
      }),
    },
  },
  plugins: [require("@tailwindcss/typography")],
};
