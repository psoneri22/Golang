package main

// func main() {
// 	var sum, count int
// 	for {
// 		fmt.Print("Enter an integer (or press Enter to finish): ")
// 		var input string
// 		_, err := fmt.Scanln(&input)
// 		if err != nil || input == "" {
// 			break // Exit loop if input is not an integer or if it's empty
// 		}

// 		num, err := strconv.Atoi(input)
// 		if err != nil {
// 			fmt.Println("Invalid input. Please enter an integer.")
// 			continue // Skip the rest of the loop and prompt for input again
// 		}

// 		sum += num
// 		count++
// 	}

// 	if count == 0 {
// 		fmt.Println("No numbers entered.")
// 		return
// 	}

// 	average := float64(sum) / float64(count)
// 	fmt.Printf("Sum: %d\n", sum)
// 	fmt.Printf("Average: %.2f\n", average)
// }
