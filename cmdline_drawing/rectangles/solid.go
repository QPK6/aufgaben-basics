package rectangles

import (
	"fmt"
	"strings"
)

// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Das Rechteck soll komplett mit `#`-Zeichen gefüllt sein.
func DrawSolidRectangle(height, width int) {
	// TODO
	for row := 0; row<height; row++{
		// Zeichne eine Zeile
		//fmt.Println(strings.Repeat("a",width)) -------- strings package
			for column :=0; column < width; column++{
				fmt.Print("#")
			}

		fmt.Println()
	} 
}
