module clout

go 1.16

replace github.com/andrewarrow/cloutcli => ../

require (
	github.com/andrewarrow/cloutcli v0.0.0-00010101000000-000000000000
	github.com/btcsuite/btcutil v1.0.2
	github.com/dgraph-io/badger/v3 v3.2103.0
	github.com/justincampbell/bigduration v0.0.0-20160531141349-e45bf03c0666 // indirect
	github.com/justincampbell/timeago v0.0.0-20160528003754-027f40306f1d
	go.mongodb.org/mongo-driver v1.5.3
)
