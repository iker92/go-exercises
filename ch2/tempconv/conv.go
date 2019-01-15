package tempconv

func CtoF(c Celsius) Fahreneit {
	return Fahreneit(c*9/5 + 32)
}

func FtoC(f Fahreneit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CtoK(c Celsius) Kelvin {
	return Kelvin(c - 273.15)
}

func FtoK(f Fahreneit) Kelvin {
	return Kelvin(((f - 32) * 5 / 9) - 273.15)
}

func KtoC(k Kelvin) Celsius {
	return Celsius(k + 273.15)
}

func KtoF(k Kelvin) Fahreneit {
	return Fahreneit((k+273.15)*9/5 + 32)
}
