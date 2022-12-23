package testdata

const (
	shop1Data = `<shop id="1">
  <name>TTN</name>
  <url>ttn.by</url>
  <working_time>
    <open>09:00:00</open>
    <close>18:00:00</close>
  </working_time>
  <offers>
    <item id="1">
      <name>Phone</name>
      <description>Samsung S11 Edge Best choice on the market</description>
      <price>1257</price>
    </item>
    <item id="2">
      <name>Notebook</name>
      <description>Asus ZenBook Stylish modern notebook for productive work</description>
      <price>4548</price>
    </item>
  </offers>
</shop>
`
	shop2Data = `<shop id="2">
  <name>5 Element</name>
  <url>5element.by</url>
  <working_time>
    <open>11:30:00</open>
    <close>20:30:00</close>
  </working_time>
  <offers>
    <item id="3">
      <name>Monitor</name>
      <description>HP 2W925AA Best choice for working and gaming</description>
      <price>1257</price>
    </item>
    <item id="4">
      <name>Computer</name>
      <description>MultiGame 5C104FD Reliable high performance package</description>
      <price>4548</price>
    </item>
  </offers>
</shop>
`
)

var (
	Shop1       = []byte(shop1Data)
	Shop2       = []byte(shop2Data)
	BothShopsV1 = []byte(shop1Data + shop2Data)
	BothShopsV2 = []byte(shop2Data + shop1Data)
)
