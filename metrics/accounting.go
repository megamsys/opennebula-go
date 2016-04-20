package metrics

import (
	"github.com/megamsys/opennebula-go/api"
	"strconv"
	"time"
	log "github.com/Sirupsen/logrus"
)

type Accounting struct {
	Api       *api.Rpc
	StartTime int64
	EndTime   int64
}

func (a *Accounting) Get() ([]interface{}, error) {
	log.Debugf("showback Get (%d, %d) started", a.StartTime, a.EndTime)
	args := []interface{}{a.Api.Key, -2, -1, a.StartTime, a.EndTime}
	res, err := a.Api.Call(api.VMPOOL_ACCOUNTING, args)
	
	if err != nil {
		return nil, err
	}
	return res, nil
}

type VmState int
type LcmState int

const (
	//VmState starts at 0
	INIT VmState = iota
	PENDING
	HOLD
	ACTIVE
	STOPPED
	SUSPENDED
	DONE
	FAILED

	//LcmState starts at 0
	LCM_INIT LcmState = iota
	PROLOG
	BOOT
	RUNNING
	MIGRATE
	SAVE_STOP
	SAVE_SUSPEND
	SAVE_MIGRATE
	PROLOG_MIGRATE
	PROLOG_RESUME
	EPILOG_STOP
	EPILOG
	SHUTDOWN
	CANCEL
	FAILURE
	CLEANUP
	UNKNOWN
)

type History struct {
	HostName string `xml:"HOSTNAME"`
	Stime    string `xml:"STIME"`
	Etime    string `xml:"ETIME"`
	VM       *VM    `xml:"VM"`
}

type VM struct {
	Name      string    `xml:"NAME"`
	State     string    `xml:"STATE"`
	Lcm_state string    `xml:"LCM_STATE"`
	Stime     string    `xml:"STIME"`
	Etime     string    `xml:"ETIME"`
	Template  *Template `xml:"TEMPLATE"`
}

type Template struct {
	Context     Context `xml:"CONTEXT"`
	Cpu         string  `xml:"CPU"`
	Cpu_cost    string  `xml:"CPU_COST"`
	Vcpu        string  `xml:"VCPU"`
	Memory      string  `xml:"MEMORY"`
	Memory_cost string  `xml:"MEMORY_COST"`
	Disk_size   string  `xml:"SIZE"`
}

type Context struct {
	Name          string `xml:"NAME"`
	Accounts_id   string `xml:"ACCOUNTS_ID"`
	Assembly_id   string `xml:"ASSEMBLY_ID"`
	Assemblies_id string `xml:"ASSEMBLIES_ID"`
}

type OpenNebulaStatus struct {
	History_Records []*History `xml:"HISTORY"`
}

func (h *History) Cpu() string {
	return h.VM.Template.Cpu
}

func (h *History) CpuCost() string {
	return h.VM.Template.Cpu_cost
}

func (h *History) Memory() string {
	return h.VM.Template.Memory
}

func (h *History) MemoryCost() string {
	return h.VM.Template.Memory_cost
}

func (h *History) AssemblyName() string {
	return h.VM.Name
}

func (h *History) AccountsId() string {
	return h.VM.Template.Context.Accounts_id
}

func (h *History) AssembliesId() string {
	return h.VM.Template.Context.Assemblies_id
}

func (h *History) AssemblyId() string {
	return h.VM.Template.Context.Assembly_id
}

func (h *History) State() string {
	return h.VM.stateString()
}

func (h *History) LcmState() string {
	return h.VM.lcmStateString()
}

func TimeAsInt64(tm string) int64 {
	if i, err := strconv.ParseInt(tm, 10, 64); err != nil {
		return i
	}
	return 0
}

func (h *History) Elapsed() string {
	return strconv.FormatFloat(time.Since(time.Unix(TimeAsInt64(h.VM.Stime), 0)).Hours(), 'E', -1, 64)
}

func (v *VM) stateAsInt(s string) int {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}
	return 22
}

func (v *VM) stateString() string {
	switch VmState(v.stateAsInt(v.State)) {
	case INIT:
		return "Init"
	case PENDING:
		return "Pending"
	case HOLD:
		return "Hold"
	case ACTIVE:
		return "Active"
	case STOPPED:
		return "Stopped"
	case SUSPENDED:
		return "Suspended"
	case DONE:
		return "Done"
	case FAILED:
		return "Failed"
	default:
	}
	return "Unknown"
}
func (v *VM) lcmStateString() string {
	switch LcmState(v.stateAsInt(v.Lcm_state)) {
	case LCM_INIT:
		return "Lcm Init"
	case PROLOG:
		return "Prolog"
	case BOOT:
		return "Boot"
	case RUNNING:
		return "Running"
	case MIGRATE:
		return "Migrate"
	case SAVE_STOP:
		return "Save stop"
	case SAVE_SUSPEND:
		return "Save suspend"
	case SAVE_MIGRATE:
		return "Save migrate"
	case PROLOG_MIGRATE:
		return "Prolog migrate"
	case PROLOG_RESUME:
		return "Prolog resume"
	case EPILOG_STOP:
		return "Eplilog stop"
	case EPILOG:
		return "Epilog"
	case SHUTDOWN:
		return "Shutdown"
	case CANCEL:
		return "Cancel"
	case FAILURE:
		return "Failure"
	case CLEANUP:
		return "Cleanup"
	default:
		return "Unknown"
	}
}
