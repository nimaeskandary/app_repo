# Wails Guide

You can expect a wails app `cmd/` folder to have its own README, but this guide is more general info that applies to all wails projects:

## Ios logs

`./bin/wails-gordle task ios:logs:dev`

or

```
xcrun simctl launch \
--terminate-running-process \
--console-pty \
booted \
com.nimaeskandary.gordle.dev
```
