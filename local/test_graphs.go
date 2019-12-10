package local

import "github.com/suzuito/bacon-number-go/entity"

// https://www.cs.rit.edu/~ark/winter2012/730/team/1/report.pdf
var Nodes01 = NodeStoreImpl{
	Nodes: map[entity.NodeID]*entity.Node{
		"01": &entity.Node{
			ID: "01",
			Adjacencies: []entity.NodeID{
				"02",
				"04",
			},
		},
		"02": &entity.Node{
			ID: "02",
			Adjacencies: []entity.NodeID{
				"01",
				"03",
				"05",
			},
		},
		"03": &entity.Node{
			ID: "03",
			Adjacencies: []entity.NodeID{
				"02",
				"05",
			},
		},
		"04": &entity.Node{
			ID: "04",
			Adjacencies: []entity.NodeID{
				"01",
				"05",
			},
		},
		"05": &entity.Node{
			ID: "05",
			Adjacencies: []entity.NodeID{
				"02",
				"03",
				"04",
			},
		},
	},
}

// https://www.geeksforgeeks.org/dijkstras-shortest-path-algorithm-greedy-algo-7/
var Nodes02 = NodeStoreImpl{
	Nodes: map[entity.NodeID]*entity.Node{
		"00": &entity.Node{
			ID:          "00",
			Adjacencies: []entity.NodeID{"01", "07"},
		},
		"01": &entity.Node{
			ID:          "01",
			Adjacencies: []entity.NodeID{"00", "07", "02"},
		},
		"02": &entity.Node{
			ID:          "02",
			Adjacencies: []entity.NodeID{"01", "08", "05", "03"},
		},
		"03": &entity.Node{
			ID:          "03",
			Adjacencies: []entity.NodeID{"02", "05", "04"},
		},
		"04": &entity.Node{
			ID:          "04",
			Adjacencies: []entity.NodeID{"03", "05"},
		},
		"05": &entity.Node{
			ID:          "05",
			Adjacencies: []entity.NodeID{"02", "03", "04", "06"},
		},
		"06": &entity.Node{
			ID:          "06",
			Adjacencies: []entity.NodeID{"05", "07", "08"},
		},
		"07": &entity.Node{
			ID:          "07",
			Adjacencies: []entity.NodeID{"00", "01", "06", "08"},
		},
		"08": &entity.Node{
			ID:          "08",
			Adjacencies: []entity.NodeID{"02", "06", "07"},
		},
	},
}

var Nodes03 = NodeStoreImpl{
	Nodes: map[entity.NodeID]*entity.Node{
		"00": &entity.Node{
			ID:          "00",
			Adjacencies: []entity.NodeID{"02", "01"},
		},
		"01": &entity.Node{
			ID:          "01",
			Adjacencies: []entity.NodeID{"00", "02"},
		},
		"02": &entity.Node{
			ID:          "02",
			Adjacencies: []entity.NodeID{"00", "01"},
		},
	},
}
