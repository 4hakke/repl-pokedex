package networkclient

type NetworkClientInterface interface {
	Get(url string, resultedObject *any) error
}
