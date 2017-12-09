package stack


type mystack struct {
	data []string
}

func (m *mystack) push(s string)  {
	m.data = append(m.data, s);
}

func (m *mystack) peek() string  {
	if len(m.data) > 0 {
		return m.data[len(m.data)-1]
	}
	return ""
}

func (m *mystack) pop() string  {
	if (len(m.data) > 0){
		result := m.data[len(m.data)-1]
		m.data = m.data[0:len(m.data)-1]
		return result
	}
	return ""

}