package main

import "github.com/suzuito/bacon-number-go/entity"

// https://www.cs.rit.edu/~ark/winter2012/730/team/1/report.pdf
var nodes01 = NodeStoreImpl{
	nodes: map[entity.NodeID]entity.Node{
		"01": entity.Node{
			ID: "01",
			Adjacencies: []entity.NodeID{
				"02",
				"04",
			},
		},
		"02": entity.Node{
			ID: "02",
			Adjacencies: []entity.NodeID{
				"01",
				"03",
				"05",
			},
		},
		"03": entity.Node{
			ID: "03",
			Adjacencies: []entity.NodeID{
				"02",
				"05",
			},
		},
		"04": entity.Node{
			ID: "04",
			Adjacencies: []entity.NodeID{
				"01",
				"05",
			},
		},
		"05": entity.Node{
			ID: "05",
			Adjacencies: []entity.NodeID{
				"02",
				"03",
				"04",
			},
		},
	},
}
