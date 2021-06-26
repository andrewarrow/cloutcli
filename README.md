# cloutcli
library to make building things with bitclout easy

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
  clout sqlite query          # --term=foo
  clout sqlite likes          # --username=foo
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

