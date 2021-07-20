let proxyObj = {};
proxyObj['/ws'] = {
    ws: true,
    target: "ws://localhost:8090"
};
proxyObj['/'] = {
    ws: false,
    target: 'http://localhost:8090',
    changeOrigin: true,
    pathRewrite: {
        '^/': ''
    }
};
module.exports = {
    devServer: {
        host: 'localhost',
        port: 8081,
        proxy: proxyObj,
        disableHostCheck: true,
    }
};