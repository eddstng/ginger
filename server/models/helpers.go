package models

func PtrInt(v int) *int             { return &v }
func PtrFloat64(v float64) *float64 { return &v }
func PtrString(v string) *string    { return &v }
func PtrBool(v bool) *bool          { return &v }
