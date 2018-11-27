package constant

type Transaction struct {
	from  Address
	to    Address
	value float64
	vin   map[Serial]float64
	vout  map[Serial]float64

	// 0: normal tx
	// 1: add a producer
	// 2: remove a producer
	// 3: new leader
	// 4: new cross-shard update from another shard
	txType int

	data interface{}
}

// only needed when txType is 4
type CrossShardData struct {

	// new utxo for target shard
	utxo MerkleTree

	// info on committee of source shard
	committee Committee

	// new account state for target shard
	// map[account]([variable]value)
	state map[Address](map[string]interface{})

	// for validation
	prevExternalBlock Hash
	sigs              []Signature
}
