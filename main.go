package main

import (
	"fmt"
	"os"

	"foosoft.net/projects/jmdict"
)

func read() (jmdict.Jmdict, error) {
	reader, err := os.Open("dictionary/JMdict_e")
	if err != nil {
		fmt.Println(err)
		return jmdict.Jmdict{}, err
	}
	defer reader.Close()
	dictionary, _, err := jmdict.LoadJmdictNoTransform(reader)
	if err != nil {
		fmt.Println(err)
		return jmdict.Jmdict{}, err
	}
	return dictionary, nil
}

func get_entry(entry jmdict.JmdictEntry) string {
	kanji := ""
	out := ""
	for _, k := range entry.Kanji {
		kanji += k.Expression + ", "
	}
	if len(kanji) != 0 {
		kanji = kanji[:len(kanji)-2]
	}
	reading := ""
	for _, r := range entry.Readings {
		reading += r.Reading + ", "
	}
	reading = reading[:len(reading)-2]
	gloss := ""
	for _, g := range entry.Sense {
		for _, s := range g.Glossary {
			gloss += s.Content + ", "
		}
	}
	gloss = gloss[:len(gloss)-2]
	if len(kanji) == 0 {
		out = fmt.Sprintf("「%s」%s\n", reading, gloss)
	} else {
		out = fmt.Sprintf("「%s ・ %s」%s\n", kanji, reading, gloss)
	}
	return out
}

func main() {
	dictionary, err := read()
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err := os.Create("res.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, entry := range dictionary.Entries {
		f.WriteString(get_entry(entry))
	}
	f.Close()
}
