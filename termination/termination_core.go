package termination

import (
	"sync"
	"woohoo/stream"
)

type Termination struct {

}

type ValveTerm interface {
	Fire(in interface{}, output *chan interface{},wg *sync.WaitGroup)
}

func (term *Termination) FireTerm(stream stream.Stream) {
	//开一个携程触发了就不管了

	stream.DoFireUp()

	output := []interface{}{}

	out := stream.output

	for {
		if data, ok := <-*out; ok {
			output = append(output, data)
		} else {
			break
		}
	}

	return output

}

