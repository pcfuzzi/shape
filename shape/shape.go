package main





// Definition der Rectangle-Struktur
type Rectangle struct {
    Width  float32
    Height float32
}
// Implementierung der Perimeter-Methode für Rectangle
func (r Rectangle) Perimeter() float32 {
    return 2*r.Width + 2*r.Height
}

func (r Rectangle) Area() float32{
return 2*r.Width + 2*r.Height 
}