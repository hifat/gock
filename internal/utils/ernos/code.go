package ernos

type c struct {
	RECORD_NOT_FOUND      string
	INTERNAL_SERVER_ERROR string
}

var C = c{
	RECORD_NOT_FOUND:      "RECORD_NOT_FOUND",
	INTERNAL_SERVER_ERROR: "INTERNAL_SERVER_ERROR",
}
