package utils

type mockClient struct {
	response   []byte
	statusCode int
	err         error
}

func NewApiMockClient(response []byte, statusCode int, err error) Client {
	return &mockClient{response, statusCode, err}
}

func (c *mockClient) GetDataFromUrl(url string) ([]byte, int, error) {
	return c.response, c.statusCode, c.err
}
