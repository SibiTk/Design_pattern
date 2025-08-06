package main

import (
	"fmt"
)


type EnglishSpeaker interface {
	SpeakEnglish() string
}

type TamilSpeaker struct{}


func (s TamilSpeaker) TamilSpeaker() string {
	return "vanakam , Epdi irukinga"
}


type TranslatorAdapter struct {
	tamilspeaker TamilSpeaker
}

func (t TranslatorAdapter) SpeakEnglish() string {

	tamil := t.tamilspeaker.TamilSpeaker()
	return "Hello, how are you? (Translated from: English To Tamil " + tamil + " )"
}


func main() {
	
	tamil := TamilSpeaker{}


	var englishSpeaker EnglishSpeaker = TranslatorAdapter{tamilspeaker: tamil}

	
	fmt.Println(englishSpeaker.SpeakEnglish())
}

