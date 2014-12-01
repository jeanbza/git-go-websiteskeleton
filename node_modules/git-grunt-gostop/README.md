# git-grunt-gostop v0.1.0

[![Build Status](https://travis-ci.org/jadekler/git-grunt-gostop.svg)](https://travis-ci.org/jadekler/git-grunt-gostop)

> Stop Running Go programs



## Getting Started
This plugin requires Grunt `~0.4.0`

If you haven't used [Grunt](http://gruntjs.com/) before, be sure to check out the [Getting Started](http://gruntjs.com/getting-started) guide, as it explains how to create a [Gruntfile](http://gruntjs.com/sample-gruntfile) as well as install and use Grunt plugins. Once you're familiar with that process, you may install this plugin with this command:

```shell
npm install git-grunt-gostop --save-dev
```

Once the plugin has been installed, it may be enabled inside your Gruntfile with this line of JavaScript:

```js
grunt.loadNpmTasks('git-grunt-gostop');
```

*This plugin was designed to work with Grunt 0.4.x. If you're still using grunt v0.3.x it's strongly recommended that [you upgrade](http://gruntjs.com/upgrading-from-0.3-to-0.4), but in case you can't please use [v0.3.2](https://github.com/gruntjs/grunt-contrib-cssmin/tree/grunt-0.3-stable).*



## Gostop task
_Run this task with the `grunt gostop` command._

Stop Go programs inline as a Grunt task.
### Options

###### src
Set to your go file with func main().
### Usage Examples

#### Basic compilation

```javascript
gostop: {
}
```

## Release History

 * 2014-10-13   v0.1.0   Initial commit with basic functionality - Go Stop.

---

Task submitted by [Jean de Klerk](jeandeklerk.com)

*This file was generated on Sat Nov 08 2014 20:09:15.*
