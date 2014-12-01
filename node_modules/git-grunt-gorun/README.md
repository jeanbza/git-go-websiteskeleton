# git-grunt-gorun v0.2.3

[![Build Status](https://travis-ci.org/jadekler/git-grunt-gorun.svg)](https://travis-ci.org/jadekler/git-grunt-gorun)

> Run Go programs



## Getting Started
This plugin requires Grunt `~0.4.0`

If you haven't used [Grunt](http://gruntjs.com/) before, be sure to check out the [Getting Started](http://gruntjs.com/getting-started) guide, as it explains how to create a [Gruntfile](http://gruntjs.com/sample-gruntfile) as well as install and use Grunt plugins. Once you're familiar with that process, you may install this plugin with this command:

```shell
npm install git-grunt-gorun --save-dev
```

Once the plugin has been installed, it may be enabled inside your Gruntfile with this line of JavaScript:

```js
grunt.loadNpmTasks('git-grunt-gorun');
```

*This plugin was designed to work with Grunt 0.4.x. If you're still using grunt v0.3.x it's strongly recommended that [you upgrade](http://gruntjs.com/upgrading-from-0.3-to-0.4), but in case you can't please use [v0.3.2](https://github.com/gruntjs/grunt-contrib-cssmin/tree/grunt-0.3-stable).*



## Gorun task
_Run this task with the `grunt gorun` command._

Run Go programs inline as a Grunt task.
### Options

###### src
Set to your go file with func main().

###### flags
Hash of runtime flags.
### Usage Examples

#### Basic running

```javascript
gorun: {
  first: {
    src: "main.go"
  }
}
```

#### Running with flags

```javascript
gorun: {
  first: {
    src: "main.go"
  },
  flags: {
    boom: "bam",
    foo: 5
  }
}
```

## Release History

 * 2014-11-24   v0.2.3   Actually removed README from package.json.
 * 2014-11-24   v0.2.2   Removed README from package.json.
 * 2014-11-24   v0.2.1   Fixed package.json to point to correct repo.
 * 2014-11-23   v0.2.0   Added support for run time flags.
 * 2014-10-13   v0.1.0   Initial commit with basic functionality - Go Run.

---

Task submitted by [Jean de Klerk](jeandeklerk.com)

*This file was generated on Mon Nov 24 2014 08:03:13.*
