// gzip压缩
const CompressionPlugin = require('compression-webpack-plugin');
let proxyObj = {};

proxyObj['/'] = {
    ws: false,
    target: 'http://localhost:8090/',
    changeOrigin: true,
    pathRewrite: {
        '^/': ''
    }
};
module.exports = {
    publicPath: process.env.NODE_ENV === 'production' ? './' : '/', // 构建好的文件输出到哪里
    runtimeCompiler: false,
    productionSourceMap: false,
    devServer: {
        host: 'localhost',
        stats: 'errors-only', // 打包日志输出输出错误信息
        port: 8081,
        proxy: proxyObj,
        disableHostCheck: false,
    },
    configureWebpack: {
        // 修改打包后的css文件名称
        plugins: [
            new CompressionPlugin({
                algorithm: 'gzip', // 使用gzip压缩
                test: /\.js$|\.html$|\.css$/, // 匹配文件名
                filename: '[path].gz[query]', // 压缩后的文件名(保持原文件名，后缀加.gz)
                minRatio: 1, // 压缩率小于1才会压缩
                threshold: 10240, // 对超过10k的数据压缩
                deleteOriginalAssets: false, // 是否删除未压缩的源文件，谨慎设置，如果希望提供非gzip的资源，可不设置或者设置为false（比如删除打包后的gz后还可以加载到原始资源文件）
            }),
        ],
    },
};
