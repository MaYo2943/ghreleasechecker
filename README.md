# gh_release_checker

ghReleaseChecker is a small CLI utility that uses the Github API to
periodically check a list of repositories for new releases.

[![license](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://raw.githubusercontent.com/McKael/ghreleasechecker/master/LICENSE)
[![Build Status](https://travis-ci.org/McKael/ghreleasechecker.svg?branch=master)](https://travis-ci.org/McKael/ghreleasechecker)
[![Go Report Card](https://goreportcard.com/badge/github.com/McKael/ghreleasechecker)](https://goreportcard.com/report/github.com/McKael/ghreleasechecker)

By default, it outputs the new versions to stdout, which might be suitable for
a cron job, but it is possible to use a simple (Go) template or a JSON/YAML
format that can be used for automation.

A YAML configuration file is required; you can find a sample in the repository
root directory.

ghReleaseChecker uses a JSON state file, which path should be defined in the
configuration file.


Here's a sample use case:
```
% ghreleasechecker --config ./ghreleasechecker.yaml -o plain
New release for kubernetes/kubernetes: v1.10.0
  Tag: v1.10.0
  Date: 2018-03-27 01:57:17 +0200 CEST
New release for BurntSushi/ripgrep: 0.8.1
  Tag: 0.8.1
  Date: 2018-02-21 03:11:44 +0100 CET
New release for restic/restic: restic 0.8.3
  Tag: v0.8.3
  Date: 2018-02-26 21:41:52 +0100 CET
```

The output can be customized with templates, here's an example with the
default configuration file:
```
% ghreleasechecker --config ghreleasechecker.yaml
McKael/ghreleasechecker ghReleaseChecker 0.0.1 -- Initial release v0.0.1 2018-04-02 00:06:03 +0200 CEST
kubernetes/kubernetes v1.10.0 v1.10.0 2018-03-27 01:57:17 +0200 CEST
BurntSushi/ripgrep 0.8.1 0.8.1 2018-02-21 03:11:44 +0100 CET
restic/restic restic 0.8.3 v0.8.3 2018-02-26 21:41:52 +0100 CET
```

Colors can be used on terminals supporting ANSI sequences.

![Screenshot](ghreleasechecker_template.png "Screenshot")

Please check the commented YAML sample configuration file provided with the
source code for the details, and the online help for CLI usage.
