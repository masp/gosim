//go:build !cgo

package cache

import _ "modernc.org/sqlite"

// sqliteDriver is the database/sql driver used for the cache. Without cgo
// (common on Windows, where no C compiler is installed by default) the cache
// uses the pure Go modernc.org/sqlite, which reads and writes the same file
// format as github.com/mattn/go-sqlite3.
const sqliteDriver = "sqlite"
