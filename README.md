# cloutcli
library to make building things with bitclout easy

# building the "clout" executable
There is no main.go file in the root directory.

Instead cd to the "cmd" directory and run:

```
go mod download
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

