package main

var HttpStatus = map[int]string{
	200: "StatusOK",
	300: "StatusMultipleChoices",
	301: "StatusMovedPermanently",
	400: "StatusBadRequest",
	401: "StatusUnauthorized",
	500: "StatusInternalServerError",
	501: "StatusNotImplemented",
}
