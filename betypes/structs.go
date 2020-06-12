package betypes

//Point - We can save the location of the object with X and Y (For example, the location of the clan)
type Point struct {
	X, Y int
}

//Mine - It will give us resources depending on the level
type Mine struct {
	Location Point //Each mine has its location on the map (X and Y)
	Level    int   //The speed of resource extraction depends on it
	ToBelong int   //The number of the clan that owns the mine
}
