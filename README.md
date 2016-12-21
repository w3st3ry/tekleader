# TekLeader

[![Build Status](https://travis-ci.org/w3st3ry/tekleader.svg?branch=master)](https://travis-ci.org/w3st3ry/tekleader)

Tekleader is a POSIX compliant and multiplatform CLI allowing to
establish a rank between {EPITECH} students from any city and/or
promotion, and much more.

## Disclaimer

This is not an official EPITECH project. Just a student who wants to
add new cool stuff and features using intranet API.

## Features

* Dynamic student GPA ranking from any city/promotion
* Intranet status access
* Dynamic API wrapper in Golang

## Roadmap

* Improve promotion ranking with all cities
* Separate API wrapper and core project
* Improve benchmarks using more concurrency
* Define more preferences using config file
* Fix bocal.exe :noel:

# Getting started

## Installation

You have three ways to get it:

- Go get the app

```
go get github.com/w3st3ry/tekleader
```

- Build the `Dockerfile`

```
docker build -t tekleader .
```

- Download the latest release according your OS
  [here](https://github.com/w3st3ry/tekleader/releases).

## Configuration

You just have to add the configuration file `.tekleader.yml` in
`/etc/.` or `$HOME/.` with your personal authkey (which can be found
[here](https://intra.epitech.eu/admin/autolog)) written in the same
format as in the example (without `https://intra.epitech.eu/`).

You can also set `TEK_AUTHKEY` as a env variable or directly as binary
parameter.

Warning: If you use this app in a container, you must share the file
or expose env variable for authentication.

## Usage

### Commands

* `leader`: leader establish a rank between students.
* `status`: give intranet status in continue.
* `version`: prints the tekleader version and any available update.

### Options

#### Global options

* `auth-key`: your authentication key, available on the intranet.
* `timeout`: timeout (in sec) for check status requests (default: 2).

#### Leader options

* `location`: set your city (default: lyon).
* `promotion`: set your promotion (default: tek2).
* `race`: enable race condition to print users (default: false).
* `find`: find student by login or first name/last name (ex:
  valentin.pichard, solomon.hykes ...).

## Hacking

With pleasure, fork with us :rocket:
