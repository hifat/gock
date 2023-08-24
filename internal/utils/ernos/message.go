package ernos

type m struct {
	RECORD_NOT_FOUND      string
	INTERNAL_SERVER_ERROR string
}

var M = m{
	RECORD_NOT_FOUND:      "record not found",
	INTERNAL_SERVER_ERROR: "internal server error",
}
