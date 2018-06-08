package fakes

type FakeSSHD struct{}

func NewFakeSSHD() *FakeSSHD {
	return &FakeSSHD{}
}

func (*FakeSSHD) StartListen() error {
	return nil
}

func (*FakeSSHD) StopListen() error {
	return nil
}
