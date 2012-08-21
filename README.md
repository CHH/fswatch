# fswatch

_Watch a directory, run a command when something changes._

## Install

Via `go get`:

	% go get github.com/CHH/fswatch

You should now have a `fswatch` executable. Run `fswatch -h` to get all
options.

## Usage

The `fswatch` utility watches the current working directory and runs the
command line given as arguments, each time a file is changed.

Some possible use cases:

- Recompile [LESS](http://lesscss.org) files everytime you change them.
- Run the project's tests.

To run `make test` everytime a file is changed:

	% fswatch make test

## Known Issues

- `fswatch` will trigger the command on every file change. You have to 
  see for yourself if the command should be run or not. Therefore it's
  recommended to use a build tool like `make` or `rake` as target for
  `fswatch`.
- You've to restart `fswatch` to make it pick up files in any created
  directories.

## License



