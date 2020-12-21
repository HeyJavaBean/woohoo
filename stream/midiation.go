package stream

import "woohoo/functions"

func (stream *Stream) Map(mapFunc functions.MapFunc) *Stream {

	return stream.AddStage(functions.NewMap(mapFunc))

}

func (stream *Stream) Filter(filterFunc functions.FilterFunc) *Stream {

	return stream.AddStage(functions.NewFilter(filterFunc))

}

func (stream *Stream) Limit(limitNum int) *Stream {

	return stream.AddStage(functions.NewLimit(limitNum))

}

func (stream *Stream) Skip(skipNum int) *Stream {

	return stream.AddStage(functions.NewSkip(skipNum))

}

func (stream *Stream) Distinct() *Stream {

	return stream.AddStage(functions.NewDistinct())

}

func (stream *Stream) SafePeek(peekF functions.PeekFunc) *Stream {

	return stream.AddStage(functions.NewSafePeek(peekF))

}

func (stream *Stream) Sort(comparator functions.Comparator) *Stream {

	return stream.AddStatefulStage(functions.NewSort(comparator))

}

func (stream *Stream) FlatMap(fmF functions.FlatMapFunc) *Stream {

	return stream.AddStage(functions.NewFlatMap(fmF))

}

func (stream *Stream) Peek(peekF functions.PeekFunc) *Stream {

	return stream.AddStage(functions.NewPeek(peekF))

}

