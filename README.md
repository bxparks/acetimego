# AceTimeGo

[![Go Tests](https://github.com/bxparks/AceTimeGo/actions/workflows/verify.yml/badge.svg)](https://github.com/bxparks/AceTimeGo/actions/workflows/verify.yml)

A date, time, and timezone library in Go lang targeting bare-metal
microcontroller environments supported by the
[TinyGo](https://github.com/tinygo-org/tinygo) compiler. All ~600 timezones
defined by the [IANA TZ database](https://github.com/eggert/tz) are supported
from the year ~3 until the year 10000. The library is self-contained and does
not depend on external files from the OS.

To reduce RAM memory consumption, the TZDB is parsed and compiled as binary data
into `const string` variables consuming approximately 35 kB of flash memory for
the years `[2000,10000)` and about 72 kB for the entire TZDB from `[1844,2087]`.
To further reduce flash memory consumption, the library does not depend on the
standard [time](https://pkg.go.dev/time) package nor the
[fmt](https://pkg.go.dev/fmt) package.

This library implements the algorithms from the
[AceTime](https://github.com/bxparks/AceTime),
[AceTimePython](https://github.com/bxparks/AceTimePython), and
[AceTimeC](https://github.com/bxparks/AceTimeC) libraries.

**Version**: 0.2.0 (2023-02-13, TZDB version 2022g)

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
Discussions](https://github.com/bxparks/AceTimeGo/discussions) for this project.
If you have bug reports, please file a ticket in [GitHub
Issues](https://github.com/bxparks/AceTimeGo/issues). Feature requests should go
into Discussions first because they often have alternative solutions which are
useful to remain visible, instead of disappearing from the default view of the
Issue tracker after the ticket is closed.

Please refrain from emailing me directly unless the content is sensitive. The
problem with email is that I cannot reference the email conversation when other
people ask similar questions later.

<a name="Authors"></a>
## Authors

* Created by Brian T. Park (brian@xparks.net).
