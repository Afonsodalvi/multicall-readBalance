package exemplo

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"github.com/omnes-tech/multicall"
)

func GetWETHBalance(addresses []*common.Address) ([]*big.Int, error) {
	///Configuracao logging
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Ldate)

	//Conectar ao RPC
	rpcURL := "https://eth.llamarpc.com"
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RPC: %v", err)
	}
	log.Println("Connected to RPC")

	//Iniciar multicall
	mcall, err := multicall.NewMultiCall(multicall.GENERAL, client, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize multicall: %v", err)
	}

	// Endereço do contrato WETH
	wethAddress := common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	
	// Preparar os arrays para as chamadas
	targets := make([]common.Address, len(addresses))
	funcSigs := make([]string, len(addresses))
	argss := make([][]any, len(addresses))
	returnTypes := make([][]string, len(addresses))

	// Preencher os arrays com os dados necessários
	for i := 0; i < len(addresses); i++ {
		targets[i] = wethAddress
		funcSigs[i] = "balanceOf(address)"
		argss[i] = []any{addresses[i]}
		returnTypes[i] = []string{"uint256"}
	}

	// Criar as chamadas multicall
	calls := multicall.NewCalls(targets, funcSigs, argss, returnTypes, nil)

	// Executar as chamadas
	results, err := mcall.AggregateStatic(calls, client)
	if err != nil {
		return nil, fmt.Errorf("failed to execute multicall: %v", err)
	}

	// Processar os resultados
	balances := make([]*big.Int, len(addresses))
	resultArray := results.Result.([]interface{})
	for i, address := range addresses {
		balance := resultArray[i].([]interface{})[0].(*big.Int)
		balances[i] = balance
		log.Printf("Balance for address %s: %s", address.Hex(), balance.String())
	}

	return balances, nil
}