package functions

import "github.com/HeyJavaBean/woohoo/stream"

func (s *stream.Stream) Map(mapFunc MapFunc) *Stream {

	
	
	return s.AddStage( NewMap(mapFunc))

}

func (stream *Stream) Filter(filterFunc FilterFunc) *Stream {

	return stream.AddStage( NewFilter(filterFunc))

}

func (stream *Stream) Limit(limitNum int) *Stream {

	return stream.AddStage( NewLimit(limitNum))

}

func (stream *Stream) Skip(skipNum int) *Stream {

	return stream.AddStage( NewSkip(skipNum))

}

func (stream *Stream) Distinct() *Stream {

	return stream.AddStage( NewDistinct())

}

func (stream *Stream) SafePeek(peekF  PeekFunc) *Stream {

	return stream.AddStage( NewSafePeek(peekF))

}

func (stream *Stream) Sort(comparator  Comparator) *Stream {

	return stream.AddStatefulStage( NewSort(comparator))

}

func (stream *Stream) FlatMap(fmF  FlatMapFunc) *Stream {

	return stream.AddStage( NewFlatMap(fmF))

}

func (stream *Stream) Peek(peekF  PeekFunc) *Stream {

	return stream.AddStage(NewPeek(peekF))

}

