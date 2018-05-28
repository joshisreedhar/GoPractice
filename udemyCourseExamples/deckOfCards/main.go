package main

func main() {
	println("-----------Take a new deck of cards---------------")
	d := newDeck()
	println(d.toString())

	println("-----------shuffle and deal ---------------")
	sd := shuffle(d)
	h, rd := deal(sd, 11)
	println(h.toString())
	println(rd.toString())

	println("----------Saving to file------------")
	hf := "hand_deckfile"
	rdf := "deck_deckfile"
	save(h, hf)
	save(rd, rdf)

	println("--------------restoring from file -----------")
	rh := deckFromFile(hf)
	rrd := deckFromFile(rdf)
	println(rh.toString())
	println(rrd.toString())
}
