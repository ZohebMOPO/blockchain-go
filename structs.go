package main

type Block struct {
	Index     int
	Timestamp string
	text      string
	Hash      string
	PrevHash  string
}

type Message struct {
	text string
}
