package compute

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateVM(t *testing.T) {
	Convey("CreateVM", t, func() {
		vmObj := VirtualMachine{OpenNebulaTemplateName: "fullfledged", Cpu: "10", VCpu: "10", Memory: "4"}
		creds := Credentials{Username: "oneadmin", Password: "RaifZuewjoc4", Endpoint: "http://localhost:2633/RPC2"}

		vmObj.CreateVM(&creds)
		//	So(err, ShouldBeNil)
	})
}
