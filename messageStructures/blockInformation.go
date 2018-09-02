package messageStructures

type BlockInformation struct {
	BlockNumber     uint32
	BlockHash       [32]byte
	BlockMerkleRoot [32]byte
}
