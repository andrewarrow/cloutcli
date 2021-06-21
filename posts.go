package cloutcli

func GlobalPosts() []Post {
	p := Post{}
	p.Body = "test"
	return []Post{p}
}
