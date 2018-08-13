package main

import (
	"fmt"
	"strings"
)

// ExpertConnection is a handle to ask questions
type ExpertConnection struct {
	questionChannel chan []int
	answerChannel   chan []string
}

// NewExpertConnection pipe questions to function
func NewExpertConnection(f func([]int) []string) (expert *ExpertConnection) {
	expert = &ExpertConnection{}
	expert.InitExpertConnection(f)
	return
}

// InitExpertConnection launches
func (expert *ExpertConnection) InitExpertConnection(f func([]int) []string) {
	expert.questionChannel = make(chan []int, 1)
	expert.answerChannel = make(chan []string, 1)
	go func() {
		for {
			q, more := <-expert.questionChannel
			if !more {
				fmt.Println("Thank you")
				return
			}
			a := f(q)
			expert.answerChannel <- a
		}
	}()
}

// Ask asks one question
func (expert *ExpertConnection) Ask(q []int) []string {
	fmt.Println("Sending question")
	expert.questionChannel <- q
	fmt.Println("Awaiting answer")
	a := <-expert.answerChannel
	fmt.Println("Returning answer")
	return a
}

// Close the connection
func (expert *ExpertConnection) Close() {
	fmt.Println("Hanging up")
	close(expert.questionChannel)
}

// NewAdder creates one handle
func NewAdder() *ExpertConnection {
	return NewExpertConnection(func(q []int) []string {
		sum := 0
		for _, v := range q {
			sum += v
		}
		s := fmt.Sprintf("Sum is %d", sum)
		return []string{s}
	})
}

// NewMultiplier creates one handle
func NewMultiplier() *ExpertConnection {
	return NewExpertConnection(func(q []int) []string {
		product := 1
		for _, v := range q {
			product *= v
		}
		s := fmt.Sprintf("Product is %d", product)
		return []string{s}
	})
}

// NewOracle creates one handle
func NewOracle(experts ...*ExpertConnection) *ExpertConnection {
	return NewExpertConnection(func(q []int) (answer []string) {
		answer = make([]string, 0)
		for _, e := range experts {
			a := e.Ask(q)
			answer = append(answer, a...)
		}
		return
	})
}

// interface Expert {
// 	Ask(q []int) []string
// }

func main() {
	adder := NewAdder()
	defer adder.Close()
	oracle := NewOracle(NewAdder(), NewMultiplier())
	defer oracle.Close()
	q := []int{1, 2, 3}
	a := adder.Ask(q)
	s := strings.Join(a, ", ")
	fmt.Println("Ask: ", q, " -> ", s)
	q = []int{1, 2, 3, 4}
	a = oracle.Ask(q)
	s = strings.Join(a, ", ")
	fmt.Println("Ask: ", q, " -> ", s)

}
