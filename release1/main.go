package main

import (
	"fmt"
	"sort"
	"sync"
)

// Driver merepresentasikan seorang driver
type Driver struct {
	ID int
}

// Task dengan prioritas (semakin kecil = semakin penting)
type Task struct {
	Name     string
	Priority int
}

// Fungsi driver untuk menangani task
func (d *Driver) work(taskChannel <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range taskChannel {
		fmt.Printf("Driver %d is handling task: %s\n", d.ID, task.Name)
	}
}

func main() {
	var wg sync.WaitGroup

	// Daftar driver
	drivers := []*Driver{
		{ID: 1},
		{ID: 2},
		{ID: 3},
	}

	// Daftar task dengan prioritas
	tasks := []Task{
		{"Pick up passenger A", 2},
		{"Emergency Medical Transport", 1}, // high priority
		{"Deliver package C", 3},
		{"Pick up passenger D", 2},
		{"Pick up passenger E", 2},
	}

	// Urutkan task berdasarkan prioritas
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Priority < tasks[j].Priority
	})

	// Buat buffered channel untuk task
	taskChannel := make(chan Task, len(tasks))

	// Masukkan semua task ke channel
	for _, task := range tasks {
		taskChannel <- task
	}
	close(taskChannel)

	// Jalankan driver
	for _, driver := range drivers {
		wg.Add(1)
		go driver.work(taskChannel, &wg)
	}

	// Tunggu semua task selesai
	wg.Wait()
	fmt.Println("All tasks have completed.")
}
