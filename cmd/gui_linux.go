package cmd

import (
	"strings"
)

func LaunchIDPrompt() {
	teamID, err := shellCommandOutput(`
		#!/bin/bash
		teamid=$(
			zenity --entry \
			--title="TeamID" \
			--text "Enter your TeamID:" \
            --width=700
		)
		echo $teamid
	`)

	strippedTeamID := strings.Replace(teamID, " ", "", -1)

	if err == nil {
		writeFile(mc.DirPath+"TeamID.txt", strippedTeamID)
	} else {
		failPrint("Error saving TeamID!")
		sendNotification("Error saving TeamID!")
	}
}

func LaunchConfigGui() {
	warnPrint("The script doesn't currently have the ability to add multiple check or fail conditions-- you must still do these manually.")
	_, err := shellCommandOutput("bash ./misc/gui_linux.sh")
	if err == nil {
		infoPrint("Configuration successfully written to scoring.conf!")
	}
}
