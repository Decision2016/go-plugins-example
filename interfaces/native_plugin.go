package interfaces

type NativeConverter interface {
	Run(string) interface{}
}
