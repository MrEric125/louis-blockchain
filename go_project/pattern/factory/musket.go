package main

type musket struct {
	Gun
}

func newMusket() IGun {
	return &musket{
		Gun:Gun{
			name: "musket gun",
			power: 1,
		},
	}

}
