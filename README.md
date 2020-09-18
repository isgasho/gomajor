# GOMAJOR

> This is an experimental tool for upgrading major versions

### Example:

```
$ gomajor list
github.com/go-redis/redis: v6.15.9+incompatible [latest v8.1.3]
```

```
$ gomajor get github.com/go-redis/redis@latest
go get github.com/go-redis/redis/v8@v8.1.3
foo.go: github.com/go-redis/redis -> github.com/go-redis/redis/v8
bar.go: github.com/go-redis/redis -> github.com/go-redis/redis/v8
```

### Features:

* Finds latest version.
* Rewrites your import paths.
* Lets you ignore SIV on the command line.

### Warning:

* This tool has no dry-run feature. Commit before running.
* If you have multiple major versions imported, ALL of them will be rewritten.
* `@latest` scrapes pkg.go.dev and will stop working at some point.
* `@master` is not supported.
* gopkg.in imports are not supported.
* +incompatible versions are not supported.
