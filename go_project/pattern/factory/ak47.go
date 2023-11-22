package main


type Ak47 struct {
	Gun
}

func newAk47() IGun  {
	//var gun=Gun{
	//	name:"Ak47 gun",
	//	power: 4,
	//}
	//return &Ak47{
	//	Gun:gun,
	//}
	//
	return &Ak47{
		Gun:Gun{
			name:"Ak47 gun",
			power: 4,
		},
	}
}