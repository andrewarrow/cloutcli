# cloutcli
library to make building things with bitclout easy

# building the "clout" executable
There is no main.go file in the root directory.

Instead go into "cmd" directory and `go build` there.

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

