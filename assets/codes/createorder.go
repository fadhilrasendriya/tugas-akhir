func createOrderSuccess(contract *client.Contract) {

	fmt.Printf("\n--> Submit Transaction: CreateOrder, Creating an order \n")
	log.Println("INFO: creating order with 5 goods that comply with the rule")
	goods := []Good{
		{
			GoodId:   "good-000000",
			Quantity: "5",
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
	fmt.Println("\n*** Submitting transaction with params: orderId: order-000000, orderType: import, country: US, name: First Order, description: An order description,\ngoods: " + goodStr)
	timeStart := time.Now()
	_, err := contract.SubmitTransaction("CreateOrder", "order-000000", "import", "US", "First Order", "An order description", goodStr)
	timeDuration := time.Now().Sub(timeStart)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\n*** Transaction committed successfully")
	fmt.Printf("\n*** Create Order duration: %f seconds\n", timeDuration.Seconds())
}
