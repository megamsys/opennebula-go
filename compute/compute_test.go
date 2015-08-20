package compute

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateVM(t *testing.T) {
	Convey("CreateVM", t, func() {
		vmObj := VirtualMachine{OpenNebulaTemplateId: 2}
		creds := Credentials{Username: "oneadmin", Password: "RaifZuewjoc4", Endpoint: "http://localhost:2633/RPC2"}

		vmObj.CreateVM(&creds)
		//	So(err, ShouldBeNil)
	})
}
