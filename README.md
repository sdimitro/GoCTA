# GoCTA

[![Build Status](https://travis-ci.org/sdimitro/GoCTA.svg?branch=master)](https://travis-ci.org/sdimitro/GoCTA)

A Golang wrapper for the Chicago Transit Authority Train API.

### Motivation

I want to be able to track the arrival times of the train near
my apartment from the command line, my i3 status bar, or through
Acme. I haven't really decided yet but they all need data in
plain text. The reason why I chose Go instead of some scripting
language was that I just wanted to get more familiar with the
language and its tooling.

### What's implemented?

* Submitting requests to the API
* Parsing the XML responses generated from the API.
* A sample command that calls the API and prints
  the response to standard output.
* A sample command that uses that parses XML responses
  coming from standard input.
* Rudimentary testing and examples

### What's on the TODO list?

* Fill the rest of StationMapID map
* Have a local copy of the API and map IDs
* Import station IDs and map the to station names

![Image of Train](img/train.jpg)
![Image of Gopher](img/gopher.png)

