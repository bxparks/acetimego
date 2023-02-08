# Changelog

* Unreleased
    * Use one-second resolution for STDOFF and DSTOFF instead of one-minute.
        * Increases zonedb by about ~1kB on ESP32.
        * Increases RAM usage of ZoneProcessor by ~200 bytes on ESP32.
        * Decreases flash size of `acetime` package on ESP32.
    * Change `ATime` type from `int32` to `int64`.
        * Very little change in flash size of `acetime`.
* 0.1.0 (2023-01-29, TZDB 2022g)
    * First internal release.
