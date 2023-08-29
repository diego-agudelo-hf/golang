package models

type Woman struct {
	Man
	Esmadre bool
}

func (h *Woman) Respirar() { h.respirando = true }
func (h *Woman) Pensar()   { h.pensando = true }
func (h *Woman) Comer()    { h.comiendo = true }

func (h *Woman) Gender() string { return "I'm women" }
