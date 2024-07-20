package rubic

type StandardSpec struct{}

func (_ *StandardSpec) getDimensions() (FACES int, EDGES int, LAYERS int) {
	return 6, 4, 1
}

func (_ *StandardSpec) Populate(rubic Rubic) {
	for i := 0; i < rubic.FacesPerRubic; i++ {
		populateRFace(rubic.Faces[i], rubic)
	}
}

func (_ *StandardSpec) getRotations() [][]int {
	return [][]int{
		{0, 0, 0, 0},
		{1, 0, 3, 0},
		{2, 0, 2, 0},
		{3, 0, 1, 0},
		{2, 3, 0, 1},
		{0, 1, 2, 3},
	}
}

func (_ *StandardSpec) getConnections() [][]int {
	return [][]int{
		{4, 1, 5, 3},
		{4, 2, 5, 0},
		{4, 3, 5, 1},
		{4, 0, 5, 2},
		{2, 1, 0, 3},
		{0, 1, 2, 3},
	}
}

func populateRFace(face RFace, rubic Rubic) {

}
