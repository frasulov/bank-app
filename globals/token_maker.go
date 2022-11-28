package globals

import (
	"BankApp/config"
	"BankApp/token"
	"fmt"
)

var TokenMaker token.Maker

func Inject() error {
	var err error
	TokenMaker, err = token.NewPasetoMaker(config.Configuration.SymmetricKey)
	if err != nil {
		return fmt.Errorf("cannot create token maker: %v", err.Error())
	}

	return nil
}
