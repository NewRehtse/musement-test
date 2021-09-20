package utils

type mockClient struct {
	response   []byte
	statusCode int
	err         error
}

// NewApiMockClient creates a mock client useful for test purpusoses
func NewApiMockClient(response []byte, statusCode int, err error) Client {
	return &mockClient{response, statusCode, err}
}

// GetDataFromUrl returns the set response so it can be tested
func (c *mockClient) GetDataFromUrl(url string) ([]byte, int, error) {
	return c.response, c.statusCode, c.err
}
