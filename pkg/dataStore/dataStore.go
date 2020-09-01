package dataStore

import(
	"sync"
	"fmt"
	"os"
	"path/filepath"
	utils "github.com/hashsequence/Copy-Linux-Job-Worker/pkg/utils"
	linuxJobWorkerPb "github.com/hashsequence/Copy-Linux-Job-Worker/pkg/pb"
	"os/exec"
)
/*Note***********************************************************
log paths are configured in Linux Format, so will not work on windows
*****************************************************************/

	type ProcessInfo struct {
		pid int
		startTimeStamp string
		endTimeStamp string
		processName string
		logPath string
		stdoutPath string
		stderrPath string
		isRunning bool
		exitCode int	
		cmd *exec.Cmd
	}

	type ProcessTable map[string]*ProcessInfo

	type DataStore struct {
		sync.RWMutex
		ProcessTable
		logFolder string
	}

func NewDataStore() (*DataStore, error) {
	path := "logs"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	} else if !os.IsExist(err) && err != nil {
		return nil, fmt.Errorf("Failed to create Log path")
	}

	absPath, _ := filepath.Abs(path)
	return &DataStore {
		ProcessTable : ProcessTable{},
		logFolder : absPath,
	}, nil
}

//getter functions for DataStore
func (this *DataStore) isExistUUID(uuid string) bool {
	if _ ,ok := this.ProcessTable[uuid]; ok {
		return true
	} 
	return false
}

func (this *DataStore) GetPid(uuid string) (int, bool)  {
	this.RLock()
	defer func() {
		this.RUnlock()
	}()
	if _ ,ok := this.ProcessTable[uuid]; !ok {
		return 0, false
	} 
	return this.ProcessTable[uuid].pid, true
}

func (this *DataStore) GetStartTimeStamp(uuid string) (string, bool) {
	this.RLock()
	defer func() {
		this.RUnlock()
	}()
	if _ ,ok := this.ProcessTable[uuid]; !ok {
		return "", false
	} 
	return this.ProcessTable[uuid].startTimeStamp, true
}

func (this *DataStore) GetEndTimeStamp(uuid string)  (string, bool) {
	this.RLock()
	defer func() {
		this.RUnlock()
	}()
	if _ ,ok := this.ProcessTable[uuid]; !ok {
		return "", false
	} 
	return this.ProcessTable[uuid].endTimeStamp, true
}

func (this *DataStore) GetProcessName(uuid string)  (string, bool) {
	this.RLock()
	defer func() {
		this.RUnlock()
	}()
	if _ ,ok := this.ProcessTable[uuid]; !ok {
		return "", false
	} 
	return this.ProcessTable[uuid].endTimeStamp, true
}

func (this *DataStore) GetLogPath(uuid string)  (string, bool) {
	this.RLock()
	defer func() {
		this.RUnlock()
	}()
	if _ ,ok := this.ProcessTable[uuid]; !ok {
		return "", false
	} 
	return this.ProcessTable[uuid].logPath, true
}

func (this *DataStore) GetStdoutPath(uuid string)  (string, bool) {
	this.RLock()
	defer func() {
		this.RUnlock()
	}()
	if _ ,ok := this.ProcessTable[uuid]; !ok {
		return "", false
	} 
	return this.ProcessTable[uuid].stdoutPath, true
}

func (this *DataStore) GetStderrPath(uuid string)  (string, bool) {
	this.RLock()
	defer func() {
		this.RUnlock()
	}()
	if _ ,ok := this.ProcessTable[uuid]; !ok {
		return "", false
	} 
	return this.ProcessTable[uuid].stderrPath, true
}

func (this *DataStore) GetIsRunning(uuid string)  (bool, bool) {
	this.RLock()
	defer func() {
		this.RUnlock()
	}()
	if _ ,ok := this.ProcessTable[uuid]; !ok {
		return false, false
	} 
	return this.ProcessTable[uuid].isRunning, true
}

func (this *DataStore) GetExitCode(uuid string)  (int, bool) {
	this.RLock()
	defer func() {
		this.RUnlock()
	}()
	if _ ,ok := this.ProcessTable[uuid]; !ok {
		return 0, false
	} 
	return this.ProcessTable[uuid].exitCode, true
}

func (this *DataStore) GetCmd(uuid string) (*exec.Cmd, bool) {
	this.RLock()
	defer func() {
		this.RUnlock()
	}()
	if _ ,ok := this.ProcessTable[uuid]; !ok {
		return nil, false
	} 
	return this.ProcessTable[uuid].cmd, true
}

func (this *DataStore) GetRespProcessInfo(uuid string) (*linuxJobWorkerPb.ProcessInfo, bool) {
	this.RLock()
	defer func() {
		this.RUnlock()
	}()
	if this.isExistUUID(uuid) {

		pid := this.ProcessTable[uuid].pid
		startTimeStamp := this.ProcessTable[uuid].startTimeStamp
		endTimeStamp := this.ProcessTable[uuid].endTimeStamp
		processName := this.ProcessTable[uuid].processName
		isRunning := this.ProcessTable[uuid].isRunning
		exitCode := this.ProcessTable[uuid].exitCode

		return &linuxJobWorkerPb.ProcessInfo{
			Pid: int32(pid),
			StartTimeStamp : startTimeStamp,
			EndTimeStamp : endTimeStamp,
			ProcessName : processName,
			IsRunning : isRunning,
			ExitCode : int32(exitCode),
		}, true
	}
	return nil, false
}

func (this *DataStore) getRespProcessInfo(uuid string) (*linuxJobWorkerPb.ProcessInfo, bool) {
	if this.isExistUUID(uuid) {

		pid := this.ProcessTable[uuid].pid
		startTimeStamp := this.ProcessTable[uuid].startTimeStamp
		endTimeStamp := this.ProcessTable[uuid].endTimeStamp
		processName := this.ProcessTable[uuid].processName
		isRunning := this.ProcessTable[uuid].isRunning
		exitCode := this.ProcessTable[uuid].exitCode

		return &linuxJobWorkerPb.ProcessInfo{
			Pid: int32(pid),
			StartTimeStamp : startTimeStamp,
			EndTimeStamp : endTimeStamp,
			ProcessName : processName,
			IsRunning : isRunning,
			ExitCode : int32(exitCode),
		}, true
	}
	return nil, false
}
//add the uuid for the incoming request into dataStore
//create the starting timestamp
//create log folder for uuid-startTimeStamp

func (this *DataStore) Add(uuid string, cmd ...string) error {
	this.Lock()
	defer func() {
		this.Unlock()
	}()
	
	processName := ""
	for _, cmdPart := range cmd {
		processName += " " + cmdPart
	}
	
	this.ProcessTable[uuid] = &ProcessInfo {
		startTimeStamp : utils.GetTimeStamp(),
		endTimeStamp : "",
		processName : processName,
		isRunning : false,
		exitCode : 0,
		cmd : nil,
	}
	if _, ok := this.ProcessTable[uuid]; !ok {
		return fmt.Errorf("Unable to Create ProcessInfo for %v\n", uuid)
	}
	//create log folder for uuid
	this.ProcessTable[uuid].logPath = this.logFolder + "/" + uuid + "-" + this.ProcessTable[uuid].startTimeStamp + "/"
	path := this.ProcessTable[uuid].logPath
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("create path " + path)
		os.Mkdir(path, 0755)
	} else if os.IsExist(err) {
		fmt.Println("remove and create path " + path)
		os.RemoveAll(path)
		os.Mkdir(path, 0755)
	} else if err != nil {
		return fmt.Errorf("Failed to create path")
	}

	//create stdout log
	path = this.ProcessTable[uuid].logPath + "stdout.log"
	fmt.Println(path)
	this.ProcessTable[uuid].stdoutPath = path
	if err := utils.CreateFile(path); err != nil {
		return fmt.Errorf("Failed to create %v due to %v", path, err)
	}

	//create stderr log
	path = this.ProcessTable[uuid].logPath + "stderr.log"
	fmt.Println(path)
	this.ProcessTable[uuid].stderrPath = path
	if err := utils.CreateFile(path); err != nil {
		return fmt.Errorf("Failed to create %v due to %v", path, err)
	}
	return nil

}

//update request from uuid 
func (this *DataStore) UpdateWithPid(uuid string, pid int, cmd *exec.Cmd) error {
	this.Lock()
	defer func() {
		this.Unlock()
	}()

	this.ProcessTable[uuid].pid = pid
	this.ProcessTable[uuid].isRunning = true
	this.ProcessTable[uuid].cmd = cmd

	if _, ok := this.ProcessTable[uuid]; !ok {
		return fmt.Errorf("Unable to Find ProcessInfo for %v\n", uuid)
	}
	return nil

}

func (this *DataStore) UpdateFinishProcess(uuid string, exitCode int) error {
	this.Lock()
	defer func() {
		this.Unlock()
	}()
	if _, ok := this.ProcessTable[uuid]; !ok {
		return fmt.Errorf("Unable to find process info for %v", uuid)
	}
	//return if process already ended
	if this.ProcessTable[uuid].endTimeStamp != "" {
		return fmt.Errorf("Process Already Ended for %v", uuid)
	}
	//update exit codes, timestamp, running status
	this.ProcessTable[uuid].exitCode = exitCode
	this.ProcessTable[uuid].isRunning = false
	this.ProcessTable[uuid].endTimeStamp = utils.GetTimeStamp()

	return nil
}	

func (this *DataStore) UpdateFailedProcessDidNotStart(uuid string) error {
	this.Lock()
	defer func() {
		this.Unlock()
	}()
	//return if process already ended
	if this.ProcessTable[uuid].endTimeStamp != "" {
		return fmt.Errorf("Process Already Ended for uuid: %v", uuid)
	}
	if _, ok := this.ProcessTable[uuid]; !ok {
		return fmt.Errorf("Unable to find process info for %v", uuid)
	}
	//update timestamp, running status
	this.ProcessTable[uuid].isRunning = false
	this.ProcessTable[uuid].endTimeStamp = utils.GetTimeStamp()
	
	return nil
}	

func (this *DataStore) GetListOfKeys() []string {
	this.RLock()
	defer func() {
		this.RUnlock()
	}()
	keys := make([]string, len(this.ProcessTable))
	i := 0
	for k := range this.ProcessTable {
	    keys[i] = k
	    i++
	}
	return keys
}

func (this *DataStore) getListOfKeys() []string {
	keys := make([]string, len(this.ProcessTable))
	i := 0
	for k := range this.ProcessTable {
	    keys[i] = k
	    i++
	}
	return keys
}

func (this *DataStore) GetProcessTable() ( map[string]*linuxJobWorkerPb.ProcessInfo, bool) {
	this.RLock()
	defer func() {
		this.RUnlock()
	}()
	listOfKeys := this.getListOfKeys()
	processTable := map[string]*linuxJobWorkerPb.ProcessInfo{}
	good := true
	for _, key := range listOfKeys{
		processInfo, ok := this.getRespProcessInfo(key)
		processTable[key] = processInfo
		if !ok {
			good = false
		}
	}
	return processTable, good
}

