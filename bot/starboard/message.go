package starboard

type StarredMessage struct {
	ChannelID string
	OriginalMessageID string
	StarboardMessageID string
}

func (m *StarredMessage) GetStarCount() int {
	// TODO: fetch reactions from original message
	// TODO: fetch reactions from starboard message
	// TODO: exclude bot reactions
	// TODO: exclude duplicates via different emojis from the same user
	// TODO: exclude duplicates via the same user on both messages
	// TODO: return total star count
}