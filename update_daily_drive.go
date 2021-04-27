package main

import (
	"log"
)

func updateDailyDrive() {
	log.Println("Running Daily Drive Update")
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
