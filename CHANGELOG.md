# Changelog

* Unreleased
    * zonedb
        * Rename `zonedb` to `zonedball`. Consumes 72050 bytes of flash.
        * Create `zonedb2000` which supports only year 2000 and onwards.
          Consumes 34948 bytes of flash, less than 1/2 of `zonedball`.
* 0.3.0 (2023-03-10, TZDB 2022g)
    * Convert C++ doxygen docs to go doc format.
    * Unexport various internal functions and structs by lowercasing their
      names. Could not use `internal/` directory due to circular dependency to
      `ATime` type.
    * Move `strutil.go` utilities to new `strbuild` package and simplify
      names of various functions.
        * These low-level functions are used by some of the utilities in `cmd`.
        * Avoids cluttering the API of the `acetime` package.
    * Rename `LocalDateToDayOfWeek()` to `LocalDateToWeekday()`
        * For better consistency with standard `time` package.
        * Add IsoWeekday.Name() method.
    * Add `LocalDateToYearday()`
        * Returns the day of year of `(year, month, day)`.
    * Add `ZonedDateTime.ZonedExtra()` convenience method.
        * For easier access to extra information about the given date time.
    * zonedb
        * Always generate anchor rules in zonedb.
        * Allows `zone_processor.go` to work over all years `[0,10000)`
          even with truncated zonedb (e.g. `[2000,2100)`).
        * Accuracy is guaranteed only for the requested interval (e.g.
          `[2000,2100)`. But the code won't crash outside of that interval.
        * Change `START_YEAR` to 1800, incorporating the complete TZDB.
        * Use simplified ZoneRule filtering.
        * Extract `XxxRecords` objects into separate `zone*_test.go` files.
        * Autogenerate `reader_test.go`.
* 0.2.0 (2023-02-13, TZDB 2022g)
    * Support one-second resolution for Zone.STDOFF field, instead of
      one-minute.
        * Increases zonedb by about ~1kB on ESP32.
        * Increases RAM usage of ZoneProcessor by ~200 bytes on ESP32.
        * Decreases flash size of `acetime` package on ESP32.
    * Change `ATime` type from `int32` to `int64`.
        * Very little change in flash size of `acetime`.
    * Support one-second resolution for Zone.UNTIL and Rule.AT fields.
        * Allows the library to support all zones before ~1972.
    * Support one-minute resolution for Zone.DSTOFF (i.e. Zone.RULES) and
      Rule.SAVE fields.
        * Handles a few zones around ~1930 whose DSTOFF is a 00:20 minutes,
          instead of a multiple of 00:15 minutes.
    * Extend `zonedb` year interval.
        * Raw TZDB year range is `[1844,2087]`.
        * Regenerate the zonedb for `[3,10000)`.
        * Increases zonedb flash size for ESP32 from 44kB to 72kB.
    * `cmd/validatetime/`
        * AceTimeGo and the standard `time` library match perfectly,
          for all zones, from `[1800,2100)`.
* 0.1.0 (2023-01-29, TZDB 2022g)
    * First internal release.
