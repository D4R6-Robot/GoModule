package mallist

// MalwareInfo struct
type MalwareInfo struct {
	malIndex         string
	malName          string
	y                int
	procTree         int
	procXSLCnt       int
	procJSCnt        int
	procMSFormsCnt   int
	procMacrosCnt    int
	netHTTPCnt       int
	ntQueryKey       int
	ntOpenKey        int
	ntOpenFile       int
	ntWriteFile      int
	ntQueryValueKey  int
	ntDelayExecution int
	ntCreateFile     int
	fileCreated      int
	fileReCreated    int
	fileDllLoaded    int
	fileOpened       int
	fileWritten      int
	fileFailed       int
}

// NewMalwerInfo function
func (malData *MalwareInfo) NewMalwerInfo() *MalwareInfo {
	return &(MalwareInfo{})
}

// GetMalIndex function
func (malData *MalwareInfo) GetMalIndex() string {
	return malData.malIndex
}

// SetMalIndex function
func (malData *MalwareInfo) SetMalIndex(malIndex string) {
	malData.malIndex = malIndex
}

// GetMalName function
func (malData *MalwareInfo) GetMalName() string {
	return malData.malName
}

// SetMalName function
func (malData *MalwareInfo) SetMalName(malName string) {
	malData.malName = malName
}

// GetY function
func (malData *MalwareInfo) GetY() int {
	return malData.y
}

// SetY function
func (malData *MalwareInfo) SetY(y int) {
	malData.y = y
}

// GetProcTree function
func (malData *MalwareInfo) GetProcTree() int {
	return malData.procTree
}

// SetProcTree function
func (malData *MalwareInfo) SetProcTree(procTree int) {
	malData.procTree = procTree
}

// GetProcXSLCnt function
func (malData *MalwareInfo) GetProcXSLCnt() int {
	return malData.procXSLCnt
}

// SetProcXSLCnt struct
func (malData *MalwareInfo) SetProcXSLCnt(procXSLCnt int) {
	malData.procXSLCnt = procXSLCnt
}

// GetProcJSCnt struct
func (malData *MalwareInfo) GetProcJSCnt() int {
	return malData.procJSCnt
}

// SetProcJSCnt struct
func (malData *MalwareInfo) SetProcJSCnt(procJSCnt int) {
	malData.procJSCnt = procJSCnt
}

// GetProcMSFormsCnt struct
func (malData *MalwareInfo) GetProcMSFormsCnt() int {
	return malData.procMSFormsCnt
}

// SetProcMSFormsCnt struct
func (malData *MalwareInfo) SetProcMSFormsCnt(procMSFormsCnt int) {
	malData.procMSFormsCnt = procMSFormsCnt
}

// GetProcMacrosCnt struct
func (malData *MalwareInfo) GetProcMacrosCnt() int {
	return malData.procMacrosCnt
}

// SetProcMacrosCnt struct
func (malData *MalwareInfo) SetProcMacrosCnt(procMacrosCnt int) {
	malData.procMacrosCnt = procMacrosCnt
}

// GetNetHTTPCnt function
func (malData *MalwareInfo) GetNetHTTPCnt() int {
	return malData.netHTTPCnt
}

// SetNetHTTPCnt function
func (malData *MalwareInfo) SetNetHTTPCnt(netHTTPCnt int) {
	malData.netHTTPCnt = netHTTPCnt
}

// GetNtQueryKey function
func (malData *MalwareInfo) GetNtQueryKey() int {
	return malData.ntQueryKey
}

// SetNtQueryKey function
func (malData *MalwareInfo) SetNtQueryKey(ntQueryKey int) {
	malData.ntQueryKey = ntQueryKey
}

// GetNtOpenKey function
func (malData *MalwareInfo) GetNtOpenKey() int {
	return malData.ntOpenKey
}

// SetNtOpenKey function
func (malData *MalwareInfo) SetNtOpenKey(ntOpenKey int) {
	malData.ntOpenKey = ntOpenKey
}

// GetNtOpenFile function
func (malData *MalwareInfo) GetNtOpenFile() int {
	return malData.ntOpenFile
}

// SetNtOpenFile function
func (malData *MalwareInfo) SetNtOpenFile(ntOpenFile int) {
	malData.ntOpenFile = ntOpenFile
}

// GetNtWriteFile function
func (malData *MalwareInfo) GetNtWriteFile() int {
	return malData.ntWriteFile
}

// SetNtWriteFile struct
func (malData *MalwareInfo) SetNtWriteFile(ntWriteFile int) {
	malData.ntWriteFile = ntWriteFile
}

// GetNtQueryValueKey struct
func (malData *MalwareInfo) GetNtQueryValueKey() int {
	return malData.ntQueryValueKey
}

// SetNtQueryValueKey struct
func (malData *MalwareInfo) SetNtQueryValueKey(ntQueryValueKey int) {
	malData.ntQueryValueKey = ntQueryValueKey
}

// GetNtDelayExecution struct
func (malData *MalwareInfo) GetNtDelayExecution() int {
	return malData.ntDelayExecution
}

// SetNtDelayExecution struct
func (malData *MalwareInfo) SetNtDelayExecution(ntDelayExecution int) {
	malData.ntDelayExecution = ntDelayExecution
}

// GetNtCreateFile struct
func (malData *MalwareInfo) GetNtCreateFile() int {
	return malData.ntCreateFile
}

// SetNtCreateFile struct
func (malData *MalwareInfo) SetNtCreateFile(ntCreateFile int) {
	malData.ntCreateFile = ntCreateFile
}

// GetFileCreated struct
func (malData *MalwareInfo) GetFileCreated() int {
	return malData.fileCreated
}

// SetFileCreated struct
func (malData *MalwareInfo) SetFileCreated(fileCreated int) {
	malData.fileCreated = fileCreated
}

// GetFileReCreated function
func (malData *MalwareInfo) GetFileReCreated() int {
	return malData.fileReCreated
}

// SetFileReCreated struct
func (malData *MalwareInfo) SetFileReCreated(fileReCreated int) {
	malData.fileReCreated = fileReCreated
}

// GetFileDllLoaded struct
func (malData *MalwareInfo) GetFileDllLoaded() int {
	return malData.fileDllLoaded
}

// SetFileDllLoaded struct
func (malData *MalwareInfo) SetFileDllLoaded(fileDllLoaded int) {
	malData.fileDllLoaded = fileDllLoaded
}

// GetFileOpened struct
func (malData *MalwareInfo) GetFileOpened() int {
	return malData.fileOpened
}

// SetFileOpened struct
func (malData *MalwareInfo) SetFileOpened(fileOpened int) {
	malData.fileOpened = fileOpened
}

// GetFileWritten struct
func (malData *MalwareInfo) GetFileWritten() int {
	return malData.fileWritten
}

// SetFileWritten struct
func (malData *MalwareInfo) SetFileWritten(fileWritten int) {
	malData.fileWritten = fileWritten
}

// GetFileFailed struct
func (malData *MalwareInfo) GetFileFailed() int {
	return malData.fileFailed
}

// SetFileFailed struct
func (malData *MalwareInfo) SetFileFailed(fileFailed int) {
	malData.fileFailed = fileFailed
}
