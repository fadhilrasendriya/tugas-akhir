func readOrder(contract *client.Contract) {
	fmt.Println("\n--> Evaluate Transaction: ReadOrder: retrieving previously created order")
	res, _ := contract.EvaluateTransaction("ReadOrder", "order-000000")
	fmt.Println(string(res))
}
