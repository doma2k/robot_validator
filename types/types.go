package types

type Proposal struct {
	ProposalID string  `json:"proposal_id"`
	Status     string  `json:"status"`
	Content    Content `json:"content"`
}

type Content struct {
	Type        string `json:"@type"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Proposals struct {
	Proposals []Proposal `json:"proposals"`
}

type Release struct {
    Tag_name    string `json:"tag_name"`
}
