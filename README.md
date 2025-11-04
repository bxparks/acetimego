# acetimego - a tiny time zone library for Go and TinyGo

[![Go Tests](https://github.com/bxparks/acetimego/actions/workflows/verify.yml/badge.svg)](https://github.com/bxparks/acetimego/actions/workflows/verify.yml)

The `acetimego` library provides date, time, and timezone functionality for the
bare-metal microcontroller environments using the
[TinyGo](https://github.com/tinygo-org/tinygo) compiler. In such microcontroller
environments, the standard [Go time](https://pkg.go.dev/time) package cannot be
used because there is no underlying operating system, and the Go `time` library
implementation consumes too much flash memory.

This library supports all ~600 timezones defined by the [IANA TZ
database](https://github.com/eggert/tz). The library is self-contained and does
not depend on external files from the host OS. Three versions of the TZDB are
provided in this library:

- `zonedb2000`
    - All timezones with transitions for the years 2000 and onwards,
    - Consumes about 35 kB of flash memory.
- `zonedb2025`
    - All timezones with transitions for the years 2025 and onwards,
    - Consumes about 26 kB of flash memory.
- `zonedball`
    - All timezones with transitions for all years defined by the TZDB database,
      from the year 1844 onwards,
    - Consumes  about 72 kB of flash memory.

To reduce RAM memory consumption, the TZDB is parsed and compiled into binary
data encoded as `const string` variables, which allows the TinyGo compiler to
place the data structures into flash memory instead of static/dynamic RAM. To
reduce flash memory consumption even further, the library does not depend on the
standard [time](https://pkg.go.dev/time) package nor the
[fmt](https://pkg.go.dev/fmt) package.

This library implements the algorithms equivalent to the following libraries:

- [AceTime](https://github.com/bxparks/AceTime) for Arduino,
- [acetimepy](https://github.com/bxparks/acetimepy) for Python,
- [acetimec](https://github.com/bxparks/acetimec) for C.

**Version**: 0.8.0 (2025-10-21, TZDB 2025b)

**Changelog**: [CHANGELOG.md](CHANGELOG.md)

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
    - [Epoch](#epoch)
    - [UnixSeconds](#unixseconds)
    - [PlainDate](#plaindate)
    - [ISO Weekday](#iso-weekday)
    - [PlainDateTime](#plaindatetime)
    - [TimeZone](#timezone)
    - [ZoneManager](#zonemanager)
    - [ZonedDateTime](#zoneddatetime)
        - [Disambiguate Gaps and Overlaps](#disambiguate-gaps-and-overlaps)
        - [Resolved Gaps and Overlaps](#resolved-gaps-and-overlaps)
        - [Convert TimeZone](#convert-timezone)
        - [ZonedDateTime Normalization](#zoneddatetime-normalization)
    - [ZonedExtra](#zonedextra)
- [Bugs and Limitations](#bugs-and-limitations)
- [License](#license)
- [Feedback and Support](#feedback-and-support)
- [Authors](#authors)

## Installation

The main package of the `acetimego` library is the `acetime` package:

```go
import (
  "github.com/bxparks/acetimego/acetime"
)
```

There are 3 database packages which allows the end user to select the range of
validity of the TZDB, which directly affects the size of the final binary.
Normally, an application will choose only of the following:

```go
import (
  "github.com/bxparks/acetimego/zonedb2000"
  "github.com/bxparks/acetimego/zonedb2025"
  "github.com/bxparks/acetimego/zonedball"
)
```

- `zonedb2000` contains data for all timezones in the TZDB, but restricted to
  the year 2000 and onwards to reduce the size. The database is approximately 35
  kB.
- `zonedb2025` contains data for all timezones in the TZDB, but restricted to
  the year 2025 and onwards to reduce the size. The database is approximately 26
  kB.
- `zonedball` contains the entire TZDB, for all timezones, for all years in the
  TZDB from 1844 and onwards. The database size is approximately 72 kB.

The `zoneinfo` package:

```go
import (
  "github.com/bxparks/acetimego/zoneinfo"
)
```

will not normally be used by the end-user. It is the package that knows how to
parse and traverse the `zonedb*` database files.

## Usage

The `acetimego` library does not use the standard `go.time` library to conserve
space on microcontroller environments.

### Epoch

Like many other date-timezone libraries, we define a specific instant of time
relative to the Unix Epoch which is defined as 1970-01-01 00:00:00 UTC. The
UnixSeconds is the number of seconds relative to this Unix Epoch.

Also like other libraries, we measure the number of seconds in units of [POSIX
seconds](https://en.wikipedia.org/wiki/Unix_time) instead of an [SI
second](https://en.wikipedia.org/wiki/Second). Unlike the SI second, the POSIX
second is not constant in duration. For example, one POSIX second is equivalent
to 2 SI seconds during a [leap
second](https://en.wikipedia.org/wiki/Leap_second). This definition allows one
POSIX day to have exactly 86400 POSIX seconds, regardless of leap seconds, which
makes many internal calculations a lot simpler.

### UnixSeconds

This library defines an `acetime.Time` type to measure the unixSeconds. It is
defined as a 64-bit signed integer `int64`. This is vastly simpler than the
`time.Time` type in the standard library which is a [struct of 3
fields](https://go.googlesource.com/proposal/+/master/design/12914-monotonic.md#time-representation)
which are 24-bytes in size on 64-bit systems, and 20 bytes on 32-bit systems.

Other AceTime-related libraries (e.g. AceTime, acetimec) use a 32-bit signed
integer for the unixSeconds to save flash and volatile memory. However, I
assumed that the target environments for `acetimego` and TinyGo are
microcontrollers which are at least 32-bits wide, which can handle 64-bit
integers efficiently. This simplifies the API of the `acetimego` library, and
avoids the [Year 2038](https://en.wikipedia.org/wiki/Year_2038_problem) problem
which can affect systems that use a 32-bit integer for the unixSeconds.

### PlainDate

A PlainDate represents a date in the [proleptic Gregorian
calendar](https://en.wikipedia.org/wiki/Proleptic_Gregorian_calendar) which has
a 400-year cycle, with leap days occurring roughly every 4 years to account for
the difference between the earth's rotation day and the earth's revolution year
around the sun.

For flexibility and efficiency, we don't define a `PlainDate` object in this
library. Rather, we define utilities functions which accept or return the
`year`, `month`, and `day` parameters separately:

```go
  isLeap := acetime.IsLeapYear(2050) // returns false

  daysInYearMonth := acetime.DaysInYearMonth(2050, 2) // returns 28

  unixDays := acetime.PlainDateToUnixDays(2050, 1, 1) // returns 29220

  year, month, day := acetime.PlainDateFromUnixDays(29220)
```

### ISO Weekday

The `acetimego` library defines an ISO Weekday type:

```go
type IsoWeekday uint8
```

The ISO week is slightly different than the Go `time.Weekday` type because the
ISO week starts on Monday with a value of 1, and ends on Sunday with a value of
7.

The ISO weekday can be retrieved from the PlainDate using the
`PlainDateToWeekday()` function:

```go
  weekDay := acetime.PlainDateToWeekday(2050, 1, 1) // returns acetime.Saturday

  weekDayString := weekDay.Name() // returns "Saturday"
```

### PlainDateTime

The `PlainDateTime` represents a date-time instance (year, month, day, hour,
minute, second) with no information about a timezone. Sometimes it represents a
date-time in the local time zone, sometimes it represents a date-time in UTC,
depending on context.

To create an object that represents `2050-01-01 00:00:01`, we use

```go
  pdt := acetime.PlainDateTime{2050, 1, 1, 0, 0, 1}
```

The 2 important things you can do with a `PlainDateTime` is to convert it to
unixSeconds and back:

```go
  pdt1 := acetime.PlainDateTime{2050, 1, 1, 0, 0, 1}
  unixSeconds := pdt.UnixSeconds() // returns acetime.Time of 2524608001

  pdt2 := acetime.PlainDateTimeFromUnixSeconds(unixSeconds)
  equals12 := pdt1 == pdt2 // should return true

  pdt3 := acetime.PlainDateTimeFromUnixSeconds(Time(2524608001))
  equals13 := pdt1 == pdt3 // should return true
```

We can convert a `PlainDateTime` into a human-readable string in [ISO
8601](https://en.wikipedia.org/wiki/ISO_8601) format using the `String()`
function:

```go
  pdt := acetime.PlainDateTime{2050, 1, 1, 0, 0, 1}
  s := pdt.String() // returns "2025-01-01T00:00:01"
```

A `BuildString()` function is provided to allow incremental construction of a
String using the `strings.Builder` object:

```go
  pdt := acetime.PlainDateTime{2050, 1, 1, 0, 0, 1}
  b := strings.Builder()
  pdt.BuildString(b) // appends "2025-01-01T00:00:01" to 'b'
```

### ZoneManager

A `TimeZone` object represents a specific timezone in the TZDB. It is (almost
always) created by the `ZoneManager` object, so let's examine the `ZoneManager`
first.

The `ZoneManager` is initialized by passing the `DataContext` from a specific
`zonedb*` package, like this:

```go
import (
 "github.com/bxparks/acetimego/acetime"
 "github.com/bxparks/acetimego/zonedb2000"
)

func doSomething() {
  zm := acetime.ZoneManagerFromDataContext(&zonedb2000.DataContext)
  ...
}
```

Once the `ZoneManager` object is constructed, we can retrieve a handful of
metadata about the zone database that we selected:

```go
func doSomething() {
  zm := acetime.ZoneManagerFromDataContext(&zonedb2000.DataContext)

  zoneCount := zm.ZoneCount() // number of zones
  zoneNames := zm.ZoneNames() // list of zone names in the database
  zoneIds := zm.ZoneIDs() // list of zone identifiers in the database
}
```

### TimeZone

The `TimeZone` object represents a timezone. It is analogous to the
[time.Location](https://pkg.go.dev/time#Location) object in the standard Go
`time` package.

The `TimeZone` is almost always created by the `ZoneManager`.

A timezone in the `acetimego` library is identified in 2 ways:

- a string (e.g. "America/Los_Angeles"), or
- a `uint32` ZoneID (e.g. `zonedb2000.ZoneIDAmerica_Los_Angeles`)

The ZoneID integer identifier is unique and stable across multiple versions of
`acetimego`. It is intended for resource-constrained microcontroller
environments where string identifiers can be wasteful and more difficult to
store, retrieve, and transmit.

The `TimeZone` object is created from the `ZoneManager` using either of these
identifiers:

```go
func doSomething() {
  zm := acetime.ZoneManagerFromDataContext(&zonedb2000.DataContext)
  tz1 := zm.TimeZoneFromName("America/Los_Angeles")
  if tz1.IsError() {
    // handle not found
  }

  tz2 := zm.TimeZoneFromZoneID(zonedb2000.ZoneIDAmerica_Los_Angeles)
  if tz2.IsError() {
    // handle not found
  }
  ...
}
```

We can query the `TimeZone` object for its name and id like this:

```go
  tz := zm.TimeZoneFromName("America/Los_Angeles")
  name := tz.Name() // returns "America/Los_Angeles")
  id := tz.ZoneID() // returns 0xb7f7e8f2
```

Some timezones are just symbolic links to another timezone in the TZDB. Most of
the time, the end-user does not need to know that, but it is available as the
`IsLink()` function:

```go
  tz := zm.TimeZoneFromName("US/Pacific")
  isLink := tz.IsLink() // returns true
```

(I just noticed that there is no function to retrieve the name of the target
timezone that the source is linked *to*. I think this can be added if needed.)

### TimeZone UTC

For convenience, the library automatically creates a special object for the UTC
timezone. This is the only `TimeZone` object which can be created without using
a `ZoneManager` and a specific `zonedb*` database:

```go
  utc := acetime.TimeZoneUTC
  isUTC := utc.IsUTC() // returns true
```

### ZonedDateTime

The `ZonedDateTime` is a pairing of the `PlainDateTime` and a `TimeZone` object.
There are 2 ways to create that binding:

- combine an explicit `PlainDateTime` object with a `TimeZone` object,
- convert an unixSeconds to `PlainDateTime` using the `TimeZone` object

Let's create a `ZonedDateTime` for the `America/Los_Angeles` time zone for the
date 2050-01-01T00:00:01:

```go
import (
  "github.com/bxparks/acetimego/acetime"
  "github.com/bxparks/acetimego/zonedb2000"
)

func doSomething() {
  zm := acetime.ZoneManagerFromDataContext(&zonedb2000.DataContext)
  tz := zm.TimeZoneFromName("America/Los_Angeles")
  pdt := acetime.PlainDateTime{2050, 1, 1, 0, 0, 1}
  zdt := acetime.ZonedDateTimeFromPlainDateTime(pdt, tz, DisambiguateCompatible)
  ...
}
```

Let's find the `ZonedDateTime` object that corresponds to the unixSeconds of
2524636801:

```go
func doSomething() {
  zm := acetime.ZoneManagerFromDataContext(&zonedb2000.DataContext)
  tz := zm.TimeZoneFromName("America/Los_Angeles")
  pdt := acetime.PlainDateTime{2050, 1, 1, 0, 0, 1}
  zdt := acetime.ZonedDateTimeFromUnixSeconds(pdt, tz, DisambiguateCompatible)
  ...
}
```

The `DisambiguateCompatible` option determines the behavior of the conversion
during a gap to daylight saving time (DST) or an overlap back to standard time
(STD). This is explained in the next section.

We can convert `ZonedDateTime` to an unixSeconds using the `UnixSeconds()`
function:

```go
  unixSeconds := zdt.UnixSeconds() // returns 2524636801
```

#### Disambiguate Gaps and Overlaps

During a DST change where the time goes back an hour (in the northern
hemisphere, during the "fall back" in Oct/Nov), a local time
appears twice for one hour. When we convert a `PlainDateTime` to a
`ZonedDateTime`, we need to be able to specify which of the 2 date-times to
select.

During a DST change where the time jumps forward an hour (in the northern
hemisphere, during the "spring forward" in Mar/Apr), there is a gap of one hour
in the local time. When we convert a `PlainDateTime` that falls in a gap to a
`ZonedDateTime`, we can either extend forward the UTC offset prior to the gap,
or extend backward the UTC offset after the gap.

The `disambiguate` parameter in the `ZonedDateTimeFromPlainDateTime()` function
determines the behavior of this function within a gap or overlap. The parameter
is *not* required for the `ZonedDateTimeFromUnixSeconds()` because the
conversion from unixSeconds to a `ZonedDateTime` can never produce a gap or
overlap.

The parameter is inspired by the `disambiguation` parameter in the
[Temporal](https://tc39.es/proposal-temporal/docs/zoneddatetime.html) JavaScript
library, and the `disambiguate` parameter in the
[Whenever](https://whenever.readthedocs.io/en/latest/overview.html#ambiguity-in-timezones)
Python library.

It accepts 4 values:

- `DisambiguateCompatible`: select the earlier time within an overlap, and the
  later time within a gap
- `DisambiguateEarlier`: always select the earlier time
- `DisambiguateLater`: always select the later time
- `DisambiguateReversed`: the opposite of `DisambiguateCompatible`

(The `acetimego` library does not support the `raise` options of the Temporal or
Whenever because Go does not support exceptions. Instead this library adds the
`DisambiguateReversed` option so that all 4 possible combinations are
implemented.)

#### Resolved Gaps and Overlaps

When we call `ZonedDateTimeFromUnixSeconds()`, the unixSeconds always
maps to a unique `ZonedDateTime`. There is no ambiguity and the `disambiguate`
parameter does not exist on the method.

The `ZonedDateTimeFromPlainDateTime()` accepts the `disambiguate` parameter
because the `PlainDateTime` can fall either in a gap or an overlap. The
resulting `ZonedDateTime` is normalized and validated, but sometimes we want to
know how the ambiguity was resolved.

The `ZonedDateTime.Resolved` parameter provides that information. It has 5
values:

- `ResolvedUnique`: always set by `ZonedDateTimeFromUnixSeconds()`, and set by
  `ZonedDateTimeFromPlainDateTime()` if the provided `PlainDateTime` maps to a
  unique time datetime
- `ResolvedOverlapEarlier`: the `PlainDateTime` was in an overlap and resolved
  to the earlier time
- `ResolvedOverlapLater`: the `PlainDateTime` was in an overlap and resolved to
  the later time
- `ResolvedGapEarlier`: the `PlainDateTime` was in a gap and resolved to the
  earlier time
- `ResolvedGapLater`: the `PlainDateTime` was in a gap and resolved to the later
  time

As noted above, the `ZonedDateTimeFromUnixSeconds()` function always maps to a
unique time, and the `Resolved` parameter will always be `ResolvedUnique`.

#### Convert TimeZone

We can convert the `ZonedDateTime` into another timezone:

```go
  tzParis := zm.TimeZoneFromName("Europe/Paris")
  zdtParis := zdt.ConvertToTimeZone(tzParis)
```

Just like `PlainDateTime`, we can convert a `ZonedDateTime` into a
human-readable string in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
format using the `String(` function:

```go
  s := zdt.String() // returns "2050-01-01T00:00:01-08:00[America/Los_Angeles]"
```

A `BuildString()` function is provided on the `ZonedDateTime` object as well.

#### ZonedDateTime Normalization

To simplify the API of the library and reduce the compiled-size of the library,
the `ZonedDateTime` object is *mutable*. You can overwrite a specific component
of the `ZonedDateTime` object, but you must remember to call the `Normalize()`
function after the change:

```go
  zm := acetime.ZoneManagerFromDataContext(&zonedb2000.DataContext)
  tz := zm.TimeZoneFromName("America/Los_Angeles")
  pdt := acetime.PlainDateTime{2050, 1, 1, 0, 0, 1}
  zdt := acetime.ZonedDateTimeFromPlainDateTime(pdt, tz, DisambiguateCompatible)

  // Change the month to July, and see incorrect date due to DST
  zdt.Month = 7
  s := zdt.String() // returns "2050-07-01T00:00:01-08:00[America/Los_Angeles]"

  // Must normalize.
  zdt.Normalize(DisambiguateCompatible)
  s = zdt.String() // returns "2050-07-01T00:00:01-07:00[America/Los_Angeles]"
```

The `Normalize()` function accepts the same `disambiguate` parameter as the
`ZonedDateTimeFromPlainDateTime()` function. Internally, it is essentially doing
the same thing as `ZonedDateTimeFromPlainDateTime()`, it is converting the
internal version of `PlainDateTime` inside the `ZonedDateTime`, then converting
it back to a normalized `ZonedDateTime` using the known `TimeZone` object.
During this normalization, the same problems with gaps and overlaps may occur,
which must be resolved using the `disambiguate` policy.

### ZonedExtra

The `ZonedExtra` object contains additional information that could have been
included in the `ZonedDateTime` but was extracted to a separate object because
they are not as commonly used. This allows the `ZonedDateTime` object to be
smaller.

The `ZonedExtra` has the following fields:

```go
type ZonedExtra struct {
  ResolvedFold        FoldType  // type of fold (e.g. gap, overlap)
  StdOffsetSeconds    int32     // STD offset
  DstOffsetSeconds    int32     // DST offset
  ReqStdOffsetSeconds int32     // request STD offset
  ReqDstOffsetSeconds int32     // request DST offset
  Abbrev              string    // abbreviation (e.g. PST, PDT)
}
```

It is created by:

- `acetime.ZonedExtraFromUnixSeconds(unixSeconds, tz)`
- `acetime.ZonedExtraFromPlainDateTime(plainDateTime, tz, disambiguate)`
- `acetime.ZonedDateTime.ZonedExtra()`

The `ResolvedFold` specifies whether the given `PlainDateTime` is within an
overlap or a gap. It takes 5 values:

- `FoldTypeNotFound`: an internal error occurred or the `PlainDateTime` was
  outside a valid range (`ZonedExtra.IsError()` returns `true`)
- `FoldTypeExact`: the `PlainDateTime` corresponds to a unique date-time value
- `FoldTypeGap`: the `PlainDateTime` falls in a gap
- `FoldTypeOverlap`: the `PlainDateTime` falls in an overlap.

The `Abbrev` parameter is the timezone abbreviation that corresponds to the
given unixSeconds or `PlainDateTime`. For example, for the `America/Los_Angeles`
time zone, it will be "PST" (Pacific Standard Time) during the winter months,
and "PDT" (Pacific Daylight Time) during the summer months.

For convenience, `ZonedDateTime` can directly retrieve the corresponding
`ZonedExtra` object using `ZonedDateTime.ZonedExtra()`.

(TODO: add documentation of the various OffsetSeconds parameters.)

## Bugs And Limitations

- `acetimgo` does not support access to a monotonic clock of the underlying
  system. The sole purpose of acetimego is to support timezones and date-times
  in the Gregorian calendar system.
- `acetimego` does not support the `time.Duration` object. The difference
  between two `acetime.Time` values can be represented as an `int64`.
- `acetimego` does not support date arithmetics such as adding days or months.
- `acetimego` does not support generalized formatting of the `ZonedDateTime`
  object similar to `time.Time.Format()`. Only one specific ISO 8601 format is
  supported by the `String()` or `BuildString()` functions.
- The internal algorithms have been tested primarily from the year 0001 to 9999.
  There may be bugs outside of that range.

## License

[MIT License](https://opensource.org/licenses/MIT)

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

## Authors

- Created by Brian T. Park (brian@xparks.net).
