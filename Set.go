package main

import ("fmt")


const Max_Hash_Table_Size = 256

type node struct {
	value int;
	next *node;
}

type set struct {
	hash_table [Max_Hash_Table_Size]*node;
	count int;
}

func create() *set {
	return nil//&set{make([]*node, Max_Hash_Table_Size), 0};
}

func hash(v int) int {
	return (v + 1) % Max_Hash_Table_Size
}

func add(s *set, x int) {
	var hash_val int = hash(x);
	for n := s.hash_table[hash_val]; n != nil; n = n.next {
		if n.value == x {
			return
		}
	}
	var n *node = &node{x, s.hash_table[hash_val]};
	s.hash_table[hash_val] = n;
	s.count++;
}

func contains(s *set, x int) bool {
	fmt.Println(s.hash_table)
	idx := hash(x)
	fmt.Println("Using index {}", idx)
	return (s.hash_table[idx] != nil)
}

func size(s *set) int {
	return s.count
}