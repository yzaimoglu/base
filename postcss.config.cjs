module.exports = {
  plugins: [
    require('cssnano')({
      preset: 'default',
    }),
    require("tailwindcss"),
    require("autoprefixer"),
    require("postcss-sort-media-queries")(),
    require("postcss-combine-duplicated-selectors")({
      removeDuplicatedProperties: true,
    }),
  ],
};

