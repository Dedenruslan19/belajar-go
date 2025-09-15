# SIM Live-Code-2-Phase-1
## RULES
1. **Untuk kampus remote**: **WAJIB** melakukan **share screen**(**DESKTOP/ENTIRE SCREEN**) dan **unmute microphone** ketika Live Code
berjalan (tidak melakukan share screen/salah screen atau tidak unmute microphone akan di ingatkan).
2. Kerjakan secara individu. Segala bentuk kecurangan (mencontek ataupun diskusi) akan menyebabkan skor live code ini 0.
3. Clone repo dari URL yang diberikan oleh instruktur.
4. Kerjakan dengan file Go (\*.go).
5. `Pastikan solusi(point 4) untuk setiap release disimpan dalam folder terpisah` yang dinamai **Release1** dan **Release2**.
6. Waktu pengerjaan: **90 menit** untuk 2 release.
7. **Pada text editor hanya ada file yang terdapat pada repository ini**.
8. Membuka referensi eksternal seperti Google, StackOverflow, dan MDN diperbolehkan.
9. Dilarang membuka repository di organisasi tugas, baik pada organisasi batch sendiri ataupun batch lain, baik branch sendiri maupun branch orang lain (**setelah melakukan clone, close tab GitHub pada web browser kalian**).
10. Lakukan `git push origin <simlivecode-2-github_name>` **hanya jika waktu Live Code telah dinyatakan usai.
11. **Penilaian berbasis logika dan hasil akhir terkait Konkurensi**. Pastikan keduanya sudah benar.


## KETENTUAN
Here are some dos and don'ts to consider when working on this Sim livecode:

Do`s:

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

## Objectives
- Mampu menyelesaikan masalah yang diberikan.
- Memahami konsep dan cara penggunaan `looping`.
- Mengerti konsep concurrent
- Mengerti konsep `sync.WaitGroup` 
- Mengerti konsep dan cara penggunaan `go routine` dan `channel`.


SET 1
# NUMBER 1 SIM LIVE CODE 2
## Release 1 : Refactoring for an Online Ride-Hailing Application

### Directions :

You've received a piece of software from a vendor that's intended for use in your online ride-hailing application project. This software handles task assignments to drivers but suffers from race conditions, leading to unpredictable and inconsistent behaviors. This could impact the reliability and performance of your application, especially under high load. Your tech lead has tasked you with refactoring this code to ensure it operates reliably and sequentially, even under concurrent execution.

**Vendor Program Description**:

The vendor's program is designed to manage a pool of drivers and assign them tasks as they come in. It uses multiple goroutines to handle different tasks simultaneously to simulate real-time task handling. However, the lack of proper synchronization mechanisms and the presence of shared state management have introduced race conditions, resulting in tasks sometimes being assigned incorrectly or even skipped.

Please check on directory : `skeleton-code` on this repo's github.

## Instructions 

**Do**:

- Use `sync.WaitGroup` to ensure all tasks complete before the program exits.
- Use channels to control task executions and ensure tasks are executed in a sequential order.
- Provide descriptive variable names and extensive comments to clarify each function's role and goroutine operations.

**Don't**:

- Use global variables to handle state management between goroutines.
- Employ unnecessary **time.Sleep** or other delays that do not simulate real-world conditions.
- Ignore error handling or let the program crash due to unhandled exceptions.


## Expected Output - Release 1 :

![image](https://github.com/H8-FTGO-AOH-P1/FTGO-AOH-P1-V1-SLC2/blob/main/Release1.png)


## Release 2 : Refactoring for an Online Ride-Hailing Application

### Directions :

With the initial race conditions resolved, your next assignment is to enhance the system by introducing dynamic task allocation based on the driver's status and a priority scheduling system.

## Instructions 

**Do**:

- Implement a dynamic task allocation system where a special goroutine monitors the statusDriver channel and assigns tasks only when drivers are marked as not busy.
- Establish a priority scheduling system using a buffered channel, ensuring high-priority tasks are addressed first.
- Maintain clear and concise documentation and comments throughout your code to explain the new functionalities and their interactions.

**Don't**:

- Overlook the necessity of using sync.WaitGroup to synchronize the completion of all tasks.
- Hardcode critical values like the number of drivers or tasks, which may need to scale with different scenarios.
- Neglect thorough testing of the new features under various conditions to ensure robustness and reliability.

## Expected Output 

![image](https://github.com/H8-FTGO-AOH-P1/FTGO-AOH-P1-V1-SLC2/blob/main/Release2.png)


**Disclaimer** : 

- `High-priority tasks` such as `"Emergency Medical Transport"` are often processed first, showing the effectiveness of the priority system.

- `taskChannel as a buffered channel` allows tasks to be processed as they become available, reflecting the dynamic nature of concurrency.

- `"All tasks have completed."` consistently appears after all tasks have been processed, ensuring all tasks are managed properly.


### HINT (Pseudocode):

## release 1
```go
// Initialize WaitGroup and any necessary channels
// Define goroutines for each task, ensuring proper synchronization with WaitGroup
// Inside each goroutine, use channels to synchronize and control task execution order
// Wait for all goroutines to finish using WaitGroup.Wait()
// Main function should start all goroutines and ensure they complete before exiting
```

## release 2
```go
// Create a buffered channel for tasks with priority handling
// Implement a goroutine to monitor driver status and assign tasks from the channel
// Each driver runs in a goroutine, pulling tasks from the channel based on availability
// Use buffered channel to ensure high-priority tasks are executed first
// Synchronize the completion of all tasks and drivers using WaitGroup
// Ensure final output indicates completion of all tasks
```
