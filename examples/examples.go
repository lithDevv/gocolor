package main 

import (
	"gocolor"
	"fmt"
) 

func main() {
	// normal colors
	fmt.Println("Normal Colors: ")
	fmt.Println(gocolor.Red("Some Red Text") + gocolor.Blank())
	fmt.Println(gocolor.Black("Some Black Text") + gocolor.Blank())
	fmt.Println(gocolor.Green("Some Green Text") + gocolor.Blank())
	fmt.Println(gocolor.Yellow("Some Yellow Text") + gocolor.Blank())
	fmt.Println(gocolor.Blue("Some Blue Text") + gocolor.Blank())
	fmt.Println(gocolor.Purple("Some Purple Text") + gocolor.Blank())
	fmt.Println(gocolor.Cyan("Some Cyan Text") + gocolor.Blank())
	fmt.Println(gocolor.White("Some White Text") + gocolor.Blank())

	// bold colors 
	fmt.Println("\nBold Colors: ")
	fmt.Println(gocolor.BlackBold("Some Bold Black Text") + gocolor.Blank())
	fmt.Println(gocolor.RedBold("Some Bold Red Text") + gocolor.Blank())
	fmt.Println(gocolor.GreenBold("Some Bold Green Text") + gocolor.Blank())
	fmt.Println(gocolor.YellowBold("Some Bold Yellow Text") + gocolor.Blank())
	fmt.Println(gocolor.BlueBold("Some Bold Blue Text") + gocolor.Blank())
	fmt.Println(gocolor.PurpleBold("Some Bold Purple Text") + gocolor.Blank())
	fmt.Println(gocolor.CyanBold("Some Bold Cyan Text") + gocolor.Blank())
	fmt.Println(gocolor.WhiteBold("Some Bold White Text") + gocolor.Blank())
	
	// underlined colors
	fmt.Println("\nUnderlined Colors: ")
	fmt.Println(gocolor.BlackUnderline("Some Underlined Black Text") + gocolor.Blank())
	fmt.Println(gocolor.RedUnderline("Some Underlined Red Text") + gocolor.Blank())
	fmt.Println(gocolor.GreenUnderline("Some Underlined Green Text") + gocolor.Blank())
	fmt.Println(gocolor.YellowUnderline("Some Underlined Yellow Text") + gocolor.Blank())
	fmt.Println(gocolor.BlueUnderline("Some Underlined Blue Text") + gocolor.Blank())
	fmt.Println(gocolor.PurpleUnderline("Some Underlined Purple Text") + gocolor.Blank())
	fmt.Println(gocolor.CyanUnderline("Some Underlined Cyan Text") + gocolor.Blank())
	fmt.Println(gocolor.WhiteUnderline("Some Underlined White Text") + gocolor.Blank())

	// high intensity colors
	fmt.Println("\nHigh Intensity Colors: ")
	fmt.Println(gocolor.HiBlack("Some High Intensity Black Text") + gocolor.Blank())
	fmt.Println(gocolor.HiRed("Some High Intensity Red Text") + gocolor.Blank())
	fmt.Println(gocolor.HiGreen("Some High Intensity Green Text") + gocolor.Blank())
	fmt.Println(gocolor.HiYellow("Some High Intensity Yellow Text") + gocolor.Blank())
	fmt.Println(gocolor.HiBlue("Some High Intensity Blue Text") + gocolor.Blank())
	fmt.Println(gocolor.HiPurple("Some High Intensity Purple Text") + gocolor.Blank())
	fmt.Println(gocolor.HiCyan("Some High Intensity Cyan Text") + gocolor.Blank())
	fmt.Println(gocolor.HiWhite("Some High Intensity White Text") + gocolor.Blank())

	// Bold high intensity colors
	fmt.Println("\nBold High Intensity Colors: ")
	fmt.Println(gocolor.BiBlack("Some Bold High Intensity Black Text") + gocolor.Blank())
	fmt.Println(gocolor.BiRed("Some Bold High Intensity Red Text") + gocolor.Blank())
	fmt.Println(gocolor.BiGreen("Some Bold High Intensity Green Text") + gocolor.Blank())
	fmt.Println(gocolor.BiYellow("Some Bold High Intensity Yellow Text") + gocolor.Blank())
	fmt.Println(gocolor.BiBlue("Some Bold High Intensity Blue Text") + gocolor.Blank())
	fmt.Println(gocolor.BiPurple("Some Bold High Intensity Purple Text") + gocolor.Blank())
	fmt.Println(gocolor.BiCyan("Some Bold High Intensity Cyan Text") + gocolor.Blank())
	fmt.Println(gocolor.BiWhite("Some Bold High Intensity White Text") + gocolor.Blank())
}
