package core

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
)

// LogoEncoded : base64 encoded logo, ascii art
const LogoEncoded = "ICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgXyBfXyBfX18gICBfX18gIF9fXyAgICAgIF8gX18gICBfXyBfIAogfCAnXyBgIF8gXCAvIF8gXC8gX198X19fX3wgJ18gXCAvIF9gIHwKIHwgfCB8IHwgfCB8ICBfXy8gKF98X19fX198IHwgfCB8IChffCB8CiB8X3wgfF98IHxffFxfX198XF9fX3wgICAgfF98IHxffFxfXywgfAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICB8X19fLyAK"

var (
	// IPList : target list
	IPList string
	// Mode : working mode
	Mode string
	// JobCnt : how many tasks per time
	JobCnt int
	// Module : which module to use
	Module string
	// TailArgs : more args in the tail
	TailArgs []string
	// UseProxy : whether use shadowsocks for anonymity or not
	UseProxy bool
	// Filter : the filter to use when parsing masscan xml
	Filter string
	// MasscanXML : specify the source XML for xmir
	MasscanXML string
	// MasscanRange : ip range list for masscan
	MasscanRange string
)

// ArgParse : parse cmd line args for package core
func ArgParse() {
	flag.StringVar(&IPList, "list", "", "target ip list, useful in custom, exp and bruteforce mode")
	flag.StringVar(&Mode, "mode", "", "working mode, can be one of the following:\ncustom, zoomeye, masscan")
	flag.StringVar(&Filter, "filter", "", "the filter (for banners) to use when parsing masscan xml")
	flag.StringVar(&MasscanXML, "xml", "", "specify the source XML for xmir")
	flag.StringVar(&MasscanRange, "mlist", "", "scan list for masscan")
	flag.IntVar(&JobCnt, "thd", 100, "how many tasks per time")
	flag.StringVar(&Module, "module", "", "in custom mode, this is the executable to run")
	flag.BoolVar(&UseProxy, "useproxy", false, "use shadowsocks or not")

	flag.Parse()

	// TailArgs = flag.Args()
}

// PrintBanner : print mec-ng ascii logo
func PrintBanner() {
	logo, err := base64.StdEncoding.DecodeString(LogoEncoded)
	if err != nil {
		log.Panic("Logo error: ", err)
	}

	color.Set(color.FgCyan, color.Bold)
	defer color.Unset()
	fmt.Println(string(logo))
	fmt.Println(strings.Repeat(" ", 26) + "by jm33-ng\n")
	fmt.Println("examples:\n mec-ng -mode custom -module ./built-in/exp/exp -list ./conf/list -thd 50 -useproxy")
	fmt.Println(" mec-ng -mode masscan -mlist /tmp/range.list")
	fmt.Println(" mec-ng -mode xmir -xml ./output/masscan.xml -filter 'SSH-2.0-OpenSSH_7.4p1'")
	fmt.Println(" mec-ng -mode bruteforce -list ./conf/list")
	fmt.Print(" mec-ng -mode zoomeye\n\n\n")
}
