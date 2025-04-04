package service

type APIClient interface {
    FetchData(id string) (string, error)
}

type DataService struct {
    client APIClient
}

func NewDataService(client APIClient) *DataService {
    return &DataService{client: client}
}

func (s *DataService) GetProcessedData(id string) (string, error) {
    data, err := s.client.FetchData(id)
    if err != nil {
        return "", err
    }
    return "Processed: " + data, nil
}
