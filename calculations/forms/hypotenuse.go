package forms
import "math"
// Erwartet die Längen der Katheten eines rechtwinkligen Dreiecks.
// Liefert die Länge der Hypotenuse.
func Hypotenuse(a, b float64) float64 {
	n := (a*a)+(b*b)
	m :=math.Sqrt(n)
	return m
}
