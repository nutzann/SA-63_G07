// Code generated by entc, DO NOT EDIT.

package level

const (
	// Label holds the string label denoting the level type in the database.
	Label = "level"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"

	// EdgeArea holds the string denoting the area edge name in mutations.
	EdgeArea = "area"

	// Table holds the table name of the level in the database.
	Table = "levels"
	// AreaTable is the table the holds the area relation/edge.
	AreaTable = "areas"
	// AreaInverseTable is the table name for the Area entity.
	// It exists in this package in order to avoid circular dependency with the "area" package.
	AreaInverseTable = "areas"
	// AreaColumn is the table column denoting the area relation/edge.
	AreaColumn = "level_area"
)

// Columns holds all SQL columns for level fields.
var Columns = []string{
	FieldID,
	FieldName,
}
