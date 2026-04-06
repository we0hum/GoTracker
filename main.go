package main

import "fmt"

func main() {
	orderId := 1001
	customerName := "Иван"
	isDelivered := false
	fmt.Println("Order ID:", orderId)
	fmt.Println("Customer Name:", customerName)
	fmt.Println("Is Delivered:", isDelivered)

	orderIDs := []int{}
	orderIDs = append(orderIDs, 101)
	orderIDs = append(orderIDs, 102)
	orderIDs = append(orderIDs, 103)
	fmt.Println("OrderIDs:", orderIDs)
	fmt.Println("len(orderIDs) =", len(orderIDs), "cap(orderIDs) =", cap(orderIDs))

	orderCount := map[string]int{
		"Alice": 5,
		"Bob":   1,
	}
	fmt.Println("OrderCount:", orderCount)

	isReadyToShip(orderCount, "Alice")
}

func isReadyToShip(orderCount map[string]int, name string) bool {
	res := false
	if count, ok := orderCount[name]; ok {
		if count > 2 {
			res = true
		}
		fmt.Printf("Is %s ready to ship? %t\n", name, res)
	}
	return res
}
