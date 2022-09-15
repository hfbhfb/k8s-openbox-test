module.exports = {
    // publicPath: process.env.NODE_ENV === "production" ? "/qushuiyin" : "/",
    publicPath: process.env.NODE_ENV === "production" ? "/" : "/",
    chainWebpack: (config) => {
        config.plugin("html").tap((args) => {
            args[0].title = "获取发布视频"; // Replace your title here
            return args;
        });
    },
    css: {
        loaderOptions: {
            less: {
                javascriptEnabled: true,
            },
        },
    },
    devServer:{
        port: 1026,
        host: '0.0.0.0',
        disableHostCheck: true,
        proxy:{
            '/api':{
                target: 'http://localhost:1027',
                changeOrigin: true
            }
        }
    }
};