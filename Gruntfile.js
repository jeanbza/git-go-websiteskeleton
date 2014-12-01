module.exports = function(grunt) {
    // Project configuration.
    grunt.initConfig({
        pkg: grunt.file.readJSON('package.json'),
        watch: {
            scripts: {
                files: ['**/*.go'],
                tasks: ['gostop', 'gorun'],
                options: {
                    spawn: false,
                },
            },
        },
        gorun: {
            app: {
                src: "main.go"
            }
        },
        gostop: {
            placeholder: {}
        }
    });

    // Load the plugin that provides the "uglify" task.
    grunt.loadNpmTasks('grunt-contrib-watch');
    grunt.loadNpmTasks('git-grunt-gorun');
    grunt.loadNpmTasks('git-grunt-gostop');

    // Default task(s).
    grunt.registerTask('default', ['gorun', 'watch']);
};