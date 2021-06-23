package lib

type StakeEntryStats struct {
}

const HashSizeBytes = 32

type PKID [33]byte
type BlockHash [HashSizeBytes]byte

type StakeEntry struct {
}

func (b *BlockHash) Bytes() []byte {
	items := []byte{}
	for _, val := range b {
		items = append(items, val)
	}
	return items
}

type PostEntry struct {
	PostHash                 *BlockHash
	PosterPublicKey          []byte
	ParentStakeID            []byte
	Body                     []byte
	RecloutedPostHash        *BlockHash
	IsQuotedReclout          bool
	CreatorBasisPoints       uint64
	StakeMultipleBasisPoints uint64
	ConfirmationBlockHeight  uint32
	TimestampNanos           uint64
	IsHidden                 bool
	StakeEntry               *StakeEntry
	LikeCount                uint64
	RecloutCount             uint64
	QuoteRecloutCount        uint64
	DiamondCount             uint64
	stakeStats               *StakeEntryStats
	isDeleted                bool
	CommentCount             uint64
	IsPinned                 bool
	PostExtraData            map[string][]byte
}

type ProfileEntry struct {
	PublicKey   []byte
	Username    []byte
	Description []byte
	ProfilePic  []byte
	IsHidden    bool
	CoinEntry
	isDeleted                bool
	StakeMultipleBasisPoints uint64
	StakeEntry               *StakeEntry
	stakeStats               *StakeEntryStats
}

type DiamondEntry struct {
	SenderPKID      *PKID
	ReceiverPKID    *PKID
	DiamondPostHash *BlockHash
	DiamondLevel    int64
	isDeleted       bool
}

type LikeEntry struct {
	LikerPubKey   []byte
	LikedPostHash []byte
	isDeleted     bool
}

type RecloutEntry struct {
	ReclouterPubKey   []byte
	RecloutPostHash   *BlockHash
	RecloutedPostHash *BlockHash
	isDeleted         bool
}
