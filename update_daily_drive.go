package main

import (
	"log"
)

func updateDailyDrive() {
	log.Println("Running Daily Drive Update")

	// Extract the token.
	token, err := client.Token()
	if err != nil {
		log.Fatal(err)
	}
	// Load token again and create client from it.
	client = clientFromToken(token)

	user, err := client.CurrentUser()
	if err != nil {
		log.Fatal(err)
	}

	dailyDrive := findDailyDrive(&client)
	myDailyDrive := findMyDailyDrive(&client)
	if myDailyDrive == nil {
		createMyDailyDrivePlaylist(&client, user)
		myDailyDrive = findMyDailyDrive(&client)
	}
	clearMyDailyDrive(&client, myDailyDrive)
	fillMyDailyDrive(&client, dailyDrive, myDailyDrive)
}
