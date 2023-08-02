package main

import("fmt")
func main(){
	var word string
	var morse_word []string
	var text_to_morse = map[string]string{
		"a":".-",
		"b":"-...",
		"c":"-.-.",
		"d":"-..",
		"e":".",
		"f":"..-.",
		"g":"--.",
		"h":"....",
		"i":"..",
		"j":".---",
		"k":"-.-",
		"l":".-..",
		"m":"--",
		"n":"-.",
		"o":"---",
		"p":".--.",
		"q":"--.-",
		"r":".-.",
		"s":"...",
		"t":"-",
		"u":"..-",
		"v":"...-",
		"w":".--",
		"x":"-..-",
		"y":"-.--",
		"z":"--..",
		"1":".----",
		"2":"..---",
		"3":"...--",
		"4":"....-",
		"5":".....",
		"6":"-....",
		"7":"--...",
		"8":"---..",
		"9":"----.",
		"0":"-----"}

	fmt.Printf("Word: ")
	fmt.Scan(&word)

	for index,val := range word{
		morse_word = append(morse_word,text_to_morse[string(val)]+" ")
		fmt.Println(index)
	}

	for i:=0;i<len(morse_word);i++{
		print(morse_word[i])
	}
	fmt.Println("")


}
