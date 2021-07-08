package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/andrewarrow/cloutcli"
)

func PrintPostHelp() {
	fmt.Println("")
	fmt.Println("  clout post inspect [payload]      # full url or hex hash")
	fmt.Println("  clout post csv [payload]          # full url or hex hash")
	fmt.Println("")
}
func HandlePosts() {
	if len(os.Args) < 3 {
		PrintPostHelp()
		return
	}
	command := os.Args[2]
	if command == "inspect" {
		payload := os.Args[3]
		HandleOnePost(payload)
	} else if command == "csv" {
		payload := os.Args[3]
		MakePostAndRepliesCSV(payload)
	}
}

func ExtractJustHash(payload string) string {
	hash := payload
	if strings.HasPrefix(payload, "https") {
		hash = payload[27:]
	}
	return hash
}

func HandleOnePost(payload string) {
	hash := ExtractJustHash(payload)
	p := cloutcli.SinglePost(hash)
	fmt.Printf("%+v\n", p)
}

func addQuotes(s []string) []string {
	items := []string{}
	for _, thing := range s {
		items = append(items, "\""+thing+"\"")
	}
	return items
}
func MakePostAndRepliesCSV(payload string) {
	hash := ExtractJustHash(payload)
	p := cloutcli.SinglePost(hash)
	if p.PostHashHex == "" {
		return
	}

	headers := []string{"PostHashHex", "PosterPublicKeyBase58Check",
		"ParentStakeID", "Body", "TimestampNanos", "ProfileEntryResponse.Username",
		"LikeCount", "DiamondCount", "PostExtraData", "CommentCount", "RecloutCount",
		"QuoteRecloutCount"}

	fmt.Println(strings.Join(addQuotes(headers), ","))

	row1 := []string{p.PostHashHex, p.PosterPublicKeyBase58Check,
		p.ParentStakeID, strings.Replace(p.Body, "\"", "'", -1),
		fmt.Sprintf("%d", p.TimestampNanos), p.ProfileEntryResponse.Username,
		fmt.Sprintf("%d", p.LikeCount),
		fmt.Sprintf("%d", p.DiamondCount),
		fmt.Sprintf("%v", p.PostExtraData),
		fmt.Sprintf("%d", p.CommentCount),
		fmt.Sprintf("%d", p.RecloutCount),
		fmt.Sprintf("%d", p.QuoteRecloutCount)}

	fmt.Println(strings.Join(addQuotes(row1), ","))

	for _, r := range p.Comments {
		row := []string{r.PostHashHex, r.PosterPublicKeyBase58Check,
			r.ParentStakeID, strings.Replace(r.Body, "\"", "'", -1),
			fmt.Sprintf("%d", r.TimestampNanos), r.ProfileEntryResponse.Username,
			fmt.Sprintf("%d", r.LikeCount),
			fmt.Sprintf("%d", r.DiamondCount),
			fmt.Sprintf("%v", r.PostExtraData),
			fmt.Sprintf("%d", r.CommentCount),
			fmt.Sprintf("%d", r.RecloutCount),
			fmt.Sprintf("%d", r.QuoteRecloutCount)}

		fmt.Println(strings.Join(addQuotes(row), ","))
	}
}
