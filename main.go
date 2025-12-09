package main

import "fmt"

func main() {

var loan_amount  int

print("Enter loan amount: ")
fmt.Scan(&loan_amount)


var interest_rate int

print("Enter interest rate: ")
fmt.Scan(&interest_rate)

var yeasrs int

print("Enter time period(years): ")
fmt.Scan(&yeasrs)

var installment_amount int

installment_amount = (loan_amount * interest_rate * yeasrs / 100 + loan_amount) / (yeasrs * 12)
fmt.Printf("Installment amount is: %d\n", installment_amount)

  
}
