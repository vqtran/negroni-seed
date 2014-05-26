var gulp = require('gulp');

var include = require('gulp-include');
var concat = require('gulp-concat');
var uglify = require('gulp-uglify');
var minify = require('gulp-minify-css');
var less = require('gulp-less');

var paths = {
  scripts: ['assets/javascripts/*.js'],
  less: ['assets/stylesheets/*.less',
         'assets/stylesheets/bootstrap/bootstrap.less'
        ]
};

gulp.task('scripts', function() {
  // Minify and copy all JavaScript (except vendor scripts)
   return gulp.src(paths.scripts)
              .pipe(include())
              .pipe(uglify())
              .pipe(concat('all.min.js'))
              .pipe(gulp.dest('public/assets'));
});

gulp.task('less', function() {
   return gulp.src(paths.less)
              .pipe(less({
                 paths: [ 'assets/stylesheets' ]
              }))
              .pipe(minify())
              .pipe(concat('all.min.css'))
              .pipe(gulp.dest('public/assets'));
});

// Rerun the task when a file changes
gulp.task('watch', function() {
  gulp.watch(paths.scripts, ['scripts']);
  gulp.watch(paths.less, ['less']);
});

// The default task (called when you run `gulp` from cli)
gulp.task('default', ['scripts', 'less', 'watch']);
