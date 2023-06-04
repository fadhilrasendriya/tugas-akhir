func createOrderConcurrently(contract *client.Contract, n int) {
	fmt.Printf("\n--> Submit Transaction: submitting %d concurrent requests\n", n)
	c := make(chan bool)
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
	goodsJson, _ := json.Marshal(goods)
	goodsStr := string(goodsJson)
	timeStart := time.Now()
	for i := 0; i < n; i++ {
		orderId := fmt.Sprintf("order-%06d", i)
		go func(contract *client.Contract, orderId, goods string, c chan bool) {
			_, err := contract.SubmitTransaction("CreateOrder", orderId, "import", "US", "concurrent order", "Order description", goods)
			if err != nil {
				c <- false
				return
			}

			c <- true
		}(contract, orderId, goodsStr, c)
	}
	failedTx := 0
	for i := 0; i < n; i++ {
		res := <-c
		if res == false {
			failedTx++
		}
	}
	secondsDuration := time.Now().Sub(timeStart).Seconds()
	if failedTx > 0 {
		fmt.Printf("*** There are %d failed request(s) on %d concurrent request(s)\n", failedTx, n)
	}
	fmt.Printf("*** %d concurrent request(s) finished in %f seconds\n", n, secondsDuration)
}
