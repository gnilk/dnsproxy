var path = require("path");

var BUILD_DIR = path.resolve(__dirname, './public/');
var APP_DIR = path.resolve(__dirname, './src/');

var config = {

    entry: {
      ui: APP_DIR + '/public_main.js',
    },
    output: {
        path: BUILD_DIR,
        filename: '[name].bundle.js',
        publicPath: '',
    },

    resolve: {
        extensions: [".vue", ".js" ]
    },
    module: {
      rules:[
          {
            test : /\.vue$/,
            include: APP_DIR,
            exclude: /node_modules/,
            loader: 'vue-loader'
          },
          {
            test: /\.js$/,
            include: APP_DIR,
            exclude: /node_modules/,
            loader: 'babel-loader',
          },
          {
            test: /\.css$/,
            use: [
              'style-loader',
              'css-loader'
              ],
          },
          {
            test: /\.scss$/,
            use: [
              'style-loader',
              {
                loader: 'css-loader',
                options: {
                  importLoaders: 1
                }
              },
              {
                loader: 'sass-loader',
                options: {
                  indentSyntax: true
                }
              }
            ]
          },
          // {
          //   test: /\.(jpe?g|png|gif|svg|woff|woff2)$/i,
          //   use: ['url-loader?limit=10000','img-loader','url-loader']
          // },
          {
            test: /\.(png|jpg|gif|svg|eot|ttf|woff|woff2)$/,
            loader: 'url-loader',
            options: {
              limit: 10000
            }
          }
        ]
    },
};

module.exports = config;