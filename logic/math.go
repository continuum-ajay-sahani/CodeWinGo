package logic

//Math conatain math logic
type Math struct {
}

//Addition add number
func (m Math) Addition(a, b float64) float64 {
	return a + b
}

//Subtraction subtract number
func (m Math) Subtraction(a, b float64) float64 {
	return a - b
}

//Multiplication multiply number
func (m Math) Multiplication(a, b float64) float64 {
	return a * b
}

//Division divide number
func (m Math) Division(a, b float64) float64 {
	return a / b
}
