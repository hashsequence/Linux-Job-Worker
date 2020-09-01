package dataStore

import (
	//"errors"
	"fmt"
	"testing"
	assert "github.com/stretchr/testify/assert"
	guuid "github.com/google/uuid"
	"os"
)


func TestNewDataStore(t *testing.T) {
	fmt.Println("Running TestNewDataStore Test")
	dataStore, err := NewDataStore()
	assert.NotEqual(t, dataStore, nil, "dataStore should not be nil")
	fmt.Println(dataStore.logFolder)
	_, err = os.Stat(dataStore.logFolder) 
	fmt.Println(err)
	assert.Equal(t, os.IsNotExist(err), false, "log folder should have been created")
	assert.Equal(t, err, nil, "error should be nil")
	//os.RemoveAll(dataStore.logFolder)

}

func TestAdd(t *testing.T) {
	fmt.Println("Running TestAdd Test")
	dataStore, _ := NewDataStore()
	assert.Equal(t, dataStore.Add(guuid.New().String(),"ls","-al"), nil, "dataStore.Add() should return nil")
	fmt.Println(dataStore)
	//os.RemoveAll(dataStore.logFolder)
}


func TestUpdatePid(t *testing.T) {
	fmt.Println("Running TestUpdatePid Test")
	dataStore,_ := NewDataStore()
	uuid := guuid.New().String()
	assert.Equal(t, dataStore.Add(uuid,"ls","-al"), nil)
	//make a fake pid
	pid := 1345
	dataStore.UpdateWithPid(uuid,pid, nil)
	assert.Equal(t,dataStore.ProcessTable[uuid].pid,1345, "pid should match the one assigned")
	//os.RemoveAll(dataStore.logFolder)
}


func TestUpdateFinishProcess(t *testing.T) {
	fmt.Println("Running TestUpdateFinishProcessTest")
	dataStore,_ := NewDataStore()
	uuid := guuid.New().String()
	assert.Equal(t, dataStore.Add(uuid,"ls","-al"), nil, "dataStore.Add() should return nil")
	//make a fake pid
	pid := 235
	dataStore.UpdateWithPid(uuid,pid,nil)
//	assert.Equal(t,dataStore[uuid].pid,1346, "pid should match the one assigned")
	dataStore.UpdateFinishProcess(uuid, 123)
	assert.Equal(t,dataStore.ProcessTable[uuid].exitCode,123)
	//os.RemoveAll(dataStore.logFolder)
}

func TestUpdateFailedProcessDidNotStart(t *testing.T) {
	fmt.Println("Running TestUpdateFinishProcessTest")
	dataStore,_ := NewDataStore()
	uuid := guuid.New().String()
	assert.Equal(t, dataStore.Add(uuid,"ls","-al"), nil, "dataStore.Add() should return nil")
	dataStore.UpdateFailedProcessDidNotStart(uuid)
	assert.Equal(t,len(dataStore.ProcessTable[uuid].endTimeStamp),29)
	os.RemoveAll(dataStore.logFolder)
}
