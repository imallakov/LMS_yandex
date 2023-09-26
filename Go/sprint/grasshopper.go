package main

import (
	"errors"
)

type Grasshopper struct {
	place int
	target int
}

type Jumper interface {
    WhereAmI() int // выводит текущее положение кузнечика на линейке
    Jump() (int, error) // кузнечик прыгает к зерну. Выводит новое положение кузнечика, или ошибку, если он уже ест зерно
}

func PlaceJumper(place, target int) Jumper {
	var g Grasshopper
	g.place=place
	g.target=target
	return &g
}

func (g *Grasshopper) Jump() (int, error) {
	if g.place==g.target {
		return 0,errors.New("Error")
	}
	if g.target>g.place{
		g.place+=min(g.target-g.place,5)
	}else{
		g.place-=min(g.place-g.target,5)
	}
	return g.place,nil
}

func (g *Grasshopper) WhereAmI() int {
	return g.place
}