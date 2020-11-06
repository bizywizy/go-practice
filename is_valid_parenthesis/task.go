package is_valid_parenthesis

//Valid Parentheses
//Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//An input string is valid if:
//* Open brackets must be closed by the same type of brackets.
//* Open brackets must be closed in the correct order.
//https://leetcode.com/problems/valid-parentheses/

const (
	openedRound  = '('
	closedRound  = ')'
	openedCurve  = '{'
	closedCurve  = '}'
	openedSquare = '['
	closedSquare = ']'
)

type stack struct {
	array [10001]rune
	ptr   uint16
}

func (s *stack) Push(symbol rune) {
	s.ptr++
	s.array[s.ptr] = symbol
}

func (s *stack) Pop() rune {
	if s.Empty() {
		panic("pop from empty stack")
	}
	s.ptr--
	return s.array[s.ptr+1]
}

func (s *stack) Empty() bool {
	return s.ptr == 0
}

func IsValid(s string) bool {
	stack := &stack{}

	for _, char := range s {
		switch char {
		case openedRound:
			fallthrough
		case openedCurve:
			fallthrough
		case openedSquare:
			stack.Push(char)
		case closedRound:
			if stack.Empty() || (stack.Pop() != openedRound) {
				return false
			}
		case closedCurve:
			if stack.Empty() || (stack.Pop() != openedCurve) {
				return false
			}
		case closedSquare:
			if stack.Empty() || (stack.Pop() != openedSquare) {
				return false
			}

		}
	}
	return stack.Empty()
}
