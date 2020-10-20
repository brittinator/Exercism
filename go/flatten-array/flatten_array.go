package flatten

// Flatten returns a single flattened list with all values.
func Flatten(input interface{}) []interface{} {
	flattened := []interface{}{}
	flattenNested(input, &flattened)

	// fmt.Println("done ", flattened)
	return flattened
}

func flattenNested(input interface{}, flattened *[]interface{}) {
	// fmt.Printf("input %v\n", input)
	inputAsSlice, ok := input.([]interface{})
	if !ok {
		return
	}

	for _, item := range inputAsSlice {
		num, ok := item.(int)
		if ok {
			// fmt.Println("adding ", num)
			*flattened = append(*flattened, num)
		} else {
			// fmt.Println("going deeper...")
			flattenNested(item, flattened)
		}
	}
}
