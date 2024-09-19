package main

type MinHeap struct {
	data []int
}

func (m *MinHeap) Len() int {
	return len(m.data)
}

func (m *MinHeap) Delete() int {
	if m.Len() == 0 {
		return -1
	}

	val := m.data[0]
	if m.Len() == 1 {
		m.data = []int{}
		return val
	}
	m.data[0] = m.data[m.Len()-1]
	m.data = m.data[:m.Len()-1]
	m.heapifyDown(0)
	return val
}

func (m *MinHeap) Insert(val int) {
	m.data = append(m.data, val)
	idx := m.Len() - 1
	m.heapifyUp(idx)
}

func (m *MinHeap) heapifyDown(idx int) {
	if idx >= m.Len() {
		return
	}

	lIdx := m.leftChild(idx)
	rIdx := m.rightChild(idx)

	if lIdx >= m.Len() {
		return
	}

	smallest := idx
	if lIdx < m.Len() && m.data[lIdx] < m.data[smallest] {
		smallest = lIdx
	}
	if rIdx < m.Len() && m.data[rIdx] < m.data[smallest] {
		smallest = rIdx
	}
	if smallest != idx {
		m.swap(idx, smallest)
		m.heapifyDown(smallest)
	}
}

func (m *MinHeap) heapifyUp(idx int) {
	if idx <= 0 {
		return
	}

	parentIdx := m.getParentIdx(idx)

	if m.data[parentIdx] < m.data[idx] {
		return
	}

	if parentIdx >= 0 {
		m.swap(parentIdx, idx)
	}
	m.heapifyUp(parentIdx)
}

func (m *MinHeap) swap(i, j int) {
	m.data[i], m.data[j] = m.data[j], m.data[i]
}

func (m *MinHeap) leftChild(idx int) int {
	return idx*2 + 1
}

func (m *MinHeap) rightChild(idx int) int {
	return idx*2 + 2
}

func (m *MinHeap) getParentIdx(i int) int {
	return (i - 1) / 2
}
