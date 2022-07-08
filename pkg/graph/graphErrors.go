package graph

type graphError struct {
	message string
}

func GraphError(message string) error {
	return graphError{
		message: message,
	}
}

func (error graphError) Error() string {
	return error.message
}

var (
	assymetricMatrixError = GraphError("Adjacency matrix is not symmetric")
	invalidListError      = GraphError("Invalid adjacency list")
)
