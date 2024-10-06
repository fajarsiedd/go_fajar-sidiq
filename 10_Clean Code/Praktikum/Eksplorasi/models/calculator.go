package models

type Calculator struct {
	NumA int `json:"a"`
	NumB int `json:"b"`
}

func (c *Calculator) Add() int {
	return c.NumA + c.NumB
}

func (c *Calculator) Subtract() int {
	return c.NumA - c.NumB
}

func (c *Calculator) Multiply() int {
	return c.NumA * c.NumB
}

func (c *Calculator) Divide() int {
	if c.NumB == 0 {
		return 0
	}
	return c.NumA / c.NumB
}
