package argon2id

type Argon2Handlers interface {
	HashPassword()
}

type Argon2Configs struct {
	HashRaw    []byte
	Salt       []byte
	TimeCost   uint32
	MemoryCost uint32
	Threads    uint8
	KeyLength  uint32
}

// Example:$argon2id$v=19$m=65536,t=3,p=4$G8NYSxrA+UMGHJbZVIXXXQ$UrHyBcYfCEms+92QVzGmfYqrWtH54WJY9FuROBQi/X8
// Format Components:

// argon2id: Algorithm variant identifier

// v=19: Argon2 version

// m=65536,t=3,p=4: Parameter encoding (memory, time, parallelism)

// G8NYSxrA+UMGHJbZVIXXXQ: Base64-encoded salt

// UrHyBcYfCEms+92QVzGmfYqrWtH54WJY9FuROBQi/X8: Base64-encoded hash

func NewArgonConfig() *Argon2Configs {
	return &Argon2Configs{
		
	}
}