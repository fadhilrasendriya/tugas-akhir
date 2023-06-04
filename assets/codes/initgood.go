func initGoods(contract *client.Contract) {
	fmt.Printf("\n--> Init Goods: Initiate 10 goods \n")
	prefixId := "good-"
	for i := 0; i < 10; i++ {

		goodId := fmt.Sprintf("%s%06d", prefixId, i)
		goodName := fmt.Sprintf("Good number %d", i)
		fmt.Printf("\n*** Submitting Transaction: CreateGood with params: goodId: %s, goodName: %s, unit: kg, importLimit: 100.00, exportLimit 200.0\n\n", goodId, goodName)
		timeStart := time.Now()
		_, err := contract.SubmitTransaction("CreateGood", goodId, goodName, "kg", "100.00", "200.0")
		secondDuration := time.Now().Sub(timeStart).Seconds()
		if err != nil {
			fmt.Println("\n*** Submit Transaction error")
			fmt.Println(err)
		}
		fmt.Printf("%s created in %f seconds\n", goodId, secondDuration)
	}
}
