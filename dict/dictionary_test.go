package dict

import "fmt"

func ExampleDictionary() {
	root := NewDictionary()

	root.SetValue("a", "wörter")
	root.SetValue("ba", "werden")
	root.SetValue("bba", "häufigsten")
	root.SetValue("bbba", "wenn")
	root.SetValue("bbbba", "die")
	root.SetValue("bbbbba", "durch")
	root.SetValue("bbbbbba", "kürzere")

	fmt.Println(root.value)
	fmt.Println(root.left.value)
	fmt.Println(root.right.left.value)
	fmt.Println(root.right.right.left.value)
	fmt.Println(root.right.right.right.left.value)
	fmt.Println(root.right.right.right.right.left.value)
	fmt.Println(root.right.right.right.right.right.left.value)
	fmt.Println(root.right.right.right.right.right.right.left.value)

	// Output:
	// wörter
	// werden
	// häufigsten
	// wenn
	// die
	// durch
	// kürzere
}
