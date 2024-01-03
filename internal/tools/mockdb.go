package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"leo": {
		AuthToken: "123",
		Username:  "leo",
	},
	"alex": {
		AuthToken: "456",
		Username:  "alex",
	},
	"jason": {
		AuthToken: "789",
		Username:  "jason",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"leo": {
		Coins:    100,
		Username: "leo",
	},
	"alex": {
		Coins:    200,
		Username: "alex",
	},
	"jason": {
		Coins:    300,
		Username: "jason",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
