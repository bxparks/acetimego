# acetimego - AceTime time zone library for Go and TinyGo

[![Go Tests](https://github.com/bxparks/acetimego/actions/workflows/verify.yml/badge.svg)](https://github.com/bxparks/acetimego/actions/workflows/verify.yml)

The `acetimego` library provides date, time, and timezone functionality for the
bare-metal microcontroller environments using the
[TinyGo](https://github.com/tinygo-org/tinygo) compiler. In such microcontroller
environments, the standard [go.time](https://pkg.go.dev/time) cannot be used
because there is no underlying operating system, and the `go.time` library
implementation consumes too much flash memory.

This library supports all ~600 timezones defined by the [IANA TZ
database](https://github.com/eggert/tz). The library is self-contained and does
not depend on external files from the host OS. Three versions of the TZDB are
provided in this library:

* `zonedball`
    * All timezones with transitions for all years defined by the TZDB database,
      from the year 1844 onwards,
    * Consumes  about 72 kB of flash memory.
* `zonedb`
    * All timezones with transitions for the years 2000 and onwards,
    * Consumes about 35 kB of flash memory.
* `zonedbtesting`
    * A small subset of timezones for internal testing purposes.

To reduce RAM memory consumption, the TZDB is parsed and compiled into binary
data encoded as `const string` variables, which allows the TinyGo compiler to
place the data structures into flash memory instead of static/dynamic RAM. To
reduce flash memory consumption even further, the library does not depend on the
standard [time](https://pkg.go.dev/time) package nor the
[fmt](https://pkg.go.dev/fmt) package.

This library implements the algorithms equivalent to the following libraries:

* [AceTime](https://github.com/bxparks/AceTime) for Arduino,
* [acetimepy](https://github.com/bxparks/acetimepy) for Python,
* [acetimec](https://github.com/bxparks/acetimec) for C.

**Version**: 0.5.2 (2024-07-26, TZDB 2024a)

**Changelog**: [CHANGELOG.md](CHANGELOG.md)

## Table of Contents

<a name="Example"></a>
## Example

<a name="Installation"></a>
## Installation

<a name="License"></a>
## License

[MIT License](https://opensource.org/licenses/MIT)

<a name="FeedbackAndSupport"></a>
## Feedback and Support

If you have any questions, comments, or feature requests for this library,
please use the [GitHub
Discussions](https://github.com/bxparks/acetimego/discussions) for this project.
If you have bug reports, please file a ticket in [GitHub
Issues](https://github.com/bxparks/acetimego/issues). Feature requests should go
into Discussions first because they often have alternative solutions which are
useful to remain visible, instead of disappearing from the default view of the
Issue tracker after the ticket is closed.

Please refrain from emailing me directly unless the content is sensitive. The
problem with email is that I cannot reference the email conversation when other
people ask similar questions later.

<a name="Authors"></a>
## Authors

* Created by Brian T. Park (brian@xparks.net).
