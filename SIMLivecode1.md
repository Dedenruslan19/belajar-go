# Sim Live-Code-1-Phase-1
## RULES
1. **Untuk kampus remote**: **WAJIB** melakukan **share screen**(**DESKTOP/ENTIRE SCREEN**) dan **unmute microphone** ketika Live Code
berjalan (tidak melakukan share screen/salah screen atau tidak unmute microphone akan di ingatkan).
2. Kerjakan secara individu. Segala bentuk kecurangan (mencontek ataupun diskusi) akan menyebabkan skor live code ini 0.
3. Waktu pengerjaan: **90 menit** untuk **2 soal**.
4. **Pada text editor hanya ada file yang terdapat pada repository ini**.
5. Membuka referensi eksternal seperti Google, StackOverflow, dan MDN diperbolehkan.
6. Dilarang membuka repository di organisasi tugas, baik pada organisasi batch sendiri ataupun batch lain, baik branch sendiri maupun branch orang
lain (**setelah melakukan clone, close tab GitHub pada web browser kalian**).
7. Lakukan `git push origin <simulasi-livecode-1-github_name>` **hanya jika waktu Live Code telah dinyatakan usai.
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


SET 1
# NUMBER 1 LIVE CODE 1

## **Hacktiv8 Order System**

## Restrictions
- Don't use global variables unnecessarily. Instead, use local variables and pass them as parameters to functions as needed.
- Don't forget to handle errors returned by functions, using either an if err != nil statement or a panic statement where appropriate.
- Don't forget to use the correct type assertions when working with interfaces.
- Incorrect type assertions can result in runtime errors and unexpected behavior.


## Objectives
- Mampu menyelesaikan masalah yang diberikan.
- Mengerti konsep dan cara penggunaan logic dengan `conditional`, `looping` dan `data primitive`.

## Directions
You are a software engineer working on a point-of-sale system for a restaurant. You need to implement three features related to conditional and looping.

Hacktiv8 Restaurant Menu : 

![image](https://github.com/H8-FTGO-AOH-P1/FTGO-AOH-P1-V1-SLC1/blob/main/menu.png) 


The Requirements requires the program to handle more complex scenarios and logic. Here are some possible requirements for the 2 release:

Release 1 : 

- **Add support for ordering multiple items**: The program should allow the user to order multiple items at once, and display the total cost of the order.

Release 2 : 

- **Add support for discounts and promotions**: The program should allow the user to apply a `10% discount` to their order `if orderTotal > 10.00`. An example formula for calculating the final price after the discount: `Final Price=Original Price−Discount` WHERE `Discount = 10% of the Original Price`.

To fulfill these requirements, the program will need to implement more control structures such as nested loops, multiple conditionals, and switch statements. The code will also need to handle more complex data structures, such as arrays or maps, to store and manipulate the order data.


## Notes

Release 1 : 

![image](https://github.com/H8-FTGO-AOH-P1/FTGO-AOH-P1-V1-SLC1/blob/main/Release1SLC1.png) 


- Jika *ITEM2* ditambahkan maka tampilkan *Total-Cost-Order* untuk *ITEM1* ditambah *ITEM2*


```go

	// Initialize menu and prices
	
	// Display menu to the customer

	// Take customer's order
	
```

Release 2 : 

![image](https://github.com/H8-FTGO-AOH-P1/FTGO-AOH-P1-V1-SLC1/blob/main/Release2SLC1.png) 

- Implementasikan sebuah fitur order dimana menampilkan item yang dipilih oleh customer dari menu dan hitung total cost termasuk discount sebesar 10%.



```go

	// Initialize menu and prices
	
	// Display menu to the customer

	// Take customer's order
	
	// Display total to the customer
	
	// Implement Discount Logic
	
	
```
