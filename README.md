# csv-worker
simple, concurrent csv worker in go


## Example
```go
func main() {

	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	worker1 := NewCsvWorker("test1.csv")
	worker2 := NewCsvWorker("test2.csv")
	worker3 := NewCsvWorker("test3.csv")
	worker4 := NewCsvWorker("test4.csv")

	for i := 1; i < 10; i++ {
		worker1.Recieve(records)
		worker2.Recieve(records)
		worker3.Recieve(records)
		worker4.Recieve(records)
	}

	worker1.Close()
	worker2.Close()
	worker3.Close()
	worker4.Close()
}
```