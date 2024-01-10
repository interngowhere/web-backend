package main

import "github.com/interngowhere/web-backend/ent"

var SeedUser1 = ent.User{
	Email: "aiken@gmail.com",
	FirstName: "Aiken",
	LastName: "Tan",
	Username: "AT",
}

var SeedUser2 = ent.User{
	Email: "dueet@gmail.com",
	FirstName: "Dueet",
	LastName: "Tan",
	Username: "DT",
}

var SeedTopic = ent.Topic{
	Title: "Test Topic",
	Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
	ShortDescription: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
}

var SeedThread = ent.Thread{
	Title: "This is a test thread",
	Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
}

var SeedComment1 = ent.Comment{
	Content: "This is a test comment",
}

var SeedComment2 = ent.Comment{
	Content: "This is a reply to the test comment",
}

var SeedComment3 = ent.Comment{
	Content: "This is another reply to the test comment",
}

var SeedTag1 = ent.Tag{
	TagName: "Behavioural Interview",
}

var SeedTag2 = ent.Tag{
	TagName: "Online Assessment",
}

var SeedTag3 = ent.Tag{
	TagName: "Leetcode",
}