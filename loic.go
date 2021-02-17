package main

import (
	"fmt"

	"is4you"
)

func main() {
	message := `
	wxdb bxvvnb mndg oaèanb sdvnjdg zdr erexwb à ynd mn mrbcjwln.
	wxdb exhxwb unb krnwb nc unb vjdg vjrb wxdb wn yjauxwb zd'nw brunwln.
	
	jccnwcrxw: rub bxwc erejwcb vjrb ln wn bxwc wr mnb ynabxwwnb, wr mnb jwrvjdg, wr mnb eépécjdg.
	`
	decodedMessage := is4you.Traduce(message)
	fmt.Println(fmt.Sprintf(`
	Bonsoir à toi mon loic,

	Voici une enigme qui te permettra de trouver un nouvel indice. Par contre, le système vient de coder le message contenant l'enigme.
	Je compte sur toi pour implémenter la fonction Traduce de is4you.go et, ainsi, pouvoir résoudre cette enigme ;)

	Voici ton enigme :
	----
	%v
	----
	`, decodedMessage))
}
