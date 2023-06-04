func readGoods(contract *client.Contract) {
	fmt.Println("\n--> Read Goods: read previously created 10 goods\n")
	for i := 0; i < 10; i++ {
		timeStart := time.Now()
		goodId := fmt.Sprintf("good-%06d", i)
		fmt.Printf("\n*** Reading good %s\n", goodId)
		evaluateResult, err := contract.EvaluateTransaction("GetGoodById", goodId)
		secondsDuration := time.Now().Sub(timeStart).Seconds()
		if err != nil {
			fmt.Println("\n*** Failed to get data")
			fmt.Println(err)
		}

		result := string(evaluateResult)
		fmt.Println("\n*** Retrieved good: " + result)
		fmt.Printf("time to read: %f seconds\n", secondsDuration)
	}
}
