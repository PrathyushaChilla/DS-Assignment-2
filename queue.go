program1

package main
import "fmt"

func enqueue(queue[] int, element int) []int {
  queue = append(queue, element); 
  fmt.Println("Enqueued:", element);
  return queue
}

func dequeue(queue[] int) ([]int) {
  element := queue[0]; 
  fmt.Println("Dequeued:", element)
  return queue[1:]; 
}

func main() {
  var queue[] int; 

  queue = enqueue(queue, 10);
  queue = enqueue(queue, 20);
  queue = enqueue(queue, 30);

  fmt.Println("Queue:", queue);

  queue = dequeue(queue);
  fmt.Println("Queue:", queue);

  queue = enqueue(queue, 40);
  fmt.Println("Queue:", queue);
}
output:
Enqueued: 10
Enqueued: 20
Enqueued: 30
Queue: [10 20 30]
Dequeued: 10
Queue: [20 30]
Enqueued: 40
Queue: [20 30 40]

Program exited.







