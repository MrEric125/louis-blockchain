package main

type Adidas struct {

}

func (adidas *Adidas) makeShoe() IShoe  {
	return &AdidasShoe{
		Shoe:Shoe{
			logo:"adidas shoe",
			size:15,
		},
	}
}

func (adidas *Adidas) makeShirt() IShirt  {
	return &AdidasShirt{
		Shirt:Shirt{
			logo: "adidas shirt",
			size: 18,
		},
	}
}

