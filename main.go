package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
    CappuccinoPrice = 150
    AmericanoPrice = 250
    ExpressoPrice = 200
)

type Cart struct{
    name string
    quantity int
    price float32
    amount float32
}

type Coffee struct{
    name string;
    price float32;
}

func main() {

    coffees := map[int]Coffee {
        1: {
            name: "Cappuccino",
            price: CappuccinoPrice,
        },
        2: {
            name: "Americano",
            price: AmericanoPrice,
        },
        3: {
            name: "Expresso",
            price: ExpressoPrice,
        },
    }

    selectedItems := make(map[int]Cart)
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Println("--------- Self Service Coffee Shop ------------")

        fmt.Println("1. Cappuccino")
        fmt.Println("2. Americano")
        fmt.Println("3. Expresso")
        fmt.Println("4. View Items")
        fmt.Println("5. Confirm Purchase")
        fmt.Println("6. Exit")

        fmt.Print("Choose an option: ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        choice, error := strconv.Atoi(input)
        if error != nil {
            fmt.Println("Invalid input. Please enter a number between 1 to 6")
            continue
        }

        switch choice {
        case 1, 2, 3:
            fmt.Print("Enter quantity: ")
            quantityInput, _ := reader.ReadString('\n')
            quantityInput = strings.TrimSpace(quantityInput)
            quantity, err := strconv.Atoi(quantityInput)
            if err != nil || quantity <= 0 {
                fmt.Println("Invalid quantity. Please enter a valid quantity")
                continue
            }

            item, exists := selectedItems[choice]
            if exists {
                item.quantity += quantity
                item.amount = float32(item.quantity) * item.price
                selectedItems[choice] = item
            } else {
                selectedItems[choice] = Cart{
                    name: coffees[choice].name,
                    quantity: quantity,
                    price: coffees[choice].price,
                    amount: coffees[choice].price * float32(quantity),
                }
            }
        case 4:
            if len(selectedItems) == 0 {
                fmt.Println("Empty bucket. Please select an item first.")
                continue
            }
            
            var totalAmount float32 = 0.0
            fmt.Printf("%-15s %-15s %-15s %-15s\n", "Item Name", "Price", "Quantity", "Amount")
            for _, item := range selectedItems {
                fmt.Printf("%-15s %-15.2f %-15d %-15.2f\n", item.name, item.price, item.quantity, item.amount)
                totalAmount += item.amount
            }
            fmt.Println("---------------------------------------------------------------")
            fmt.Printf("%-47s %.2f\n", "Total Amount", totalAmount)
        case 5:
            if len(selectedItems) == 0 {
                fmt.Println("Empty bucket. Please select an item first.")
                continue
            }

            var totalAmount float32 = 0.0
            fmt.Printf("%-15s %-15s %-15s %-15s\n", "Item Name", "Price", "Quantity", "Amount")
            for _, item := range selectedItems {
                fmt.Printf("%-15s %-15.2f %-15d %-15.2f\n", item.name, item.price, item.quantity, item.amount)
                totalAmount += item.amount
            }
            fmt.Println("---------------------------------------------------------------")
            fmt.Printf("%-47s %.2f\n", "Total Amount", totalAmount)

            var approval string
            fmt.Print("Type 'yes' for payment: ")
            fmt.Scanln(&approval)

            if approval == "yes" {
                fmt.Println("Thanks for shopping with us!")
                return
            }
        case 6:
            fmt.Println("Existing from the coffee purchasing booth. Goodbye!")
            return
        default:
            fmt.Println("Invalid option. Please enter a number between 1 to 6")
        }
    }
}                                             