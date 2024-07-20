package rubic

type RFCorner struct {
	TabCell [][]int
}

type RFSlice struct {
	VecCell []int
}

type RFace struct {
	Corner      RFCorner
	Slice       RFSlice
	Rotations   []int
	Connections []int
}

type Rubic struct {
	FacesPerRubic int
	EdgesPerFace  int
	LayersPerEdge int
	Faces         []RFace
}

type Specifier interface {
	populate(rubic Rubic)
	getDimensions() (FACES int, EDGES int, LAYERS int)
	getConnections() [][]int
	getRotations() [][]int
}

func makeRFSlice(layersPerEdge int) RFSlice {
	return RFSlice{
		VecCell: make([]int, layersPerEdge),
	}
}
func makeRFCorner(layersPerEdge int) RFCorner {
	tabCells := make([][]int, layersPerEdge)
	for j := 0; j < layersPerEdge; j++ {
		tabCells[j] = make([]int, layersPerEdge)
	}
	return RFCorner{
		TabCell: tabCells,
	}
}

func makeRFace(edgeCount int, layersPerEdge int) RFace {
	return RFace{
		Corner:    makeRFCorner(layersPerEdge),
		Slice:     makeRFSlice(layersPerEdge),
		Rotations: make([]int, 0),
	}
}

func Make(specifier Specifier) Rubic {
	FACES, EDGES, LAYERS := specifier.getDimensions()
	var faces = make([]RFace, FACES)
	for i := 0; i < FACES; i++ {
		faces[i] = makeRFace(EDGES, LAYERS)
	}
	rubic := Rubic{
		FacesPerRubic: FACES,
		EdgesPerFace:  EDGES,
		LayersPerEdge: LAYERS,
		Faces:         faces,
	}
	specifier.populate(rubic)
	return rubic
}
