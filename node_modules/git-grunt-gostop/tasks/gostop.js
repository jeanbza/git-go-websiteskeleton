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
        }
    };

    grunt.registerMultiTask('gostop', 'Stop Running Go programs.', function() {
        var done = this.async();
        var commandText = "pkill -f go-build";
        process.env.GOOS="darwin";

        var proc = spawn('pkill', ["-f", "go-build"], opts);

        proc.on('stdout', function(data) {
            opts.stdout(data);
        });

        proc.on('stderr', function(data) {
            opts.stderr(data);
        });

        proc.on('error', function(err) {
            console.log(err);
        });

        proc.on('exit', function (status) {
            if (status !== 0) {
                grunt.fail.fatal("Failure executing "+commandText+" with exit code "+status+".");
            }
            
            done();
        });
    });
};