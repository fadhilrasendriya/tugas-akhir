func createOrderFailed(contract *client.Contract) {
	fmt.Printf("\n--> Submit Transaction: CreateOrder, Creating an order \n")
	log.Println("INFO: creating order with goods that does not comply with rules")

	goods := []Good{
		{
			GoodId:   "good-000000",
			Quantity: "500",
			Unit:     "kg",
		},
		{
			GoodId:   "good-000001",
			Quantity: "6",
			Unit:     "kg",
		},
		...
	}
	for _, good := range goods {
		fmt.Println(good)
	}

	goodJson, _ := json.Marshal(goods)
	goodStr := string(goodJson)

	log.Println("INFO: Submitting invalid order")

	_, err := contract.SubmitTransaction("CreateOrder", "order-000001", "import", "US", "First Order", "An order description", goodStr)
	if err != nil {
		fmt.Println(err)
		fmt.Println("\n*** Transaction not committed")
		return
	}
	fmt.Println("\n*** Transaction committed successfully")
}
