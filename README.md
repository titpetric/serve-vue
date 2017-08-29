# serve-vue

Serve a built vue-js project folder

## Building

If you have go, it should be as simple as:

`go run main.go` or `go build main.go`.

If you have docker, you can build and install by running `./build.sh`.

## Usage

I'm tired of setting up nginx-es for every small vuejs project.
This lets me serve the project with a simple one-liner after it's
built.

~~~
serve-vue [folder] (-port NNN)
~~~

By default it will serve your current folder, so you should run
`serve-vue` inside your projects `src/dist/` folder where the resulting
built project is located.

What's the catch? Why specific to VueJS? Any 404 files will serve
/index.html. This enables vue-router pustState navigation without
those ugly `#` shebangs. I suppose this supports any other framework
that can use pushState routing (Angular, React, ...)

## Words

Enjoy :)