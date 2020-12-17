package functions



type ValvePassthrough struct{

}

func (valve *ValvePassthrough) Fire(in interface{},output *chan interface{}){


	*output<-in



}

func NewPassthrough() *ValvePassthrough{
	return new(ValvePassthrough)
}