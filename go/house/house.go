package house

import "strings"

func Embed(relPhrase, nounPhrase string) string {
	return relPhrase + " " + nounPhrase
}

func Verse(
	subject string,
	relPhrases []string,
	nounPhrase string,
) string {

	return Embed(subject, strings.Join(append(relPhrases, nounPhrase), " "))
}

func Song() string {
	return song

}
