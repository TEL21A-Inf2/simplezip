package dict

// Datentyp für die Knoten unseres Wörterbuchs.
type Dictionary struct {
	value       string
	left, right *Dictionary
}

// Erzeugt einen neuen leeren Knoten.
func NewDictionary() *Dictionary {
	return &Dictionary{"", nil, nil}
}

// Setzt den Wert zur gegebenen Sequenz. Steigt dabei rekursiv in den Baum ab, bis
func (node *Dictionary) SetValue(sequence, value string) {
	// Wenn die Sequenz leer ist, können wir hier den Wert setzen.
	if sequence == "" {
		node.value = value
		return
	}

	head, tail := sequence[:1], sequence[1:]

	// Wenn der Anfang der Sequenz ein "a" ist, dann müssen wir im linken Kind
	// den Wert eintragen. Falls das linke Kind noch nicht existert, wird es erstellt.
	if head == "a" {
		if node.left == nil {
			node.left = NewDictionary()
		}
		node.left.SetValue(tail, value)
		return
	}

	// Wenn wir hierher kommen, ist head == "b" und die Sequenz muss noch weitergehen.
	// Wir rufen SetValue() rekursiv mit tail auf dem rechten Kind auf.
	// Falls das rechte Kind noch nicht existiert, wird es vorher noch erstellt.
	if node.right == nil {
		node.right = NewDictionary()
	}
	node.right.SetValue(tail, value)
}

// Konsumiert den Text, bis ein Leerzeichen gelesen wird oder der Text leer ist.
// Folgt dabei den Kanten im Baum.
// Liefert den Wert des Knotens, der auf diese Weise gefunden wird.
// Liefert außerdem den restlichen Text.
func GetNextValue(node Dictionary, text string) (string, string) {
	// TODO
	return "value", "rest of text"
}
