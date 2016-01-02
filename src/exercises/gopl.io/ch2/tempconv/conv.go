package tempconv

// Conversion functions
func CToK(c Celsius) Kelvin {
    return Kelvin(c - AbsoluteZeroC)
}

func KToC(k Kelvin) Celsius {
    return Celsius(k) + AbsoluteZeroC
}

func FToC(f Fahrenheit) Celsius {
    return Celsius((f - 32) * 5 / 9)
}

func CToF(c Celsius) Fahrenheit {
    return Fahrenheit(c * 9 / 5 + 32)
}

func FToK(f Fahrenheit) Kelvin {
    return CToK(FToC(f))
}

func KToF(k Kelvin) Fahrenheit {
    return CToF(KToC(k))
}