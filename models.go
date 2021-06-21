package cloutcli

type Post struct {
	PostHashHex                string
	PosterPublicKeyBase58Check string
	ParentStakeID              string
	Body                       string
	PostExtraData              PostExtraData
	ImageURLs                  []string
	TimestampNanos             int64
	ProfileEntryResponse       ProfileEntryResponse
	LikeCount                  int64
	Comments                   []Post
	RecloutedPostEntryResponse *Post
	CommentCount               int64
	RecloutCount               int64
}

type PostExtraData struct {
	EmbedVideoURL string
}

type ProfileEntryResponse struct {
	PublicKeyBase58Check   string
	Username               string
	Description            string
	CoinEntry              CoinEntry
	CoinPriceBitCloutNanos int64
}

type CoinEntry struct {
	CreatorBasisPoints      int64
	BitCloutLockedNanos     int64
	NumberOfHolders         int64
	CoinsInCirculationNanos int64
	CoinWatermarkNanos      int64
}
