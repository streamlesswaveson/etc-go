package stack

import "testing"

func TestStackPush(t *testing.T)  {

	s := mystack{}
	s.push("grinch")

	if len(s.data) != 1 {
		t.Errorf("Expected length 1, got %v", len(s.data))
	}

}

func TestStackPeek(t *testing.T)  {
	s := mystack{}
	s.push("grinch")

	result  := s.peek()
	if result != "grinch" {
		t.Errorf("Expected the grinch got %v", result)
	}
}

func TestStackPeekWithTwo(t *testing.T)  {
	s := mystack{}
	s.push("grinch")
	s.push("Cindy Lou")

	result  := s.peek()
	if result != "Cindy Lou" {
		t.Errorf("Expected the Cindy Lou got %v", result)
	}
}

func TestStackPop(t *testing.T)  {
	s := mystack{}
	s.push("grinch")
	s.push("Cindy Lou")

	result := s.pop()
	if result != "Cindy Lou" {
		t.Errorf("Expected the Cindy Lou got %v", result)
	}

	result = s.pop()
	if result != "grinch" {
		t.Errorf("Expected the grinch got %v", result)
	}

	result = s.pop()
	if result != "" {
		t.Errorf("expected empty string got %v", result)
	}

}