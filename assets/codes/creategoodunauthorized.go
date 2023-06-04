func createGoodUnauthorized(contract *client.Contract) {
	fmt.Println("\n--> Submit Transaction: CreateGood: creating good from unauthorized organization: tradingcompany1.example.com\n")
	_, err := contract.SubmitTransaction("CreateGood", "good-100123", "Good number 100123", "kg", "100.00", "200.0")
	if err != nil {
		fmt.Println(err)
		fmt.Println("*** Transaction failed")
	}
}
