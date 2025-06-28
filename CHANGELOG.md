# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.0.2] — More features

### Added

- `Adapt*` to adapt a function to the given signature.
- `Lazy*` for lazy evaluation of function arguments.
- `Head*` to make partial currying.
- `Wrap` to build a chain of data processors.
- `seq` package, that provides some sequence processing functions.
- `Signature*` and `FnSig*` to compute function parameters and return types without using reflection.
- `Eq` and `Not` predicates.
- `Return*` to return the passed argument.
- `Thunk*` makes a thunk function that returns the passed argument.

### Changed

- Generated functions up to 5 parameters.
- Renamed `Curry*` functions to `F*`.
- Renamed function suffixes from `Two`, `Three`, etc. to `2`, `3`, etc.

## [0.0.1] — Initial release

### Added

- Basic functionality and unit tests.
- GitHub setup.
- README.md and LICENSE.
