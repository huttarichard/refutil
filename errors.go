package refutil

import "errors"

// ErrArgumentNotIndexable is telling you that argument supplied to function
// is not a searchable for specific index
var ErrArgumentNotIndexable = errors.New("argument can't be used to search index")
