package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func roastmeCommandCreate(s *discordgo.Session, m *discordgo.MessageCreate){
	if m.Content == "!roastme" {
		createMeCommand(s,m)}
	if m.Content == "!RoastMe" {
		createMeCommand(s,m)}
	if m.Content == "!roastMe" {
		createMeCommand(s,m)}
	if m.Content == "!Roastme" {
		createMeCommand(s,m)}
	if m.Content == "!ROASTME" {
		createMeCommand(s,m)}

	if m.Content == "Roastme" {
		createMeCommand(s,m)}
	if m.Content == "ROASTME" {
		createMeCommand(s,m)}
	if m.Content == "roastme"{
		createMeCommand(s,m)}
}

func roastyouCommandCreate(s *discordgo.Session, m *discordgo.MessageCreate) {	// Implement random roast person if no user mentioned
	if strings.Contains(m.Content, "roastyou") && !isUserMentioned(m.Content) {
		log.Println("in the wrong section")
		s.ChannelMessageSend(m.ChannelID, createRoast())} // + randomPersonArray
	if strings.Contains(m.Content, "Roastyou") && !isUserMentioned(m.Content) {
		s.ChannelMessageSend(m.ChannelID, createRoast())} // + randomPersonArray
	if strings.Contains(m.Content, "RoastYou") && !isUserMentioned(m.Content) {
		s.ChannelMessageSend(m.ChannelID, createRoast())} // + randomPersonArray
	if strings.Contains(m.Content, "!roastyou") && !isUserMentioned(m.Content) {
	s.ChannelMessageSend(m.ChannelID, createRoast())} // + randomPersonArray

}
//	creates the mention roasts
func mentionsRoastYouCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.Contains(m.Content, "roastyou") && isUserMentioned(m.Content) {
		log.Println(m.Content, " printing content of message")
		createMentionYouCommand(s, m)}
	if strings.Contains(m.Content, "Roastyou") && isUserMentioned(m.Content) {
		log.Println(m.Content, " printing content of message")
		createMentionYouCommand(s, m)}
	if strings.Contains(m.Content, "RoastYou") && isUserMentioned(m.Content) {
		log.Println(m.Content, " printing content of message")
		createMentionYouCommand(s, m)}
	if strings.Contains(m.Content, "!roastyou") && isUserMentioned(m.Content) {
		log.Println(m.Content, " printing content of message")
		createMentionYouCommand(s, m)}}

func createMeCommand(s *discordgo.Session, m *discordgo.MessageCreate){
	s.ChannelMessageSend(m.ChannelID, createRoast() + MentionUser(m.Author.ID))}

func createMentionYouCommand(s *discordgo.Session, m *discordgo.MessageCreate){
	s.ChannelMessageSend(m.ChannelID, createRoast() + returnUsersMentioned(m.Content))} //

func roastmeCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	roastmeCommandCreate(s, m)} // Inits the roastme commands

func roastyouCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	roastyouCommandCreate(s, m)} // Inits the roastyou commands