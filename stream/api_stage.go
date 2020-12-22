package stream

import "github.com/HeyJavaBean/woohoo/stage"

func (s *Stream) Map(mapFunc stage.MapFunc)  *Stream {

	return s.AddStage(true,stage.NewMap(mapFunc))

}

func (s *Stream) Filter(filterFunc stage.FilterFunc)  *Stream {

	return s.AddStage(true,stage.NewFilter(filterFunc))

}

func (s *Stream) Limit(limitNum int)  *Stream {

	return s.AddStage(true, stage.NewLimit(limitNum))

}

func (s *Stream) Skip(skipNum int)  *Stream {

	return s.AddStage( true,stage.NewSkip(skipNum))

}

func (s *Stream) Meter(count stage.MeterCounter)  *Stream {

	return s.AddStage( true,stage.NewMeter(count))

}

func (s *Stream) Distinct()  *Stream {

	return s.AddStage(true,stage.NewDistinct())

}

func (s *Stream) SafePeek(peekF stage.PeekFunc)  *Stream {

	return s.AddStage( true,stage.NewSafePeek(peekF))

}

func (s *Stream) Sort(comparator stage.Comparator)  *Stream {

	return s.AddStage(false,stage.NewSort(comparator))

}

//输出后的内容变成了Entry类型
func (s *Stream) ThenGroup(identifyFunc stage.IdentifyFunc)  *Stream {

	return s.AddStage(false,stage.NewThenGroup(identifyFunc))

}

//输出后的内容变成了EntryCount
func (s *Stream) ThenGroupCount(identifyFunc stage.IdentifyFunc)  *Stream {

	return s.AddStage(false,stage.NewThenGroupCount(identifyFunc))

}



func (s *Stream) FlatMap(fmF stage.FlatMapFunc)  *Stream {

	return s.AddStage(true, stage.NewFlatMap(fmF))

}

func (s *Stream) Peek(peekF stage.PeekFunc)  *Stream {

	return s.AddStage(true,stage.NewPeek(peekF))

}

