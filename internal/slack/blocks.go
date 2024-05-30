package slack

type TextBlock struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type SectionBlock struct {
	Type string    `json:"type"`
	Text TextBlock `json:"text"`
}

type Blocks struct {
	Blocks []SectionBlock `json:"blocks"`
}

func NewBlocksMap(blocks []SectionBlock) (Blocks, error) {
	return Blocks{Blocks: blocks}, nil
}
