package main

import "fmt"

const PAINT_PRICE = 10

func main() {

    walls := []struct {
        width float32
        height float32
        label string
    } {
        {0, 7.5, "1st wall"},
        {15.0, 7.5, "2nd wall"},
        {9.0, 7.5, "3rd wall"},
        {7.0, 7.5, "4th wall"},
    }

	for _, wall := range walls {
        cost, error := calculateWallPaintCost(wall.width, wall.height)
        if error != nil {
            fmt.Println("Error:", error)
        } else {
            fmt.Printf("%s cost %.2f\n", wall.label, cost)
        }
    }
    

	// 2nd wall cost
	// width, height = 15.0, 7.5
    // secondWallCost, error := calculateWallPaintCost(width, height)
    // if error != nil {
    //     fmt.Println("Error:", error)
    // } else {
    //     fmt.Printf("2nd wall cost %.2f\n", secondWallCost)
    // }

	// 3rd wall cost
	// width, height = 9.0, 7.5
    // thirdWallCost, error := calculateWallPaintCost(width, height)
    // if error != nil {
    //     fmt.Println("Error:", error)
    // } else {
    //     fmt.Printf("3rd wall cost %.2f\n", thirdWallCost)
    // }

	// 4th wall cost
	// width, height = 7.0, 7.5
    // fourthWallCost, error := calculateWallPaintCost(width, height)
    // if error != nil {
    //     fmt.Println("Error:", error)
    // } else {
    //     fmt.Printf("4th wall cost %.2f\n", fourthWallCost)
    // }
}

func calculateWallPaintCost(width, height float32) (float32, error) {
    if width <= 0 {
        return 0, fmt.Errorf("Width can't be zero or negative")
    }

    if height <= 0 {
        return 0, fmt.Errorf("Height can't be zero or negative")
    }

    area := width * height
    return area * PAINT_PRICE, nil
}