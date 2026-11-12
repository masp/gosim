//go:build cgo

package cache

import _ "github.com/mattn/go-sqlite3"

// sqliteDriver is the database/sql driver used for the cache. With cgo the
// cache uses github.com/mattn/go-sqlite3.
const sqliteDriver = "sqlite3"
