package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"
)

//Block
type Block struct {
	//hash du block
	Hash []byte
	//données du block
	Data []byte
	//hash du block précédent
	PrevHash []byte
	//nonce
	Nonce int
}

//Blockchain avec un tableau contenant les pointeurs vers chaque block
type Blockchain struct {
	blocks []*Block
	diff   int
}

//création d'un hash
func (b *Block) HashingBlock() {
	//join prend un tableau 2d avec le hash du block et ses données puis les convertit
	blockData := bytes.Join([][]byte{b.Hash, b.Data, []byte(strconv.Itoa(b.Nonce))}, []byte{})
	// fmt.Println(blockData)
	blockHash := sha256.Sum256(blockData)
	// fmt.Println(blockHash)
	b.Hash = blockHash[:]
}

func (b *Block) Mining(difficulty int) {
	var toHash string
	for i := 0; i < difficulty; i++ {
		toHash += string(b.Hash[i])
	}
	v, err := strconv.Atoi(toHash)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	for v != 0 {
		b.Nonce += 1
		b.HashingBlock()

	}
	fmt.Println("Block mined !")

}

//création d'un block
//données du block, hash du block précédent, retourne le pointeur du block créé
func CreateBlock(data string, previousHash []byte) *Block {
	//créé un block avec sa référence
	//conversion des données en bytes
	block := &Block{[]byte{}, []byte(data), previousHash, 0}
	//hash le block
	block.Mining(2)
	block.HashingBlock()
	return block
}

//ajout d'un block
func (chain *Blockchain) AddBlock(data string) {
	//récupération du block précédent
	previousBlock := chain.blocks[len(chain.blocks)-1]
	//création du nouveau block
	newBlock := CreateBlock(data, previousBlock.Hash)
	//ajout du block à la blockchain
	chain.blocks = append(chain.blocks, newBlock)
}

func Genesis() *Block {
	return CreateBlock("Genesis Block", []byte{})
}

func InitBlockchain() *Blockchain {
	//retourne l'adresse de la Blockchain avec un tableau de bloc et le block genesis
	return &Blockchain{[]*Block{Genesis()}, 2}
}

func main() {
	fmt.Println("Initialisation de la Blockchain")
	chain := InitBlockchain()

	fmt.Println("Ajout de blocks")
	chain.AddBlock("Block 1")
	chain.AddBlock("Block 2")
	chain.AddBlock("Block 3")

	for _, block := range chain.blocks {
		fmt.Println("***********************************************************")
		fmt.Printf("*Hash du block prévédent : %x\n", block.PrevHash)
		fmt.Printf("*Données du block : %s\n", block.Data)
		fmt.Printf("*Hash du block : %x\n", block.Hash)

		fmt.Println("         |\n         |\n         |\n")
		fmt.Println("***********************************************************")
	}

}
