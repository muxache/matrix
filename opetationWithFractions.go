package matrix

//Multiply combinates two elements according fraction multiplication rules
func Multiply(a, b Element) Element {
	var m Element
	m.Number = a.Number * b.Number
	m.Fraction = a.Fraction * b.Fraction
	return m
}

//Multiply combinates one element and single number according fraction multiplication rules
func MultiplyWithNumber(a Element, b float64) Element {
	var m Element
	m.Number = a.Number * (int64(b))
	m.Fraction = a.Fraction
	return m
}

//Division makes a division operation with frations
func Division(a, b Element) Element {
	var m Element
	m.Number = a.Number * b.Fraction
	m.Fraction = a.Fraction * b.Number
	return m
}

//Addition summarized two fractions
func Addition(a, b Element) Element {
	var m Element
	if a.Number == 0 {
		return b
	}
	if b.Number == 0 {
		return a
	}
	if a.Fraction == b.Fraction {
		m.Number = a.Number + b.Number
		m.Fraction = a.Fraction
		return m
	}
	m.Number = (a.Number * b.Fraction) + (b.Number * a.Fraction)
	m.Fraction = a.Fraction * b.Fraction
	return m
}

//Subtraction summarized two fractions
func Subtraction(a, b Element) Element {
	var m Element
	if a.Fraction == b.Fraction {
		m.Number = a.Number - b.Number
		m.Fraction = a.Fraction
		return m
	}
	m.Number = (a.Number * b.Fraction) - (b.Number * a.Fraction)
	m.Fraction = a.Fraction * b.Fraction
	return m
}

//Decimal returns decimal number after division of fraction
func (e Element) Decimal() float64 {
	return float64(float64(e.Number) / float64(e.Fraction))
}
