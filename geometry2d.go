package geometry2d

import (
    "image/color"
    "container/list"
)

type Point struct {
        X float64
        Y float64
}

type Line struct {
    P1 Point
    P2 Point
    color.RGBA
}

type Polyline struct {
    list.List
}


type Arc struct {
    Center Point
    BeginAngle float64
    EndAngle float64
    color.RGBA
}

type Circle struct {
    Center Point
    Radius float64
    color.RGBA
}

type Bezier struct {
    P1 Point
    P2 Point
    Control1 Point
    Control2 Point
    color.RGBA
}

