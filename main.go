package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	
	"multicall2/exemplo"
	"github.com/ethereum/go-ethereum/common"
)

type BalanceRequest struct {
	Addresses []string `json:"addresses"`
}

type BalanceResponse struct {
	Status   string           `json:"status"`
	Balances []AddressBalance `json:"balances"`
}

type AddressBalance struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
}

func getBalanceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var req BalanceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Erro ao decodificar request: "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(req.Addresses) == 0 {
		http.Error(w, "Nenhum endereço fornecido", http.StatusBadRequest)
		return
	}

	// Converter strings para *common.Address
	addresses := make([]*common.Address, len(req.Addresses))
	for i, addr := range req.Addresses {
		address := common.HexToAddress(addr)
		addresses[i] = &address
	}

	balances, err := exemplo.GetWETHBalance(addresses)
	if err != nil {
		http.Error(w, "Erro ao obter saldos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := BalanceResponse{
		Status:   "success",
		Balances: make([]AddressBalance, len(req.Addresses)),
	}

	for i, addr := range req.Addresses {
		response.Balances[i] = AddressBalance{
			Address: addr,
			Balance: balances[i].String(),
		}
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Erro ao codificar resposta: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	http.HandleFunc("/getBalance", getBalanceHandler)

	port := ":8080"
	fmt.Printf("Servidor iniciado na porta %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}
}
