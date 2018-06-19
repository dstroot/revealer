# Revealer

[![Build Status](https://travis-ci.org/dstroot/revealer.svg?branch=master)](https://travis-ci.org/dstroot/revealer)
[![Go Report Card](https://goreportcard.com/badge/github.com/dstroot/revealer)](https://goreportcard.com/report/github.com/dstroot/revealer)
[![GoDoc](https://godoc.org/github.com/dstroot/revealer?status.svg)](https://godoc.org/github.com/dstroot/revealer)

A [go](http://www.golang.org) (or 'golang' for search engine friendliness) tool for "de-obfuscating" email addresses.  Pass in an obfuscated email in string format and it will attempt to figure out the valid email address.  

**NOTE:** Requires Go 1.10 or above due to use of "strings.Builder".

## Examples

See [the project documentation](https://godoc.org/github.com/dstroot/revealer) for examples of usage.

## Project Status & Versioning

The API should be considered stable. Feedback and feature requests are appreciated.  

This project uses [Semantic Versioning 2.0.0](http://semver.org).  Accepted pull requests will land on `master`.  Periodically, versions will be tagged from `master`.  You can find all the releases on [the project releases page](https://github.com/dstroot/revealer/releases).

## More

Documentation can be found [on godoc.org](http://godoc.org/github.com/dstroot/revealer).

## TODO 
* [ ] Make sure we handle Unicode properly
* [ ] Support international addresses/punycode

### References
http://jasonpriem.com/obfuscation-decoder/
https://stackoverflow.com/questions/2049502/what-characters-are-allowed-in-an-email-address
https://en.wikipedia.org/wiki/Email_address#Local-part
https://social.technet.microsoft.com/Forums/ie/en-US/69f393aa-d555-4f8f-bb16-c636a129fc25/what-are-valid-and-invalid-email-address-characters?forum=exchangesvradminlegacy
