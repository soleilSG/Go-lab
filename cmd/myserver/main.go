package main

import (
	"fmt"

	"github.com/soleilSG/Go-lab/internal/database"
	"github.com/soleilSG/Go-lab/pkg/calculator"
)

func main() {
	sum := calculator.Add(5, 7)
	fmt.Printf("5 + 7 = %d\n", sum)

	dbStatus := database.Connect()
	fmt.Println(dbStatus)
}
