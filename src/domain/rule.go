package domain

type RuleStrings struct {
	Key   string   `json:",omitempty"`
	Value []string `json:",omitempty"`
}

type RuleInt struct {
	Key   string `json:",omitempty"`
	Value int    `json:",omitempty"`
}
