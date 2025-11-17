# Changelog

- unreleased
- 0.9.0 (2025-11-17, TZDB 2025b)
    - **Breaking** Replace  `ZonedExtra.FoldType` with `ZonedExtra.Resolved`
      which has the same meaning as `ZonedDateTime.Resolved`.
    - **Breaking** Use typesafe types for enums (`findResultType` instead of
      `uint8`, `ResolvedType` instead of `uint8`).
    - **Breaking** Remove `ZonedDataTime.ZonedExtra()` that converts
      `ZonedDateTime` into a `ZonedExtra`.
        - Use `ZonedExtraFromPlainDateTime()` or `ZonedExtraFromUnixSeconds()`
          instead.
        - For compatibility with `AceTime` and `acetimec` libraries. Simplifies
          the API.
    - Regenerate memory benchmarks using Go 1.25.3 and TinyGo 0.39.0.
- 0.8.0 (2025-10-21, TZDB 2025b)
    - **Breaking** Rename LocalXxx to PlainXxx, following the conventions
      used by more modern timezone libraries (JavaScript Temporal, Python
      whenever).
        - LocalDate -> PlainDate
        - LocalDateTime -> PlainDateTime
        - LocalTime -> PlainTime
    - **Breaking** Zone database refactoring.
        - Rename `zonedb` package to `zonedb2000`.
        - Add new `zonedb2025` package which contains timezone data from the
          year 2025 and onwards. Reduces size of the database from 35 kB to
          about 26 kB.
    - **Breaking** Rename `EpochSeconds()` to `UnixSeconds()`
        - Also rename `EpochDays()` to `UnixDays().
        - The acetimego library uses a 64-bit integer for its acetime.Time type.
          It does not need to provide a user-configurable EpochDate. Therefore,
          we can always use the Unix epoch of 1970-01-01.
        - Using `UnixSeconds()` makes this convention self-documenting, and
          avoid confusion over which Epoch date is being used.
- 0.7.0 (2025-04-25, TZDB 2025b)
    - merge various AceTime projects into single AceTimeSuite repo
    - [upgrade to TZDB 2025a](https://lists.iana.org/hyperkitty/list/tz-announce@iana.org/thread/MWII7R3HMCEDNUCIYQKSSTYYR7UWK4OQ/)
        - Paraguay adopts permanent -03 starting spring 2024.
        - Improve pre-1991 data for the Philippines.
        - Etc/Unknown is now reserved.
    - [upgrade to TZDB 2025b](https://lists.iana.org/hyperkitty/list/tz-announce@iana.org/thread/6JVHNHLB6I2WAYTQ75L6KEPEQHFXAJK3/)
        - New zone for Aysén Region in Chile which moves from -04/-03 to -03.
          (Creates new zone named America/Coyhaique)
- 0.6.0 (2024-12-14, TZDB 2024b)
    - Support new `%z` value in FORMAT column.
    - Upgrade TZDB to 2024b
        - https://lists.iana.org/hyperkitty/list/tz-announce@iana.org/thread/IZ7AO6WRE3W3TWBL5IR6PMQUL433BQIE/
        - "Improve historical data for Mexico, Mongolia, and Portugal. System V
          names are now obsolescent. The main data form now uses %z. The code
          now conforms to RFC 8536 for early timestamps. Support POSIX.1-2024,
          which removes asctime_r and ctime_r. Assume POSIX.2-1992 or later for
          shell scripts. SUPPORT_C89 now defaults to 1."
- 0.5.2 (2024-07-26, TZDB 2024a)
    - Upgrade TZDB to 2024a
        - https://mm.icann.org/pipermail/tz-announce/2024-February/000081.html
        - "Kazakhstan unifies on UTC+5 beginning 2024-03-01. Palestine springs
          forward a week later after Ramadan. zic no longer pretends to support
          indefinite-past DST. localtime no longer mishandles Ciudad Juárez in
          2422."
- 0.5.1 (2024-01-12, TZDB 2023d)
    - zonedb
        - Simplify `tzcompiler.py` flags in various Makefiles.
    - Upgrade TZDB to 2023d
        - https://mm.icann.org/pipermail/tz-announce/2023-December/000080.html
        - "Ittoqqortoormiit, Greenland changes time zones on 2024-03-31. Vostok,
          Antarctica changed time zones on 2023-12-18. Casey, Antarctica changed
          time zones five times since 2020. Code and data fixes for Palestine
          timestamps starting in 2072. A new data file zonenow.tab for
          timestamps starting now."
- 0.5.0 (2023-05-31, TZDB 2023c)
    - Add `examples/helloacetime` demo.
    - Rename `ATime` to `Time`.
        - Causes no conflict with `time.Time` because the package prefix
          `acetime` is always required.
    - `ZonedExtra`
        - Rename `ZonedExtra.Zetype` to `FoldType`.
        - Rename `ZonedExtraXxx` to `FoldTypeXxx` for self-documentation.
        - Add `OffsetSeconds()` convenience method which calculates
          `StdOffsetSeconds + DstOffsetSeconds`.
    - `OffsetDateTime`, `ZonedDateTime`
        - Use struct embedding feature of Go lang to simplify and reuse code.
        - `PlainDateTime` embeds into `OffsetDateTime`.
        - `OffsetDateTime` embeds into `ZonedDateTime`.
    - Added to `AceTimeValidation`
        - Validated with AceTimeC and C++ Hinnant libraries from [1800,2100).
- 0.4.0 (2023-05-21, TZDB 2023c)
    - Rename project from `AceTimeGo` to `acetimego`
        - More consistent with Go library naming convention.
    - zonedb
        - Upgrade to TZDB 2023c
        - Rename `zonedb` to `zonedball`. Consumes 72050 bytes of flash.
        - Repurpose `zonedb` to support only year 2000 and onwards.
          Consumes 34948 bytes of flash, less than 1/2 of `zonedball`.
    - Replace out-of-band `isFilled` with in-band `year!=InvalidYear`.
    - Support DS3231 RTC chip under TinyGo
        - Add `ds3231` package
        - Add `examples/ds3231demo`
- 0.3.0 (2023-03-10, TZDB 2022g)
    - Convert C++ doxygen docs to go doc format.
    - Unexport various internal functions and structs by lowercasing their
      names. Could not use `internal/` directory due to circular dependency to
      `ATime` type.
    - Move `strutil.go` utilities to new `strbuild` package and simplify
      names of various functions.
        - These low-level functions are used by some of the utilities in `cmd`.
        - Avoids cluttering the API of the `acetime` package.
    - Rename `PlainDateToDayOfWeek()` to `PlainDateToWeekday()`
        - For better consistency with standard `time` package.
        - Add IsoWeekday.Name() method.
    - Add `PlainDateToYearday()`
        - Returns the day of year of `(year, month, day)`.
    - Add `ZonedDateTime.ZonedExtra()` convenience method.
        - For easier access to extra information about the given date time.
    - zonedb
        - Always generate anchor rules in zonedb.
        - Allows `zone_processor.go` to work over all years `[0,10000)`
          even with truncated zonedb (e.g. `[2000,2100)`).
        - Accuracy is guaranteed only for the requested interval (e.g.
          `[2000,2100)`. But the code won't crash outside of that interval.
        - Change `START_YEAR` to 1800, incorporating the complete TZDB.
        - Use simplified ZoneRule filtering.
        - Extract `XxxRecords` objects into separate `zone*_test.go` files.
        - Autogenerate `reader_test.go`.
- 0.2.0 (2023-02-13, TZDB 2022g)
    - Support one-second resolution for Zone.STDOFF field, instead of
      one-minute.
        - Increases zonedb by about ~1kB on ESP32.
        - Increases RAM usage of ZoneProcessor by ~200 bytes on ESP32.
        - Decreases flash size of `acetime` package on ESP32.
    - Change `ATime` type from `int32` to `int64`.
        - Very little change in flash size of `acetime`.
    - Support one-second resolution for Zone.UNTIL and Rule.AT fields.
        - Allows the library to support all zones before ~1972.
    - Support one-minute resolution for Zone.DSTOFF (i.e. Zone.RULES) and
      Rule.SAVE fields.
        - Handles a few zones around ~1930 whose DSTOFF is a 00:20 minutes,
          instead of a multiple of 00:15 minutes.
    - Extend `zonedb` year interval.
        - Raw TZDB year range is `[1844,2087]`.
        - Regenerate the zonedb for `[3,10000)`.
        - Increases zonedb flash size for ESP32 from 44kB to 72kB.
    - `cmd/validatetime/`
        - AceTimeGo and the standard `time` library match perfectly,
          for all zones, from `[1800,2100)`.
- 0.1.0 (2023-01-29, TZDB 2022g)
    - First internal release.
