package gocolor

var (
	// normal colors
	RED    = "\033[31m"
	BLACK  = "\033[0;30m"
	GREEN  = "\033[0;32m"
	YELLOW = "\033[0;33m"
	BLUE   = "\033[0;34m"
	BLANK  = "\033[0m"
	PURPLE = "\033[0;35m"
	CYAN   = "\033[0;36m"
	WHITE  = "\033[0;37m"
	// bold colors 
	BLACK_BOLD  =  "\033[1;30m"
	RED_BOLD    =  "\033[1;31m"
	GREEN_BOLD  =  "\033[1;32m"
	YELLOW_BOLD =  "\033[1;33m"
	BLUE_BOLD   =  "\033[1;34m"
	PURPLE_BOLD = "\033[1;35m"
	CYAN_BOLD   = "\033[1;36m"
	WHITE_BOLD  = "\033[1;37m"
	// underline colors
	BLACK_UNDERLINE  = "\033[4;30m"
	RED_UNDERLINE    = "\033[4;31m"
	GREEN_UNDERLINE  = "\033[4;32m"
	YELLOW_UNDERLINE = "\033[4;33m"
	BLUE_UNDERLINE   = "\033[4;34m"
	PURPLE_UNDERLINE = "\033[4;35m"
	CYAN_UNDERLINE   = "\033[4;36m"
	WHITE_UNDERLINE  = "\033[4;37m"
	// high intensity colors 
	HI_BLACK  = "\033[0;90m"
	HI_RED    = "\033[0;91m"
	HI_GREEN  = "\033[0;92m"
	HI_YELLOW = "\033[0;93m"
	HI_BLUE   = "\033[0;94m"
	HI_PURPLE = "\033[0;95m"
	HI_CYAN   = "\033[0;96m"
	HI_WHITE  = "\033[0;97m"
	// bold high intesity colors
	BI_BLACK  = "\033[1;90m"
	BI_RED    = "\033[1;91m"
	BI_GREEN  = "\033[1;92m"
	BI_YELLOW = "\033[1;93m"
	BI_BLUE   = "\033[1;94m"
	BI_PURPLE = "\033[1;95m"
	BI_CYAN   = "\033[1;96m"
	BI_WHITE  = "\033[1;97m"
)
 
func Red(text string) string{
	return RED + text
}

func Black(text string) string{
	return BLACK + text
}

func Green(text string) string{
	return GREEN + text
}

func Yellow(text string) string{
	return YELLOW + text
}

func Blue(text string) string{
	return BLUE + text
}

func Purple(text string) string{
	return PURPLE + text
}

func Cyan(text string) string{
	return CYAN + text
}

func White(text string) string{
	return WHITE + text
}

func BlackBold(text string) string{
	return BLACK_BOLD + text
}

func RedBold(text string) string{
	return RED_BOLD + text
}

func GreenBold(text string) string{
	return GREEN_BOLD + text
}

func YellowBold(text string) string{
	return YELLOW_BOLD + text
}

func BlueBold(text string) string{
	return BLUE_BOLD + text
}

func PurpleBold(text string) string{
	return PURPLE_BOLD + text
}

func CyanBold(text string) string{
	return CYAN_BOLD + text
}

func WhiteBold(text string) string{
	return WHITE_BOLD + text
}

func BlackUnderline(text string) string{
	return BLACK_UNDERLINE + text
}

func RedUnderline(text string) string{
	return RED_UNDERLINE + text
}

func GreenUnderline(text string) string{
	return GREEN_UNDERLINE + text
}

func YellowUnderline(text string) string{
	return YELLOW_UNDERLINE + text
}

func BlueUnderline(text string) string{
	return BLUE_UNDERLINE + text
}

func PurpleUnderline(text string) string{
	return PURPLE_UNDERLINE + text
}

func CyanUnderline(text string) string{
	return CYAN_UNDERLINE + text
}

func WhiteUnderline(text string) string{
	return WHITE_UNDERLINE + text
}

func HiBlack(text string) string{
	return HI_BLACK + text
}

func HiRed(text string) string{
	return HI_RED + text
}

func HiGreen(text string) string{
	return HI_GREEN + text
}

func HiYellow(text string) string{
	return HI_YELLOW + text
} 

func HiBlue(text string) string{
	return HI_BLUE + text
}

func HiPurple(text string) string{
	return HI_PURPLE + text
}

func HiCyan(text string) string{
	return HI_CYAN + text
}

func HiWhite(text string) string{
	return HI_WHITE + text
}

func BiBlack(text string) string{
	return BI_BLACK + text
}

func BiRed(text string) string{
	return BI_RED + text
}

func BiGreen(text string) string{
	return BI_GREEN + text
} 

func BiYellow(text string) string{
	return BI_YELLOW + text
}

func BiBlue(text string) string{
	return BI_BLUE + text
}

func BiPurple(text string) string{
	return BI_PURPLE + text
}

func BiCyan(text string) string{
	return BI_CYAN + text
} 

func BiWhite(text string) string{
	return BI_WHITE + text
}

func Blank() string{
	return BLANK
}