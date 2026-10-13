package main

import (
	"log"
	"github.com/staritfarm/fintech-trading-engine/backend/internal/domain"
)

func main() {
	log.Println("Initializing ApexTrade High-Frequency Engine...")
	engine := domain.NewExecutionEngine()
	log.Printf("ApexTrade Engine ready. State: %s\n", engine.GetStatus())
}
