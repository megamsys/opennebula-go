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


func TestGetTemplateByName(t *testing.T) {
	Convey("GetTemplateByName", t, func() {

		endpoint := "http://localhost:2633/RPC2"
		key := "oneadmin:RaifZuewjoc4"
		flav := FlavorOpts{TemplateName: "ubuntu"}
		_, error := flav.GetTemplateByName(endpoint, key)
		So(error, ShouldBeNil)

	})
}
*/


func TestUpdateTemplate(t *testing.T) {

	Convey("UpdateTemplate", t, func() {

		endpoint := "http://localhost:2633/RPC2"
		key := "oneadmin:RaifZuewjoc4"
		data := `<VMTEMPLATE><ID>6</ID><UID>0</UID><GID>3</GID><UNAME>osrasdsddmin</UNAME><GNAME>yeshwsdsanthsdeadmin</GNAME><NAME>supertest</NAME><PERMISSIONS><OWNER_U>1</OWNER_U><OWNER_M>1</OWNER_M><OWNER_A>0</OWNER_A><GROUP_U>0</GROUP_U><GROUP_M>0</GROUP_M><GROUP_A>0</GROUP_A><OTHER_U>0</OTHER_U><OTHER_M>0</OTHER_M><OTHER_A>0</OTHER_A></PERMISSIONS><REGTIME>1440059694</REGTIME><TEMPLATE><CONTEXT><NETWORK><![CDATA[YES]]></NETWORK><SSH_PUBLIC_KEY><![CDATA[$USER[SSH_PUBLIC_KEY]]]></SSH_PUBLIC_KEY></CONTEXT><CPU><![CDATA[1]]></CPU><CPU_COST><![CDATA[10]]></CPU_COST><DESCRIPTION><![CDATA[testtemplate...]]></DESCRIPTION><HYPERVISOR><![CDATA[kvm]]></HYPERVISOR><LOGO><![CDATA[images/logos/ubuntu.png]]></LOGO><MEMORY><![CDATA[512]]></MEMORY><MEMORY_COST><![CDATA[10]]></MEMORY_COST><SUNSTONE_CAPACITY_SELECT><![CDATA[YES]]></SUNSTONE_CAPACITY_SELECT><SUNSTONE_NETWORK_SELECT><![CDATA[YES]]></SUNSTONE_NETWORK_SELECT></TEMPLATE></VMTEMPLATE>`

		flav := FlavorOpts{TemplateId: 6, TemplateData: data}

	 error := flav.UpdateTemplate(endpoint, key)
		So(error, ShouldBeNil)

	})
}
