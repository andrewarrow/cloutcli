module clout

go 1.16

replace github.com/andrewarrow/cloutcli => ../

require (
	github.com/andrewarrow/cloutcli v0.0.0-00010101000000-000000000000
	github.com/dgraph-io/badger v1.6.2 // indirect
	github.com/dgraph-io/badger/v3 v3.2103.0
	github.com/justincampbell/timeago v0.0.0-20160528003754-027f40306f1d
)
