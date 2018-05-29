package main

import (
	"project/Algorithm/graph/DAG/adjacency"
)

func main() {
	dag := adjacency.New()

	dag.Add("5", "2", "0")
	dag.Add("4", "0", "1")
	dag.Add("2", "3")
	dag.Add("3", "1")

	//
	dag.Show()

	//
	dag.Kahn()

	//
	//dag.Show()
}
