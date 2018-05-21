package _struct

type Human struct {

	length int64
	height int64
	width int64

	id int64
	name string
	relationship map[string]*Human
}

func NewHuman() *Human {
	var h *Human

	return h
}

func (h *Human) Length() int64 {
	return h.length
}

func (h *Human) Height() int64 {
	return h.height
}

func (h *Human) Width() int64 {
	return h.width
}

func (h *Human) Response() {

}

func (h *Human) Effect()  {

}
