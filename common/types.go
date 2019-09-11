package common

type Block struct {
	Number          string   `bson:"number" json:"number"`
	Timestamp       string   `bson:"timestamp" json:"timestamp"`
	Transactions    []string `bson:"transactions" json:"transactions"`
	Hash            string   `bson:"hash" json:"hash"`
	ParentHash      string   `bson:"parentHash" json:"parentHash"`
	Sha3Uncles      string   `bson:"sha3Uncles" json:"sha3Uncles"`
	Miner           string   `bson:"miner" json:"miner"`
	Difficulty      string   `bson:"difficulty" json:"difficulty"`
	TotalDifficulty string   `bson:"totalDifficulty" json:"totalDifficulty"`
	Size            string   `bson:"size" json:"size"`
	GasUsed         string   `bson:"gasUsed" json:"gasUsed"`
	GasLimit        string   `bson:"gasLimit" json:"gasLimit"`
	Nonce           string   `bson:"nonce" json:"nonce"`
	Uncles          []string `bson:"uncles" json:"uncles"`
	ExtraData       string   `bson:"extraData" json:"extraData"`
	LogsBloom       string   `bson:"logsBloom" json:"logsBloom"`
	MixHash         string   `bson:"mixHash" json:"mixHash"`
	ReceiptsRoot    string   `bson:"receiptsRoot" json:"receiptsRoot"`
	StateRoot       string   `bson:"stateRoot" json:"stateRoot"`
	TransactionRoot string   `bson:"transactionsRoot" json:"transactionsRoot"`
}
