# Simulasi Live-Code-3-Phase-1

## RULES

1. **Untuk kampus remote**: **WAJIB** melakukan **share screen**(**DESKTOP/ENTIRE SCREEN**) dan **unmute microphone** ketika Live Code
berjalan (tidak melakukan share screen/salah screen atau tidak unmute microphone akan di ingatkan).
2. Kerjakan secara individu. Segala bentuk kecurangan (mencontek ataupun diskusi) akan menyebabkan skor live code ini 0.
3. Clone repo ini kemudian buatlah **branch dengan nama kalian**.
4. Kerjakan pada file go (\*.go) yang telah disediakan.
5. Waktu pengerjaan: **120 menit** untuk **2 soal**.
6. **Pada text editor hanya ada file yang terdapat pada repository ini**.
7. Membuka referensi eksternal seperti Google, StackOverflow, dan MDN diperbolehkan.
8. Dilarang membuka repository di organisasi tugas, baik pada organisasi batch sendiri ataupun batch lain, baik branch sendiri maupun branch orang
lain (**setelah melakukan clone, close tab GitHub pada web browser kalian**).
9. git push origin `<simlivecode-2-github_name>` **hanya jika waktu Live Code telah usai (bukan ketika kalian sudah selesai
mengerjakan)**.
10. **Penilaian berbasis logika dan hasil akhir**. Pastikan keduanya sudah benar.


### KETENTUAN

Here are some dos and don'ts to consider when working on this livecode:

Do's:

- Clearly define the problem statement and requirements for the challenge.
- Use realistic and representative data for the challenge.
- Provide a database schema or ER diagram for the challenge.
- Specify the SQL dialect and version to be used for the challenge.
- Provide a clear submission format and deadline for the challenge.
- Evaluate the challenge solutions based on clear and transparent criteria.
- Provide feedback and follow-up for the candidates after evaluating their solutions.

Don'ts:

- Use overly complex or unrealistic scenarios for the challenge.
- Include ambiguous or poorly defined requirements for the challenge.
- Use inconsistent or illogical data for the challenge.
- Assume the candidates have access to proprietary or confidential information for the challenge.
- Ignore the importance of data security and privacy in the challenge.
- Use vague or subjective evaluation criteria for the challenge.
- Provide no or insufficient feedback to the candidates after evaluating their solutions.

SET 1

## NUMBER 1 SIM LIVE CODE 3

## Mini Soccer Field Booking App Development Challenge

### Restrictions

- Don't use global variables unnecessarily. Instead, use local variables and pass them as parameters to functions as needed.
- Don't forget to handle errors returned by functions, using either an if err != nil statement or a panic statement where appropriate.
- Don't use fmt.Println excessively for debugging purposes. Instead, use Go's built-in testing framework or a separate debugging tool.
- Don't forget to use the correct type assertions when working with interfaces.
- Incorrect type assertions can result in runtime errors and unexpected behavior.
- Don't use unidiomatic or non-idiomatic Go code, such as using while loops instead of for loops, or using := instead of var for variable declarations when appropriate.

## Objectives

- Mampu menyelesaikan masalah yang diberikan.
- Memahami konsep `DDL` dan `DML`
- Memahami konsep `Query SQL`
- Memahami penggunaan logika kondisi
- Memahami penggunaan logika perulangan
- Memahami penggunaan konkurensi dengan goroutines, channel hingga wg
- Memahami penggunaan Golang Basic CLI

## Directions

### Title: Mini Soccer Field Booking App Development Backend Story

After years of dazzling the world with his football skills, Cristiano Ronaldo has decided to retire and delve into the world of business. Passionate about the game of soccer and always eager to give back to the community, Ronaldo has set his eyes on promoting mini soccer in Indonesia. He envisions an app where locals can easily book mini soccer fields for their games, gatherings, or practice sessions. But before that can become a reality, Ronaldo needs your help to lay the groundwork for the database that will support this app!

### RELEASE 1 - Database Analyse

Objective: Examine the given sample database for the MiniSoccer Field Booking App. Identify any inefficiencies, redundancies, or anomalies.

Here an overview database required :

##### Fields Table

![image](https://github.com/H8-FTGO-AOH-P1/FTGO-AOH-P1-V1-SLC3/blob/main/Dummy1.png)

##### Customers Table

![image](https://github.com/H8-FTGO-AOH-P1/FTGO-AOH-P1-V1-SLC3/blob/main/Dummy2.png)

##### Bookings Table

![image](https://github.com/H8-FTGO-AOH-P1/FTGO-AOH-P1-V1-SLC3/blob/main/Dummy3.png)

##### Payments Table

![image](https://github.com/H8-FTGO-AOH-P1/FTGO-AOH-P1-V1-SLC3/blob/main/Dummy4.png)

#### To-Do :

a. Review the table structures, relationships, and data.

b. Identify which parts of the database require normalization. Explain why and recommend the normalization level (1NF, 2NF, 3NF, etc.).

c. Submit a report detailing your findings and the reasoning behind your normalization choices.

### Design an Entity Relationship Diagram (ERD):

Objective: Create a visual representation of the relationships between different entities in the MiniSoccer Field Booking App's database.

#### To-Do:
a. List down all the entities involved, their attributes, and relationships.

b. Design the ERD, showcasing primary and foreign keys, cardinalities (one-to-one, one-to-many, many-to-many), and any other necessary details.

c. Use an ERD design tool of your choice. Ensure your design is clear and easy to understand.

### SQL Representation:

Objective: Convert your ERD into actual SQL table creations.

#### To-Do:

a. Write SQL queries to create the tables, establish relationships, and set up the constraints as identified in your ERD.

b. Ensure the SQL syntax is correct, and the relationships are accurately represented

c. Submit the SQL script containing all your table creation queries.

### Your task, as the backend developer, was to analyze these requirements and develop an Entity Relationship Diagram (ERD) that depicts the interconnections between these entities. 

### Please Notes

You were aware that this ERD was a vital step in revamping the system, ensuring the principles of database normalization were upheld to minimize data redundancy and maintain data integrity. Understanding this, you dived headfirst into the task, ready to tackle another challenge and support Ronaldo's growth journey.

After finalizing the ERD, you would then be tasked with translating these relationships into a database structure in XAMPP, creating SQL queries to perform CRUD operations, designing views for easy access to information, and developing advanced SQL queries for valuable business insights.

Hint **Release 1** :

- **Do Analyze ERD** based on requirement on Field Requirements above
- **Create your ERD** based on your analyze on `point a`
- **Provide your SQL Query** to create db on localhost

### Expected Output

- Analyze ERD - Found an file called `Template_analisa.md` on this repos.

## sample

![image](https://github.com/H8-FTGO-AOH-P1/FTGO-AOH-P1-V1-SLC3/blob/main/Release1-AnalyzeERD.png)

- ERD
  
## sample

![image](https://github.com/H8-FTGO-AOH-P1/FTGO-AOH-P1-V1-SLC3/blob/main/Release1B-ERD.png)

- SQL Query

## sample

![image](https://github.com/H8-FTGO-AOH-P1/FTGO-AOH-P1-V1-SLC3/blob/main/Release1C-Query.png)

Good luck, and may you help Ronaldo realize his vision for promoting mini soccer in Indonesia!

### Your Sample Data

Before starting release 2 please insert this sample data into your database :

```sql

-- Inserting data into Fields table
INSERT INTO Fields (FieldName, FieldSize, Location, HourlyRate) VALUES 
('Ronaldo Arena', '5v5', 'Jakarta', 50.00),
('Messi Park', '7v7', 'Bali', 70.00),
('Neymar Ground', '5v5', 'Surabaya', 45.00);

-- Inserting data into Customers table
INSERT INTO Customers (FirstName, LastName, Email, PhoneNumber) VALUES 
('Ahmad', 'Setiawan', 'ahmad.set@example.com', '+628123456789'),
('Siti', 'Fatimah', 'siti.fat@example.com', '+628987654321'),
('Budi', 'Harsono', 'budi.h@example.com', '+628456789012');

-- Inserting data into Bookings table
INSERT INTO Bookings (CustomerID, FieldID, BookingDate, StartTime, EndTime) VALUES 
(1, 2, '2023-08-10', '16:00:00', '18:00:00'),
(3, 1, '2023-08-11', '10:00:00', '12:00:00'),
(2, 3, '2023-08-12', '14:00:00', '15:00:00');

-- Inserting data into Payments table
INSERT INTO Payments (BookingID, PaymentDate, PaymentAmount, PaymentMethod) VALUES 
(1, '2023-08-09', 140.00, 'Credit Card'),
(2, '2023-08-10', 100.00, 'Bank Transfer'),
(3, '2023-08-11', 45.00, 'Cash');

```
### RELEASE 2 - Mini Report

Cristiano Ronaldo, having transitioned from a football legend to a successful entrepreneur in Indonesia, is keen on optimizing the operations of his Mini Soccer Field Booking App. He relies heavily on data to make informed decisions and is particularly interested in two main insights. First, understanding which soccer field is generating the most revenue and attracting more bookings is vital. It allows him to evaluate the popularity of fields such as 'Messi Park', 'Ronaldo Arena', and 'Neymar Ground' and guides his investment strategy. Second, he wishes to keep a close eye on customers who've made field bookings but are lagging in their payments. Quick follow-ups and ensuring that all booked slots are paid for are crucial for the business's profitability.

Please make sure to test the reports functionality thoroughly to ensure accurate and reliable data retrieval.

#### Objectives

Develop a Command Line Interface (CLI) program using Go (Golang) to help Ronaldo pull these insights efficiently. The program should connect to the database, retrieve, and display these insights seamlessly.

Requirements:

A. Revenue Report:

The CLI should provide an option to fetch and display the total revenue generated by each field.
The report should display the number of bookings and total revenue, sorted in descending order of revenue.
Fields should be easily identifiable by their names like 'Messi Park', 'Ronaldo Arena', etc.

B. Pending Payments Report:

Another option in the CLI should generate a list of customers who have made bookings but haven't completed their payments.
This report should display the customer name, booking ID, and the booking date.

Hint **Release 2** :

- **Maximize** SQL Query and display on *CLI Program*

Expected Output :

![image](https://github.com/H8-FTGO-AOH-P1/FTGO-AOH-P1-V1-SLC3/blob/main/R2Dummy1.png)

![image](https://github.com/H8-FTGO-AOH-P1/FTGO-AOH-P1-V1-SLC3/blob/main/R2Dummy2.png)

![image](https://github.com/H8-FTGO-AOH-P1/FTGO-AOH-P1-V1-SLC3/blob/main/R2Dummy3.png)
