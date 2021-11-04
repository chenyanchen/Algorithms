// Copyright 2020 Singularity, Inc. All rights reserved.

/*
 * Revision History:
 *     Initial: 2018/05/16        Chen Yanchen
 */

package main

type Graph struct {
	edgNum int // the number of edge
	vexNum int // the number of vertex
	adj    []Vertex
}

type Vertex struct {
	Date []byte
	e    *Edge
}

type Edge struct {
	ivex int
	next *Edge
}

// BFS Breadth-first search
func BFS() {

}
