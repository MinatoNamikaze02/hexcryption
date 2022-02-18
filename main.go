package main

import (
	"bufio"
	"fmt"
	"github.com/MinatoNamikaze02/hexcryption/hc"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)
var reader = bufio.NewReader(os.Stdin)

func opt1(){
	var inp string
	var enc string
	var key string
	println(`Enter the input text to be encoded: `)
	scanner.Scan()
	inp = scanner.Text()
	inp = strings.TrimSuffix(inp, "\n")
	enc, key = se.Encoder(inp)	
	fmt.Println("Encrypted: ", enc)
	fmt.Println("Key: ", key)
}

func opt2(){
	var inp string
	var dec string
	var key string
	println(`Enter the input text to be decoded: `)
	scanner.Scan()
	inp = scanner.Text()
	inp = strings.TrimSuffix(inp, "\n")
	println(`Enter the key: `)
	scanner.Scan()
	key = scanner.Text()
	key = strings.TrimSuffix(key, "\n")
	dec = se.Decoder(inp, key)	
	fmt.Println("Decrypted: ", dec)
}

func main(){
	if len(os.Args) < 2 {
		inputF()
		os.Exit(2)
	}
	cmd := os.Args[1]
	switch cmd {
	case "-e":
		var msgFlag string
		if len(os.Args) > 2{
			inp := os.Args[2]
			fmt.Println("\nInput String:",inp)
			enc, key := se.Encoder(string(inp))
			msgFlag = fmt.Sprint("\nEncrypted: ", enc, "\n\n", "Key: ", key)
		}
		fmt.Println(msgFlag)
	case "-d":
		var msgFlag string
		if len(os.Args) > 2{
			inp := os.Args[2]
			if strings.Compare(os.Args[3], "-d") == 1 && len(os.Args) > 4{
				key := os.Args[4]
				dec := se.Decoder(inp, key)
				msgFlag = fmt.Sprint("\nDecrypted: ", dec)
				fmt.Println(msgFlag)
			}else{
				fmt.Println("No key provided! Cannot decode without the key")
				fmt.Println("try running with --help flag to know more about the command usage!")
				os.Exit(2)
			}
		}
	case "-k":
		fmt.Println("Command to be used after '-d'")
		fmt.Println("try running with --help flag to know more about the command usage!\n")
	case "--help":
		fmt.Printf("Usage: <command> [<args>]")
		commands := map[string] string {
			"-e" : "Followed by the input text to be encoded",
			"-d" : "Followed by the input text to be decoded",
			"-k" : "Followed by the key used to decode(to be entered after -d)",
			"--help": "Prints this message",
		}
		for k, v :=  range commands{
			fmt.Printf(fmt.Sprint("\n",k,"\t",v))
		}
	default:
		fmt.Printf("Unknown Command: %s\n", cmd)
	}
}

func inputF(){
	fmt.Println("Enter your choice (e, E for encrypting a text and D or d for decrypting)")
	ch, _, _ := reader.ReadRune()
	switch ch{
	case 'e', 'E':
		opt1()
		break
	case 'd', 'D':
		opt2()
		break
	default:
		fmt.Println("Invalid choice!")
		os.Exit(2)
	}

}