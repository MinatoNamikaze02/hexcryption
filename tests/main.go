package main

import (
	"fmt"
	"github.com/MinatoNamikaze02/hexcryption/hc"
)

func main(){
	testStr := [10]string{
		`MZbXypE5oS`,
		`YXTeIPiFaT`,
		`0jWqe6Rt3c`,
		`kRsIRFcEau`,
		`RdzpqjTgEg`,
		`yI3UWMzPIT`,
		`sfURJ71YWr`,
		`gShxPwoBlR`,
		`AuSNIWlc0x`,
		`5rZUPmNKtC`}

	for _, s := range testStr{
		fmt.Println("      ----------      ")
		fmt.Printf("\nInitial Value: %s", s)
		enc, key := se.Encoder(s)
		fmt.Println("\nEncrypted: ", enc)
		fmt.Println("\nKey: ", key)
		dec := se.Decoder(enc, key)
		fmt.Println("\nDecrypted: ", dec)
	}
	
}