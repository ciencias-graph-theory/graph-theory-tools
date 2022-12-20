package doublylexordering

// Auxiliary structure to define an entry in the map.
type Entry struct {
	rowKey *IntSet
	colKey *IntSet
	value  *Block
}

// newEntry initializes a new entry.
func newEntry(Ri, Cj *IntSet, B *Block) *Entry {
	return &Entry{
		rowKey: Ri,
		colKey: Cj,
		value:  B,
	}
}

// matches returns if a key matches with entry's keys. In this case, an entry
// matches with a key if and only if the rowKeys and colKeys are the same.
func (e *Entry) matches(Ri, Cj *IntSet) bool {
	return (e.rowKey.Equals(Ri) && e.colKey.Equals(Cj))
}

// hashBlock is the hash function used to determine where to place the block in
// the map. The current method returns the sum of the values of both sets; this
// may yield some collisions.
func hashBlock(Ri, Cj *IntSet) int {
	h := 0
	for _, val := range Ri.GetValues() {
		h += val
	}

	for _, val := range Cj.GetValues() {
		h += val
	}

	return h
}

// BlockMap is a map where the keys are a pair of integer sets Ri and Cj, and
// the value is the block B = (Ri, Cj). Usually the value should have some
// attributes calculated already, i.e. its size and the row blocks sizes too.
type BlockMap struct {
	entries map[int][]*Entry
}

// NewBlockMap initializes a block map.
func NewBlockMap() *BlockMap {
	return &BlockMap{
		entries: make(map[int][]*Entry),
	}
}

// Add adds an entry to the map.
func (bm *BlockMap) Add(Ri, Cj *IntSet, B *Block) {
	// Define an entry with the given parameters.
	x := newEntry(Ri, Cj, B)

	// Hash value of the block.
	h := hashBlock(Ri, Cj)

	// There already exists a block with same hash value.
	if _, ok := bm.entries[h]; ok {
		// Traverse all of the entries that share said hash value.
		for _, e := range bm.entries[h] {
			// If some entry matches, return.
			if e.matches(Ri, Cj) {
				return
			}
		}

		// Otherwise, there is no record of the entry in the map; add it to the
		// slice.
		bm.entries[h] = append(bm.entries[h], x)
	} else {
		// It's the first block with said hash value.
		var s []*Entry
		bm.entries[h] = append(s, x)
	}
}

// Get returns the entry that matches the given key. If no entry matches the key
// return nil.
func (bm *BlockMap) Get(Ri, Cj *IntSet) *Block {
	// Hash value of the block.
	h := hashBlock(Ri, Cj)

	// There already exists a block with same hash value.
	if _, ok := bm.entries[h]; ok {
		// Traverse all of the entries that share said hash value.
		for _, e := range bm.entries[h] {
			// If some entry matches, return the entry.
			if e.matches(Ri, Cj) {
				return e.value
			}
		}
	}

	// If no entry matches the key, return nil.
	return nil
}

// Contains returns whether the block map contains the given key.
func (bm *BlockMap) Contains(Ri, Cj *IntSet) bool {
	// Hash value of the block.
	h := hashBlock(Ri, Cj)

	// There already exists a block with same hash value.
	if _, ok := bm.entries[h]; ok {
		// Traverse all of the entries that share said hash value.
		for _, e := range bm.entries[h] {
			// If some entry matches, return the entry.
			if e.matches(Ri, Cj) {
				return true
			}
		}
	}

	// If no entry matches the key, return nil.
	return false
}
