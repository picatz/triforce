package triforce

type Normal struct {
	I float32
	J float32
	K float32
}

func (n *Normal) ZeroOut() {
	n.I = 0
	n.J = 0
	n.K = 0
}
