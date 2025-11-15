package main

import (
	"log"
	"strings"
)

func MentionUser(AuthorID string) string{
	return "<@!" + AuthorID + ">"
}

func isUserMentioned(content string) bool { // Store to user[]
	if strings.Contains(content, "<@!") {
		log.Println("Mentioned Found!")
		return true
	} else {
		return false}}

func returnUsersMentioned(content string) string { // Users signature is 18 characters long
	var mentionArray string
	//var tight []string
	var signature string
	splitContent := strings.Split(content, "<@!")
	for i := 1; i <= len(splitContent)-1; i++ { //Start at i = 1 to skip text before the mention
		if len(splitContent[i]) > 18 {
			signature = splitContent[i][0:18]
			mentionArray += "<@!" + signature + ">"
		}
	}
	/*	TODO fix ths bugs I think its working right now though
	 */
	log.Println(mentionArray) // FOR DEBUGGING!
	return mentionArray
}
