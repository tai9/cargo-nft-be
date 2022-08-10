package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/tai9/cargo-nft-be/api"
	db "github.com/tai9/cargo-nft-be/db/sqlc"
	"github.com/tai9/cargo-nft-be/utils"
	"github.com/thirdweb-dev/go-sdk/thirdweb"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	// Creates a new SDK instance to get read-only data for your contracts, you can pass:
	// - a chain name (mainnet, rinkeby, goerli, polygon, mumbai, avalanche, fantom)
	// - a custom RPC URL
	thirdwebSdk, err := thirdweb.NewThirdwebSDK("rinkeby", &thirdweb.SDKOptions{
		PrivateKey: config.ThirdwebPrivateKey,
	})
	if err != nil {
		panic(err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store, thirdwebSdk)
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
