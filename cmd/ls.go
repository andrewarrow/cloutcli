package main

import (
	"fmt"

	"github.com/andrewarrow/cloutcli"
)

func HandleLs() {
	list := cloutcli.GlobalPosts()

	for _, post := range list {
		fmt.Println(post.Body)
	}
}
