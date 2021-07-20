module.exports = {
  presets: [
    '@vue/app',
    ['@babel/preset-env', {
      modules: false
    }]
  ],
  plugins: [
    [
      "component", {
      'libraryName': "umy-ui",
      "styleLibraryName": "theme-chalk"
    }, "umy-ui"],
    [
      "component", {
      "libraryName": "element-ui",
      "styleLibraryName": "theme-chalk"
    }, "element-ui"],
  ]
};
