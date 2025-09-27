package main

import (
	"sort"
	"sync"
	"time"
)

// Struct untuk Task
type Task struct {
	Name     string
	Priority int
}

func main() {
	//Define isi (Name, Priority) untuk struct Task dengan slice
	tasks := []Task{
		{"Emergency Medical Transport", 1},
		{"Delivery at Zone A", 2},
		{"Pickup at Zone B", 3},
	}

	// Urutkan task berdasarkan priority secara ASC
	// Define anon func dengan bool untuk membandingkan priority berdasarkan index
	sort.Slice(tasks, func(i, j int) bool {
		// jika index i lebih kecil dari j maka balikkan nilai sebagai true
		return tasks[i].Priority < tasks[j].Priority
	})

	// Define buffer channel untuk ditangkap oleh driver/worker
	taskChan := make(chan Task, len(tasks))
	var wg sync.WaitGroup

	// Perulangan untuk goroutines driver yang siap untuk menerima task
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go driver(i, taskChan, &wg)
	}

	// kirim task sesuai priority ke channel dan memberi jeda selama 50 * time.Millisecond
	for _, t := range tasks {
		taskChan <- t
		time.Sleep(50 * time.Millisecond)
	}
	// Menutup taskChan bahwa tidak ada lagi tugas yang dapat dikerjakan
	close(taskChan)

	wg.Wait()
}

// func driver untuk menerima task dari taskChan dan simulasi
func driver(id int, taskChan <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range taskChan {
		// Simulate processing the task
		println("Driver", id, "processing:", task.Name)
	}
}
