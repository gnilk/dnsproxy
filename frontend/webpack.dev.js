const webpack = require('webpack');
const merge = require('webpack-merge');
const common = require('./webpack.common.js');
const VueLoaderPlugin = require('vue-loader/lib/plugin');

module.exports = merge(common, {
    mode: 'development',

    resolve: {
        alias: {
            'vue': 'vue/dist/vue.esm.js'   // this is a test
        }
    },
    plugins: [
        new VueLoaderPlugin(),
        new webpack.DefinePlugin({
            'process.env.NODE_ENV': JSON.stringify('dev')      
        })
    ],

    externals: {
        'Config': JSON.stringify(require('./config.dev.json'))
    },

    devtool: 'inline-source-map',
    devServer: {
        contentBase: common.output.path,
        proxy: {
            '/': {
                target: 'http://localhost:3080',
                secure: false
            }
          }
        }
});