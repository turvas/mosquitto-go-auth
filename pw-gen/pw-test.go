package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/iegomez/mosquitto-go-auth/hashing"
)

func main() {

	var password = flag.String("p", "", "plain text password")
	var hashed = flag.String("h", "", "pbkdf2 password hash. use single quotes so $ symbols aren't taken as variables")

	flag.Parse()
 
  // Test UTF8.
	hasher = NewPBKDF2Hasher(defaultPBKDF2SaltSize, defaultPBKDF2Iterations, defaultPBKDF2Algorithm, UTF8, defaultPBKDF2KeyLen)
     
	if hasher.Compare(password, hashed) {
		fmt.Println("success utf8: plain and hashed passwords match")
		return
	}
	fmt.Println("error: plain and hashed passwords don't match")
	os.Exit(1)
}
