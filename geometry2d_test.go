package geometry2d

import "testing"

func TestEntitiesEqual(t *testing.T) {
    p1 := Point{10.5, 20.5}
    p2 := Point{10.5, 20.5}
    p3 := Point{20.5, 10.5}

    if p1 != p2 {
        t.Errorf("P1(%p) != P2(%p)", p1, p2)
    }
    if p1 == p3 {
        t.Errorf("P1(%p) == P3(%p)", p1, p3)
    }

    p4 := Point{20.5, 10.5}
    p5 := Point{30.5, 10.5}
    l1 := Line{P1: p1, P2: p3}
    l2 := Line{P1: p2, P2: p4}
    l3 := Line{P1: p1, P2: p5}

    if l1 != l2 {
        t.Error("L1(%l) != L2(%l)", l1, l2)
    }
    if l1 == l3 || l2 == l3 {
        t.Error("l1(%l) == l5(%l)", l1, l3)
    }
}
