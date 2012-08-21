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

The MIT License

Copyright (c) 2012 Christoph Hochstrasser, http://christophh.net

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE. 

