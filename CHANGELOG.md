# Changelog

* Unreleased
    * Convert C++ doxygen docs to go doc format.
    * Unexport various internal functions and structs by lowercasing their
      names. Could not use `internal/` directory due to circular dependency to
      `ATime` type.
    * Move `strutil.go` utilities to new `strbuild` package and simplify
      names of various functions.
        * These low-level functions are used by some of the utilities in `cmd`.
        * Avoids cluttering the API of the `acetime` package.
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
