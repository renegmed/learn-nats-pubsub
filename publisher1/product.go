package main

import "strconv"

type Product struct {
	ProductID      int    `json:ProductID`
	Manufacturer   string `json:Manufacturer`
	Sku            string `json:Sku`
	Upc            string `json:Upc`
	PricePerUnit   string `json:PricePerUnit`
	QuantityOnHand int    `json:QuantityOnHand`
	ProductName    string `json:ProductName`
}

func (p *Product) ProductTable() map[string]Product {

	products := []Product{
		{
			ProductID:      1,
			Manufacturer:   "Johns-Jenkins",
			Sku:            "p5z343vdS",
			Upc:            "939581000000",
			PricePerUnit:   "497.45",
			QuantityOnHand: 9703,
			ProductName:    "sticky note",
		},
		{
			ProductID:      2,
			Manufacturer:   "Hessel, Schimmel and Feeney",
			Sku:            "i7v300kmx",
			Upc:            "740979000000",
			PricePerUnit:   "282.29",
			QuantityOnHand: 9217,
			ProductName:    "leg warmers",
		},
		{
			ProductID:      3,
			Manufacturer:   "Swaniawski, Bartoletti and Bruen",
			Sku:            "q0L657ys7",
			Upc:            "111730000000",
			PricePerUnit:   "436.26",
			QuantityOnHand: 5905,
			ProductName:    "lamp shade",
		},
		{
			ProductID:      4,
			Manufacturer:   "Runolfsdottir, Littel and Dicki",
			Sku:            "x78426lq1",
			Upc:            "93986215015",
			PricePerUnit:   "537.90",
			QuantityOnHand: 2642,
			ProductName:    "flowers",
		},
		{
			ProductID:      5,
			Manufacturer:   "Kuhn, Cronin and Spencer",
			Sku:            "r4X793mdR",
			Upc:            "260149000000",
			PricePerUnit:   "112.10",
			QuantityOnHand: 6144,
			ProductName:    "clamp",
		},
		{
			ProductID:      6,
			Manufacturer:   "Quigley, Casper and Boyer",
			Sku:            "g2P499xrM",
			Upc:            "390162000000",
			PricePerUnit:   "593.53",
			QuantityOnHand: 6507,
			ProductName:    "twister",
		},
		{
			ProductID:      7,
			Manufacturer:   "Gutmann and Sons",
			Sku:            "v4j250hbi",
			Upc:            "465922000000",
			PricePerUnit:   "88.97",
			QuantityOnHand: 4348,
			ProductName:    "clay pot",
		},
		{
			ProductID:      8,
			Manufacturer:   "Bins-Hansen",
			Sku:            "d3E278nk2",
			Upc:            "20072026056",
			PricePerUnit:   "933.35",
			QuantityOnHand: 4577,
			ProductName:    "tooth picks",
		},
		{
			ProductID:      9,
			Manufacturer:   "Jones, Braun and Ratke",
			Sku:            "o3w911oal",
			Upc:            "879638000000",
			PricePerUnit:   "426.23",
			QuantityOnHand: 1882,
			ProductName:    "mirror",
		},
		{
			ProductID:      10,
			Manufacturer:   "Upton-Mraz",
			Sku:            "k0R875prt",
			Upc:            "877387000000",
			PricePerUnit:   "630.61",
			QuantityOnHand: 4036,
			ProductName:    "rug",
		},
		{
			ProductID:      11,
			Manufacturer:   "Schneider, Douglas and Franecki",
			Sku:            "h3t822kaB",
			Upc:            "507019000000",
			PricePerUnit:   "13.67",
			QuantityOnHand: 2285,
			ProductName:    "headphones",
		},
		{
			ProductID:      12,
			Manufacturer:   "Weimann, Waelchi and Kassulke",
			Sku:            "t8A474iuv",
			Upc:            "669100000000",
			PricePerUnit:   "369.46",
			QuantityOnHand: 5409,
			ProductName:    "balloon",
		},
		{
			ProductID:      13,
			Manufacturer:   "Moore-Gibson",
			Sku:            "c4i321mvh",
			Upc:            "600535000000",
			PricePerUnit:   "250.98",
			QuantityOnHand: 6219,
			ProductName:    "lip gloss",
		},
		{
			ProductID:      14,
			Manufacturer:   "Schuppe, Cummerata and Hammes",
			Sku:            "l9V771xw1",
			Upc:            "677102000000",
			PricePerUnit:   "577.90",
			QuantityOnHand: 1104,
			ProductName:    "sidewalk",
		},
		{
			ProductID:      15,
			Manufacturer:   "Ward-Quigley",
			Sku:            "x7b275hk2",
			Upc:            "110459000000",
			PricePerUnit:   "642.19",
			QuantityOnHand: 9371,
			ProductName:    "pen",
		},
		{
			ProductID:      16,
			Manufacturer:   "Cassin Inc",
			Sku:            "b27590nys",
			Upc:            "920302000000",
			PricePerUnit:   "145.19",
			QuantityOnHand: 5382,
			ProductName:    "outlet",
		},
		{
			ProductID:      17,
			Manufacturer:   "Shanahan Inc",
			Sku:            "j4y570kml",
			Upc:            "850798000000",
			PricePerUnit:   "409.02",
			QuantityOnHand: 4567,
			ProductName:    "blanket",
		},
		{
			ProductID:      18,
			Manufacturer:   "Upton-Runolfsdottir",
			Sku:            "m2K116lkV",
			Upc:            "10335859487",
			PricePerUnit:   "815.26",
			QuantityOnHand: 7083,
			ProductName:    "lotion",
		},
		{
			ProductID:      19,
			Manufacturer:   "Lubowitz Group",
			Sku:            "w61375szc",
			Upc:            "286418000000",
			PricePerUnit:   "700.53",
			QuantityOnHand: 2224,
			ProductName:    "socks",
		},
		{
			ProductID:      20,
			Manufacturer:   "Padberg, Grady and Mueller",
			Sku:            "f1D653dwZ",
			Upc:            "62442197606",
			PricePerUnit:   "176.67",
			QuantityOnHand: 4544,
			ProductName:    "CD",
		},
	}

	prods := map[string]Product{}
	for _, p := range products {
		prods[strconv.Itoa(p.ProductID)] = p
	}

	return prods
}
