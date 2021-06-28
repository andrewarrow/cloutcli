# cloutcli
library to make building things with bitclout easy

# quick start demo

```
cmd $ ./clout demo

  clout demo graph       # make clout.gv graph file
  clout demo posts       # print all clouts
  clout demo search      # search sqlite database
  clout demo sqlite      # place data into local sqlite database

  search examples:

  ./clout demo search --term=hi --table=users
  ./clout demo search --term=hi --table=posts
  ./clout demo search --term=username --table=follow --degrees=2
```

# full menu

```
cloutcli/cmd $ ./clout

  clout account               # list your various accounts
  clout ls                    # list global posts
  clout message               # send, send bulk, read
  clout mongo                 # query from mongodb
  clout sell                  # sell coins
  clout sqlite                # import from badger, query sqlite

cloutcli/cmd $ ./clout message

  clout message bulk           # --to=allfollowers [--text=foo]
  clout message inbox          # --filter=myhodlers
  clout message new            # --to=username [--text=foo]
  clout message reply          # --id=foo [--text=foo]
  clout message show           # --id=foo

cloutcli/cmd $ ./clout sell

  clout sell dust           # --limit=x [--execute]

cloutcli/cmd $ ./clout sqlite

  clout sqlite fill           # --dir=/path/to/badgerdb
  clout sqlite graph          # produce clout.gv file
  clout sqlite query          # --term=foo [--table=x]
```

# building the "clout" executable
There is no main.go file in the root directory.

```
cd cmd
go mod tidy
go build
./clout
```

This is done to keep the root directory having the package name "cloutcli".

Which allows other go programs to just import:

```
import "github.com/andrewarrow/cloutcli"
```

and then:

```
list := cloutcli.GlobalPosts()

for _, post := range list {
  fmt.Println(post.Body)
}
```

# Example
[github.com/andrewarrow/referential](https://github.com/andrewarrow/referential)

