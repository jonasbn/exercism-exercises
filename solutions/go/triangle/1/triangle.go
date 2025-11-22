package triangle

import ( "math" )

// Kind is a type for supported triangle types
type Kind string

// polygon is a type for possibly implementing a triangle
type polygon struct {
    a float64
    b float64
    c float64
}

const (
    // NaT not a triangle
    NaT = "NaT"
    // Equ indicated equilateral triangle
    Equ = "Equ"
    // Iso indicated isosceles triangle
    Iso = "Iso"
    // Sca indicated scalene triangle
    Sca = "Sca"
)

// KindFromSides should have a comment documenting ip.
func KindFromSides(a, b, c float64) Kind {
    var k Kind
    var p polygon = polygon{a, b, c}

    if hasNaN(p) {
        return NaT
    }

    if hasInf(p) {
        return NaT
    }

    if hasZero(p) {
        return NaT
    }

    if isInequal(p) {
        return NaT
    }

    if isEquilateral(p) {
        k = Equ
    } else if isIsosceles(p) {
        k = Iso
    } else if isScalene(p) {
        k = Sca
    }

	return k
}

func isScalene(p polygon) bool {
    if p.a != p.b && p.a != p.c && p.b != p.c {
        return true
    }

    return false
}

func isIsosceles(p polygon) bool {
    if p.a == p.b || p.a == p.c || p.b == p.c {
        return true
    }

    return false
}

func isEquilateral(p polygon) bool {
    if p.a == p.b && p.a == p.c && p.b == p.c {
        return true
    }

    return false
}

func isInequal(p polygon) bool {
    if p.a + p.b < p.c || p.a + p.c < p.b || p.b + p.c < p.a {
        return true;
    }

    return false;
}

func hasZero(p polygon) bool {
    if p.a <= 0 || p.b <= 0 || p.c <= 0 {
        return true
    }

    return false;
}

func hasInf(p polygon) bool {
    if math.IsInf(p.a, 0) { return true }
    if math.IsInf(p.b, 0) { return true }
    if math.IsInf(p.c, 0) { return true }

    return false
}

func hasNaN(p polygon) bool {
    if math.IsNaN(p.a) { return true }
    if math.IsNaN(p.b) { return true }
    if math.IsNaN(p.c) { return true }

    return false
}
