package valueobject

import "time"

// Metadata default Newton metadata of an aggregate; lets developers soft-remove, restore and hard-remove an specific
// aggregate
//	This value object is using enriched built-in data types, apart from being a helper object.
//	Hence, validation and encapsulation is not required.
type Metadata struct {
	CreateTime     time.Time
	UpdateTime     time.Time
	State          bool
	MarkAsMutation bool
	MarkAsRemoval  bool
}
