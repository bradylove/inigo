package fake_bbs

import "github.com/cloudfoundry-incubator/runtime-schema/models"

type FakeLRPRouterBBS struct {
	DesiredLRPChangeChan chan models.DesiredLRPChange
	desiredLRPStopChan   chan bool
	desiredLRPErrChan    chan error

	ActualLRPChangeChan chan models.ActualLRPChange
	actualLRPStopChan   chan bool
	actualLRPErrChan    chan error

	AllDesiredLRPs []models.DesiredLRP
	AllActualLRPs  []models.ActualLRP

	DesiredLRP models.DesiredLRP
	ActualLRPs []models.ActualLRP

	WhenGettingRunningActualLRPs func() ([]models.ActualLRP, error)
	WhenGettingAllDesiredLRPs    func() ([]models.DesiredLRP, error)

	WhenGettingActualLRPsByProcessGuid func(string) ([]models.ActualLRP, error)
	WhenGettingDesiredLRPByProcessGuid func(string) (models.DesiredLRP, error)
}

func NewFakeLRPRouterBBS() *FakeLRPRouterBBS {
	return &FakeLRPRouterBBS{
		DesiredLRPChangeChan: make(chan models.DesiredLRPChange, 1),
		desiredLRPStopChan:   make(chan bool),
		desiredLRPErrChan:    make(chan error),

		ActualLRPChangeChan: make(chan models.ActualLRPChange, 1),
		actualLRPStopChan:   make(chan bool),
		actualLRPErrChan:    make(chan error),
	}
}

func (fakeBBS *FakeLRPRouterBBS) WatchForDesiredLRPChanges() (<-chan models.DesiredLRPChange, chan<- bool, <-chan error) {
	return fakeBBS.DesiredLRPChangeChan, fakeBBS.desiredLRPStopChan, fakeBBS.desiredLRPErrChan
}

func (fakeBBS *FakeLRPRouterBBS) WatchForActualLRPChanges() (<-chan models.ActualLRPChange, chan<- bool, <-chan error) {
	return fakeBBS.ActualLRPChangeChan, fakeBBS.actualLRPStopChan, fakeBBS.actualLRPErrChan
}

func (fakeBBS *FakeLRPRouterBBS) GetAllDesiredLRPs() ([]models.DesiredLRP, error) {
	if fakeBBS.WhenGettingAllDesiredLRPs != nil {
		return fakeBBS.WhenGettingAllDesiredLRPs()
	}

	return fakeBBS.AllDesiredLRPs, nil
}

func (fakeBBS *FakeLRPRouterBBS) GetRunningActualLRPs() ([]models.ActualLRP, error) {
	if fakeBBS.WhenGettingRunningActualLRPs != nil {
		return fakeBBS.WhenGettingRunningActualLRPs()
	}

	return fakeBBS.AllActualLRPs, nil
}

func (fakeBBS *FakeLRPRouterBBS) GetDesiredLRPByProcessGuid(processGuid string) (models.DesiredLRP, error) {
	if fakeBBS.WhenGettingDesiredLRPByProcessGuid != nil {
		return fakeBBS.WhenGettingDesiredLRPByProcessGuid(processGuid)
	}

	return fakeBBS.DesiredLRP, nil
}

func (fakeBBS *FakeLRPRouterBBS) GetRunningActualLRPsByProcessGuid(processGuid string) ([]models.ActualLRP, error) {
	if fakeBBS.WhenGettingActualLRPsByProcessGuid != nil {
		return fakeBBS.WhenGettingActualLRPsByProcessGuid(processGuid)
	}

	return fakeBBS.ActualLRPs, nil
}
