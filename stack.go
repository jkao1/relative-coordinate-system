package main

type struct Stack {
	[][][]float64
}

func NewCStack() *Stack {
	stack := new(Stack)
	m := NewMatrix()
	MakeIdentity(m)
	return &append(stack, m)
}

func (*s Stack) Push() {
	head := Copy2d(s[len(s) - 1])
	s := append(s, head)
}

func (*s Stack) Peek() *Stack {
	return &s[len(s) - 1]
}

func Copy2d(m [][]float64) [][]float64 {
	duplicate := make([][]float64, len(m))
	for i := range m {
    duplicate[i] = make([]float64, len(m[i]))
    copy(duplicate[i], m[i])
	}
	return duplicate
}
