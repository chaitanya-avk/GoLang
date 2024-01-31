package main

import (
	"fmt"
	"slices"
)

var Product_family = map[string][]string{
	"7750":   []string{"SR-7", "SR-12", "SR-7s"},
	"7450":   []string{"ess-7", "ess-6", "ess-12"},
	"7250":   []string{"IXR-R4", "IXR-e", "IXR-R6dl"},
	"7705hm": []string{"SAR-HM", "SAR-H"},
}

func main() {

	fmt.Println("Initialized values of the map")
	printMap()

	//Adding new family to the map
	Product_family["7950"] = []string{"XRS"}
	fmt.Println("After adding a new family to the map")
	printMap()

	//Adding node to the existing family
	addNode("7750", "SR-2s")
	addNode("750", "XRS")
	addNode("7950", "XRS")
	fmt.Println("Updated map after adding the node to a existing family")
	printMap()

	deleteDeprecatedNode("7750", "SR-7")
	deleteDeprecatedNode("7950", "XRS")
	fmt.Println("Updated map after deleting the deprecated node from a existing family")
	printMap()

}

func deleteDeprecatedNode(f, depChassis string) {

	_, boolean_exists := Product_family[f]
	if boolean_exists {
		values := Product_family[f]
		exists := slices.Contains(values, depChassis)
		if exists {
			for i, _ := range values {

				//Deleting the family if only one node is present which needs to be deprecated
				if len(values) == 1 {
					delete(Product_family, f)
				} else if values[i] == depChassis {
					Product_family[f] = append(values[:i], values[i+1:]...)
				}
			}
		} else {
			fmt.Println("Node is not present in the family")
		}
	} else {
		fmt.Println("Family does not exist in the map")
	}
}

func printMap() {
	for fam, chassis := range Product_family {
		fmt.Println(fam, chassis)
	}
}

func addNode(f string, newChassis string) {

	_, boolean_exists := Product_family[f]
	if boolean_exists {
		values := Product_family[f]
		exists := slices.Contains(values, newChassis)
		if !exists {
			Product_family[f] = append(Product_family[f], newChassis)
		} else {
			fmt.Println("Node is already present in the family")
		}
	} else {
		fmt.Println("Family does not exist in the map")
	}
}
