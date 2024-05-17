package main

import (
	"fmt"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/bnb-chain/tss/client"
	"github.com/bnb-chain/tss/common"
	_ "net/http/pprof"
	"os"
	"path"
)

func loadSavedKey(home, vault, password string) keygen.LocalPartySaveData {
	wPriv, err := os.OpenFile(path.Join(home, vault, "sk.json"), os.O_RDONLY, 0400)
	if err != nil {
		common.Panic(err)
	}
	defer wPriv.Close()
	wPub, err := os.OpenFile(path.Join(home, vault, "pk.json"), os.O_RDONLY, 0400)
	if err != nil {
		common.Panic(err)
	}
	defer wPub.Close()

	result, _, err := common.Load(password, wPriv, wPub) // TODO: validate nodeKey
	if err != nil {
		common.Panic(err)
	}
	return *result
}

func main() {
	/*
		go func() {
			log.Println(http.ListenAndServe("localhost:6062", nil))
		}()
	*/
	keyShare1 := loadSavedKey("./.test1", "default", "123456789")
	keyShare2 := loadSavedKey("./.test2", "default", "123456789")

	keyShares := make([]keygen.LocalPartySaveData, 2)
	keyShares[0] = keyShare1
	keyShares[1] = keyShare2
	privKey, _ := client.Reconstruct(1, tss.S256(), keyShares)
	fmt.Println(client.GetAddress(privKey.PublicKey, ""))

	//cmd.Execute()
}
