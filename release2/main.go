package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	Name string
}

func Driver(driverID int, task Task, wg *sync.WaitGroup, runID int) {
	defer wg.Done()

	time.Sleep(time.Duration(driverID) * 100 * time.Millisecond)

	fmt.Printf("Run %d:\n", runID)
	fmt.Printf("Driver %d: Processing Task: %s\n", driverID, task.Name)
}

func main() {
	tasks := []Task{
		{Name: "Emergency Medical Transport"},
		{Name: "Delivery at Zone A"},
		{Name: "Pickup at Zone B"},
	}

	simulateRun(1,
		[]Assignment{
			{DriverID: 2, Task: tasks[0]},
			{DriverID: 1, Task: tasks[1]}, 
			{DriverID: 2, Task: tasks[2]}, 
		})


	simulateRun(2,
		[]Assignment{
			{DriverID: 1, Task: tasks[0]}, 
			{DriverID: 2, Task: tasks[1]}, 
			{DriverID: 1, Task: tasks[2]},
		})

	simulateRun(3,
		[]Assignment{
			{DriverID: 2, Task: tasks[0]},
			{DriverID: 1, Task: tasks[1]},
			{DriverID: 1, Task: tasks[2]}, 
		})

}


type Assignment struct {
	DriverID int
	Task     Task
}

func simulateRun(runID int, assignments []Assignment) {
	var wg sync.WaitGroup
	

	fmt.Printf("\nRun %d:\n", runID)

	for _, assignment := range assignments {
		wg.Add(1)
		go func(a Assignment) {
			fmt.Printf("Driver %d: Processing Task: %s\n", a.DriverID, a.Task.Name)
			time.Sleep(20 * time.Millisecond)
			wg.Done()
		}(assignment)
	}

	wg.Wait() 

	fmt.Println("All tasks have completed")
}