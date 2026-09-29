package argon2id

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"log"
	"renet/config/env"
	"strings"

	"golang.org/x/crypto/argon2"
)

type Argon2Handlers interface {
	HashPassword(inputPassword string)(string,error)
	VerifyPassword(hashedPassword string,userPassword string)(bool,error)
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

func parseArgon2dHash(hashedPassword string)(*Argon2Configs,error){
	components := strings.Split(hashedPassword, "$")

// 	components[0] = ""
// components[1] = "argon2id"
// components[2] = "v=19"
// components[3] = "m=65536,t=3,p=4"
// components[4] = "G8NYSxrA+UMGHJbZVIXXXQ"
// components[5] = "UrHyBcYfCEms+92QVzGmfYqrWtH54WJY9FuROBQi/X8"

	if len(components) != 6{
		fmt.Printf("Invalid Password Hash format")
		return nil,fmt.Errorf("invalid hash format structure")
	}

	if !strings.HasPrefix(components[1],"argon2id"){
		fmt.Printf("Unsupported argon format")
		return nil,fmt.Errorf("Unsupported argon format")
	}

	var version int
	fmt.Sscanf(components[2],"v=%d",&version)

	if version != argon2.Version{
		fmt.Printf("unsupported argon2 version")
		return nil,fmt.Errorf("unsupported argon2 version")
	}

	config :=&Argon2Configs{}

	_,err := fmt.Sscanf(components[3],"m=%d,t=%d,p=%d",&config.MemoryCost,&config.TimeCost,&config.Threads)

	if err != nil{
		fmt.Printf("invalid argon2 format")
		return nil,err
	}

	salt,err := base64.RawStdEncoding.DecodeString(components[4])

	if err != nil{
		fmt.Println("salt decoding salt")
		return nil,err
	}

	config.Salt = salt

	hash,hashErr := base64.RawStdEncoding.DecodeString(components[5])

	if hashErr!=nil{
		fmt.Println("hash decoding failed")
		return nil,hashErr
	}

	config.HashRaw = hash
	config.KeyLength = uint32(len(hash))

	return config,nil
}

func(config *Argon2Configs)VerifyPassword(hashedPassword string,userPassword string)(bool,error){
	// $argon2id$v=19$m=65536,t=3,p=4$G8NYSxrA+UMGHJbZVIXXXQ$UrHyBcYfCEms+92QVzGmfYqrWtH54WJY9FuROBQi/X8

	parsecfg,err := parseArgon2dHash(hashedPassword)

	if err != nil{
		fmt.Println("Error happend while parsing the password")
		return false,fmt.Errorf("Error happend while verigying the hash")
	}

	computedHash := argon2.IDKey([] byte(userPassword),parsecfg.Salt,parsecfg.TimeCost,parsecfg.MemoryCost,parsecfg.Threads,parsecfg.KeyLength)

	match := subtle.ConstantTimeCompare(computedHash,parsecfg.HashRaw) == 1

	return match,nil
}