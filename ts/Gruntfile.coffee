module.exports = (grunt) =>
  grunt.initConfig
    typescript:
      base:
        src: ['*.ts', '**/*.ts']
        dest: '../public/js'
        options:
          target: 'es5'
          sourceMap: true
    watch:
      scripts:
        files: ['*.ts', '**/*.ts']
        tasks: ['typescript']
  grunt.loadNpmTasks('grunt-contrib-watch')
  grunt.loadNpmTasks('grunt-typescript')
  grunt.registerTask('default', ['typescript', 'watch'])
