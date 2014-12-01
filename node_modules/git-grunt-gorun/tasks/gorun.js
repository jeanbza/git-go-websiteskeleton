'use strict';

var chalk = require('chalk');
var spawn = require('child_process').spawn;

module.exports = function(grunt) {
    var opts = {
        stdout: function (data) {
            grunt.log.writeln(data);
        },
        stderr: function (data) {
            grunt.log.error(data);
        },
        detached: true,
        stdio: [ 'ignore', process.stdin, process.stdout ]
    };

    grunt.registerMultiTask('gorun', 'Run Go programs.', function() {
        var done = this.async();
        var src = this.data.src;
        var flags = this.data.flags;
        var args = ['run', src];

        for (var key in flags) {
            args.push("--"+key+"="+flags[key]);
        }

        grunt.log.writeln("go run "+chalk.cyan(src));

        setTimeout(function() {
            spawn('go', args, opts);
            done();
        }, 1000);
    });
};