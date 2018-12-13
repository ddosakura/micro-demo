// ref: https://umijs.org/config/
export default {
  plugins: [
    // ref: https://umijs.org/plugin/umi-plugin-react.html
    ['umi-plugin-react', {
      antd: true,
      dva: true,
      dynamicImport: true,
      title: 'frontend',
      dll: true,
      routes: {
        exclude: [],
      },
      hardSource: true,
    }],
  ],
  proxy: {
    "/search": {
      "target": "http://www.baidu.com/",
      "changeOrigin": true,
      "pathRewrite": {
        "^/search": "/s"
      }
    }
  }
}
