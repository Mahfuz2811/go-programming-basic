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
    EspressoPrice = 200
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

func displayMenu() {
    fmt.Println("--------- Self Service Coffee Shop ------------")

    fmt.Println("1. Cappuccino")
    fmt.Println("2. Americano")
    fmt.Println("3. Espresso")
    fmt.Println("4. View Items")
    fmt.Println("5. Confirm Purchase")
    fmt.Println("6. Exit")
}

func getChoice(reader *bufio.Reader) (int, error) {
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

    return strconv.Atoi(input)
}

func addItemToCart(choice int, quantity int, coffees map[int]Coffee, cart map[int]Cart) {
    item, exists := cart[choice]
    if exists {
        item.quantity += quantity
        item.amount = float32(item.quantity) * item.price
        cart[choice] = item
    } else {
        cart[choice] = Cart{
            name: coffees[choice].name,
            quantity: quantity,
            price: coffees[choice].price,
            amount: coffees[choice].price * float32(quantity),
        }
    }
}

func viewCart(cart map[int]Cart) {
    if len(cart) == 0 {
        fmt.Println("Empty bucket. Please select an item first.")
        return
    }
    
    var totalAmount float32 = 0.0
    fmt.Printf("%-15s %-15s %-15s %-15s\n", "Item Name", "Price", "Quantity", "Amount")
    for _, item := range cart {
        fmt.Printf("%-15s %-15.2f %-15d %-15.2f\n", item.name, item.price, item.quantity, item.amount)
        totalAmount += item.amount
    }
    fmt.Println("---------------------------------------------------------------")
    fmt.Printf("%-47s %.2f\n", "Total Amount", totalAmount)
}

func confirmPurchase(cart map[int]Cart) {
    viewCart(cart)

    var approval string
    fmt.Print("Type 'yes' for payment: ")
    fmt.Scanln(&approval)

    if approval == "yes" {
        fmt.Println("Thanks for shopping with us!")
        return
    }
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
            name: "Espresso",
            price: EspressoPrice,
        },
    }

    selectedItems := make(map[int]Cart)
    reader := bufio.NewReader(os.Stdin)

    for {
        displayMenu()
        fmt.Print("Choose an option: ")
        choice, error := getChoice(reader)
        if error != nil {
            fmt.Println("Invalid input. Please enter a number between 1 to 6")
            continue
        }

        switch choice {
        case 1, 2, 3:
            fmt.Print("Enter quantity: ")
            quantity, err := getChoice(reader)
            if err != nil || quantity <= 0 {
                fmt.Println("Invalid quantity. Please enter a valid quantity")
                continue
            }

            addItemToCart(choice, quantity, coffees, selectedItems)
        case 4:
            viewCart(selectedItems)
        case 5:
            confirmPurchase(selectedItems)
        case 6:
            fmt.Println("Existing from the coffee purchasing booth. Goodbye!")
            return
        default:
            fmt.Println("Invalid option. Please enter a number between 1 to 6")
        }
    }
}                                             