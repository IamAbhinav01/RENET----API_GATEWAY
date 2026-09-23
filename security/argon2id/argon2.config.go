package argon2id

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"renet/config/env"

	"golang.org/x/crypto/argon2"
)

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

	
	timeCost := env.GetInt("TimeCost")
	memoryCost := env.GetInt("MemoryCost")
	threads := env.GetInt("Threads")
	keyLength := env.GetInt("KeyLength")

	return &Argon2Configs{
		TimeCost: uint32(timeCost),
		MemoryCost: uint32(memoryCost),
		Threads: uint8(threads),
		KeyLength: uint32(keyLength),
		
	}
}

func generateSalt(salt_size uint32) ([]byte,error){
	salt := make([]byte,salt_size)// creates  aa slice of empty byet[] of salt size
	_,err := rand.Read(salt)//randomly initialise the slice with values
	if err != nil{
		log.Println("Error occured while generating salt")
		return nil,err
	}
	return salt,nil
}

func(config *Argon2Configs) HashPassword(inputPassword string)(string,error){

	salt_size := env.GetInt("salt_size")
	salt,err := generateSalt(uint32(salt_size))

	if err!= nil{
		log.Println("Error while generating the salt")
		return "",err
	}

	config.Salt = salt
	config.HashRaw = argon2.IDKey([]byte(inputPassword),config.Salt,config.TimeCost,config.MemoryCost,config.Threads,config.KeyLength)


	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",argon2.Version,config.MemoryCost,config.TimeCost,config.Threads,base64.RawStdEncoding.EncodeToString(config.Salt),base64.RawStdEncoding.EncodeToString(config.HashRaw))

	return encodedHash,nil

}
