package models

type Man struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
}

func (h *Man) Respirar()      { h.respirando = true }
func (h *Man) Pensar()        { h.pensando = true }
func (h *Man) Comer()         { h.comiendo = true }
func (h *Man) Gender() string { return "I'm Men" }
