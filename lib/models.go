package lib

type PostsStateless struct {
	PostsFound []Post
}

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
	CreatorBasisPoints      uint64
	BitCloutLockedNanos     uint64
	NumberOfHolders         uint64
	CoinsInCirculationNanos uint64
	CoinWatermarkNanos      uint64
}

type SingleProfile struct {
	Profile ProfileEntryResponse
}

type TxReady struct {
	TstampNanos                      int64
	TransactionHex                   string
	ExpectedBitCloutReturnedNanos    int64
	ExpectedCreatorCoinReturnedNanos int64
	SpendAmountNanos                 int64
	TotalInputNanos                  int64
	ChangeAmountNanos                int64
	FeeNanos                         int64
}

type MessageList struct {
	NumberOfUnreadThreads       int64
	OrderedContactsWithMessages []MessageThing
	PublicKeyToProfileEntry     map[string]ProfileEntryResponse
	UnreadStateByContact        map[string]bool
}

type MessageThing struct {
	PublicKeyBase58Check string
	Messages             []Message
}

type Message struct {
	SenderPublicKeyBase58Check    string
	RecipientPublicKeyBase58Check string
	EncryptedText                 string
	TstampNanos                   int64
	IsSender                      bool
}

type PublicKeyToProfileEntry struct {
	PublicKeyToProfileEntry map[string]ProfileEntryResponse
	NumFollowers            int64
}

type UsersStateless struct {
	UserList []User
}

type User struct {
	PublicKeyBase58Check string
	ProfileEntryResponse ProfileEntryResponse
	BalanceNanos         int64
	UsersYouHODL         []HODLerThing
}
type HODLerThing struct {
	HODLerPublicKeyBase58Check  string
	CreatorPublicKeyBase58Check string
	HasPurchased                bool
	BalanceNanos                int64
	ProfileEntryResponse        ProfileEntryResponse
}

type Rate struct {
	SatoshisPerBitCloutExchangeRate int64
	NanosSold                       int64
	USDCentsPerBitcoinExchangeRate  int64
}

type HodlersWrap struct {
	Hodlers []Hodler
}

type Hodler struct {
	HODLerPublicKeyBase58Check string
	BalanceNanos               int64
	ProfileEntryResponse       ProfileEntryResponse
}
