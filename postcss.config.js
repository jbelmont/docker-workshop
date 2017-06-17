// https://github.com/michael-ciniawsky/postcss-load-config

module.exports = (ctx) => ({
  map: ctx.options.map,
  'plugins': {
    // to edit target browsers: use "browserlist" field in package.json
    'autoprefixer': {}
  }
})
