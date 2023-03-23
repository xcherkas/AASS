const { CleanWebpackPlugin } = require('clean-webpack-plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin')
const TerserPlugin = require('terser-webpack-plugin');
const CopyPlugin = require('copy-webpack-plugin');

const path = require('path');
const webpack = require('webpack');

module.exports = function(_, arg) {

  const config = {

    entry: {
      bundle: [ './app/main.js' ]
    },

    resolve: {
      alias: {
        svelte: path.resolve('node_modules', 'svelte')
      },
      extensions: [ '.mjs', '.js', '.svelte' ],
      mainFields: ['svelte', 'browser', 'module', 'main']
    },

    output: {
      filename: '[name].js',
      publicPath: ''
    },

    devtool: arg.mode === 'development' ? 'inline-source-map' : undefined,

    optimization: {
      minimize: arg.mode !== 'development',
      minimizer: [
        new TerserPlugin({
          cache: true,
          parallel: true,
          sourceMap: arg.mode === 'development',
          terserOptions: {
            output: {
              comments: false
            }
          }
        })
      ]
    },

    module: {
      rules: [
        {
          test: /\.(svelte)$/,
          use: {
            loader: 'svelte-loader',
            options: {
              hotReload: true
            }
          }
        },
        {
          test: /\.less$/,
          use: [
            'style-loader',
            'css-loader',
            'less-loader'
          ],
          exclude: /node_modules/
        },
        {
          test: /\.css$/,
          use: ['style-loader', 'css-loader']
        }
      ]
    },

    plugins: [

      new CleanWebpackPlugin(),

      new webpack.DefinePlugin({
        'process.env.NODE_ENV': arg.mode
      }),

      new CopyPlugin([
        { from: 'app/static/img', to: 'img/' },
        { from: 'app/static/img', to: 'client/img/' }
      ]),

      new HtmlWebpackPlugin({
        template: 'app/index.html'
      }),

      new webpack.ProvidePlugin({
        $: 'jquery',
        jQuery: 'jquery'
      }),

    ]
  };


  return config;
};