package main


func persist() {
      insert(buffer)
      buffer = buffer[:0]
}

func main{
buffer := []kafkaMsg{}
bufferSize := 100
timeout := 100 * time.Millisecond

for {
  select {
    case kafkaMsg := <-channel:
      buffer = append(buffer, kafkaMsg)
      if len(buffer) >= bufferSize {
        persist()
      }
    case<-time.After(timeout):
      persist()
  }
}
}

