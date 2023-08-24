package ernos

type m struct {
	RECORD_NOTFOUND       string
	INTERNAL_SERVER_ERROR string
}

var M = m{
	RECORD_NOTFOUND:       "record not found",
	INTERNAL_SERVER_ERROR: "internal server error",
}
