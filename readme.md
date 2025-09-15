# Live-Code-1 Phase-1
## RULES
1. **Untuk kampus remote**: **WAJIB** melakukan **share screen**(**DESKTOP/ENTIRE SCREEN**) dan **unmute microphone** ketika Live Code
berjalan (tidak melakukan share screen/salah screen atau tidak unmute microphone akan di ingatkan).
2. Kerjakan secara individu. Segala bentuk kecurangan (mencontek ataupun diskusi) akan menyebabkan skor live code ini 0.
3. Waktu pengerjaan: **120 menit** untuk **2 soal**.
4. **Pada text editor hanya ada file yang terdapat pada repository ini**.
5. Membuka referensi eksternal seperti Google, StackOverflow, dan MDN diperbolehkan.
6. Dilarang membuka repository di organisasi tugas, baik pada organisasi batch sendiri ataupun batch lain, baik branch sendiri maupun branch orang
lain (**setelah melakukan clone, close tab GitHub pada web browser kalian**).
7. Lakukan `git push origin <livecode-1-github_name>` **hanya jika waktu Live Code telah dinyatakan usai.
8. **Penilaian berbasis logika dan hasil akhir**. Pastikan keduanya sudah benar.


## KETENTUAN

Here are some dos and don'ts to consider when working on this livecode:

Dos:

- Do read and understand the problem statement and requirements carefully before starting to code.
- Do break down the problem into smaller, manageable parts and tackle each one individually.
- Do test your code frequently and thoroughly to ensure that it is functioning as intended.
- Do use good programming practices, such as meaningful variable names, clear comments, and properly formatted code.
- Do ask for help if you get stuck or need clarification on a specific concept or feature.


Don'ts:

- Don't rush through the problem or try to solve it all at once.
- Don't copy and paste code from external sources without fully understanding how it works.
- Don't hardcode values or rely on assumptions that may not hold true in all cases.
- Don't forget to handle error cases or edge cases, such as invalid input or unexpected behavior.
- Don't hesitate to refactor your code or make improvements based on feedback or new insights.


SET 2
# NUMBER 1 LIVE CODE 1

## **Pharmacy Booking System**

## Restrictions
- Don't use global variables unnecessarily. Instead, use local variables and pass them as parameters to functions as needed.
- Don't forget to handle errors returned by functions, using either an if err != nil statement or a panic statement where appropriate.
- Don't forget to use the correct type assertions when working with interfaces.
- Incorrect type assertions can result in runtime errors and unexpected behavior.

## Objectives
- Mampu menyelesaikan masalah yang diberikan.
- Mengerti konsep dan cara penggunaan logic dengan `conditional`, `looping` dan `data primitive`.

## Directions
As a software engineer at a pharmacy, you are tasked with designing a prescription and billing system that allows customers to select and order multiple pharmaceutical items. The project will be developed in two releases, focusing on the integration of conditional statements and looping concepts.


Services Menu : 

![image](https://github.com/H8-FTGO-AOH-P1/FTGO-AOH-P1-V2-LC1/blob/main/menu.png) 


Requirements builds upon the previous two releases and requires the program to handle more complex scenarios and logic. Here are some possible requirements for the 2 releases:

Release 1 : 

- **Add support for ordering multiple items**: The program should allow the user to order multiple items at once, and display the total cost of the order.

Release 1 Expectation : 

![image](https://github.com/H8-FTGO-AOH-P1/FTGO-AOH-P1-V2-LC1/blob/main/Release1LC.png) 


Release 2 : 

- **Add support for discounts and promotions**: The program should allow the user to apply discounts or promotions to their order based on `orderTotal > 1,000,000`, such as a percentage discount on the total cost. Additionally, a `tax rate of 0.08%` should be applied to the final amount. 

The formula for the final price should be:

`Final Price =(Original Price−Discount)+Tax`


- **Add logic for additional promotion**: If `orderTotal > 2,000,000`, add a promotion for a free "Antacids".


Release 2 Expectation : 

![image](https://github.com/H8-FTGO-AOH-P1/FTGO-AOH-P1-V2-LC1/blob/main/Release2LC.png) 

To fulfill these requirements, the program will need to implement more control structures such as nested loops, multiple conditionals, and switch statements. The code will also need to handle more complex data structures, such as arrays or maps, to store and manipulate the order data.

## Handling Incorrect Type/Error Behavior:

**The main requirement** is that if an input error is validated by the incorrect type/error handling, the program should not terminate. Instead, it should display an error message and prompt the user to re-enter the correct input.

A. Empty Input:

`No input provided. Please enter the item and the number of units.`

B. Invalid Input Format | Negative Number | Zero for Hours:

`Invalid number of units. Please enter a positive number of units in numeric format.`

C. Excessive Number of Units:

`Number of units too large. Please enter a reasonable number of units (e.g., less than 100).`


D. Invalid Item Format:

`Item not available. Please select a valid item from the menu.`


**Example Implementation of Error Handling**:

```go
// Example for handling errors without terminating the program

func main() {
    for {
        // Display items to the customer

        // Take customer's order input

        // Validate input
        if invalidInput {
            fmt.Println("Invalid input. Please try again.")
            continue
        }

        // If valid, process order
        break
    }

    // Display total to the customer
}

```

**expected Error Handling**:

![image](https://github.com/H8-FTGO-AOH-P1/FTGO-AOH-P1-V2-LC1/blob/main/1ErrHandling.png) 

![image](https://github.com/H8-FTGO-AOH-P1/FTGO-AOH-P1-V2-LC1/blob/main/2ErrHandling.png) 

![image](https://github.com/H8-FTGO-AOH-P1/FTGO-AOH-P1-V2-LC1/blob/main/3ErrHandling.png) 

## Notes

HINTS : Packages that can be used: (Optional)

`bufio`
`errors`
`fmt`
`os`
`strconv`
`strings`

Release 1 : 


- If **ITEM1** is added, display the **Total-Cost-Order** for **ITEM1** plus **ITEM2** etc


```go

	// Initialize items and prices

    // Display items to the customer

    // Take customer's order

    // Display total to the customer
	
```

Release 2 : 

HINT : 

- Calculate the total cost of the order.
- Apply a **10% discount if orderTotal > Rp 1,000,000.**
- Apply a **0.08% tax on the total cost after the discount.**
- Add a promotion for a **free "Antacids" if orderTotal > Rp 2,000,000.**

- If **orderTotal > Rp 1,000,000, a 10% discount is applied**, **a TaxRate of 0.08% is implemented**, and **a free "Antacids" is given if orderTotal > Rp 2,000,000.**


```go

	// Initialize items and prices

    // Display items to the customer

    // Take customer's order

    // Display total to the customer

    // Implement Discount, Tax, and Additional Free Item Logic
		
```
