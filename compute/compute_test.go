package compute

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateVM(t *testing.T) {
	Convey("CreateVM", t, func() {
		vmObj := VirtualMachine{OpenNebulaTemplateName: "yeshdeb", Cpu: "1", VCpu: "1", Memory: "400"} //memory interms of MB! duh!
		//creds := Credentials{Username: "oneadmin", Password: "RaifZuewjoc4", Endpoint: "http://localhost:2633/RPC2"}
		creds := Credentials{Username: "oneadmin", Password: "yib4OquafUp1", Endpoint: "http://192.168.1.100:2633/RPC2"}

		vmObj.CreateVM(&creds)
		//	So(err, ShouldBeNil)
	})
}
