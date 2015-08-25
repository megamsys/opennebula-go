package flavor

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

/*
func TestGetTemplate(t *testing.T) {
	Convey("GetTemplate", t, func() {
		endpoint := "http://localhost:2633/RPC2"
		key := "oneadmin:RaifZuewjoc4"
		flav := FlavorOpts{TemplateId: 3}
		_, error := flav.GetTemplate(endpoint, key)
		So(error, ShouldBeNil)
	})
}

*/

func TestGetTemplateByName(t *testing.T) {
	Convey("GetTemplateByName", t, func() {

		endpoint := "http://localhost:2633/RPC2"
		key := "oneadmin:RaifZuewjoc4"
		flav := FlavorOpts{TemplateName: "test11"}
		_, error := flav.GetTemplateByName(endpoint, key)
		So(error, ShouldBeNil)

	})
}

/*

func TestUpdateTemplate(t *testing.T) {

	Convey("UpdateTemplate", t, func() {

		endpoint := "http://localhost:2633/RPC2"
		key := "oneadmin:RaifZuewjoc4"
		data := `<TEMPLATE><CONTEXT><NETWORK><![CDATA[YES]]></NETWORK><SSH_PUBLIC_KEY><![CDATA[$USER[SSH_PUBLIC_KEY]]]></SSH_PUBLIC_KEY></CONTEXT><CPU><![CDATA[5]]></CPU><CPU_COST><![CDATA[10]]></CPU_COST><DESCRIPTION><![CDATA[testtemplate...]]></DESCRIPTION><HYPERVISOR><![CDATA[kvm]]></HYPERVISOR><LOGO><![CDATA[images/logos/ubuntu.png]]></LOGO><MEMORY><![CDATA[512]]></MEMORY><MEMORY_COST><![CDATA[10]]></MEMORY_COST><SUNSTONE_CAPACITY_SELECT><![CDATA[YES]]></SUNSTONE_CAPACITY_SELECT><SUNSTONE_NETWORK_SELECT><![CDATA[YES]]></SUNSTONE_NETWORK_SELECT></TEMPLATE>`

		flav := FlavorOpts{TemplateId: 25, TemplateData: data}

		error := flav.UpdateTemplate(endpoint, key)
		So(error, ShouldBeNil)

	})
}
*/
