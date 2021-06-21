package main

import (
	"fmt"

	"github.com/andrewarrow/cloutcli"
)

func main() {
	fmt.Println("list messages from the global feed...")

	list := cloutcli.GlobalPosts()

	for _, post := range list {
		fmt.Println(post.Body)
	}

}
