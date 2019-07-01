package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

var testsys *System

// var testConfig *Config
// var testRulesEngine *RulesEngine

func setup() {
	testsys = NewSystem("test_config.json")

	// testConfig, err := LoadConfig("test_config.json")
	// if err != nil {
	// 	log.Fatal("Unable to start testing, can't load configuration, error: %s", err.Error())
	// }

	// re, err := NewRulesEngine(testConfig)
	// if err != nil {
	// 	log.Fatal("Unable to start testing, failed to create Rules Engine")
	// }
	// testRulesEngine = re
	// re.Debugging(true)
	testsys.RulesEngine().Debugging(true)
	testsys.RulesEngine().SetDeviceCache(testsys.DeviceCache())
}

func TestCreateRulesEngine(t *testing.T) {
	_, err := NewRulesEngine(testsys.Config())
	if err != nil {
		t.Error(err)
	}
}

func TestStripPort(t *testing.T) {
	addr := StripPortFromAddr("127.0.0.1:2323")
	if addr != "127.0.0.1" {
		t.Error(errors.New("Address is wrong"))
	}
}

func TestLocalHostPass(t *testing.T) {
	action, err := testsys.RulesEngine().Evaluate("*.rules.test", "127.0.0.1")
	if err != nil {
		t.Error(err)
	}
	if action != ActionTypePass {
		t.Errorf("Wrong action, expected 'ActionTypePass' got '%s'\n", action.String())
	}
}

func TestPassWithTimeSpan(t *testing.T) {
	tm := time.Now()
	r := Rule{
		Type:     ActionTypePass,
		TimeSpan: fmt.Sprintf("%d:00-%d:00", tm.Hour(), tm.Hour()+2),
	}

	action, err := r.EvaluateRule()
	if err != nil {
		t.Error(err)
	}
	if action != ActionTypePass {
		t.Errorf("Wrong action, expected 'ActionTypePAss' got %s\n", action.String())
	}

}

func TestLocalHostNone(t *testing.T) {
	action, err := testsys.RulesEngine().Evaluate("site2.rules.test", "127.0.0.2")
	if err != nil {
		t.Error(err)
	}
	log.Printf("Time: %s, action: %s\n", time.Now().Format("15:04"), action.String())

	// Note: Action depends on time

	// if action != ActionTypeNone {
	// 	t.Errorf("Wrong action, expected 'ActionTypeNone' got '%s'\n", action.String())
	// }
}

func TestLocalHostBlock(t *testing.T) {
	action, err := testsys.RulesEngine().Evaluate("site3.rules.test", "127.0.0.3")
	if err != nil {
		t.Error(err)
	}
	if action != ActionTypeBlockedDevice {
		t.Errorf("Wrong action, exptected 'ActionTypeBlockedDevice' got '%s'", action.String())
	}
}

func TestNameRulePass(t *testing.T) {
	action, err := testsys.RulesEngine().Evaluate("site1.rules.test", "192.168.1.8")
	if err != nil {
		t.Error(err)
	}

	log.Printf("Time: %s, action: %s\n", time.Now().Format("15:04"), action.String())

	// Note: The action here depends on the time.
	// 16:00 - 20:00 this will give a ban
	// All other times it will give a pass
	//
	// RulesEngine don't give an option to modify input time
	//

	// if action != ActionTypePass {
	// 	t.Errorf("Wrong action, exptected 'ActionTypePass' got '%s'", action.String())
	// }
}

func TestNameRuleLastBlock(t *testing.T) {
	action, err := testsys.RulesEngine().Evaluate("www.aftonbladet.se", "192.168.1.23")
	if err != nil {
		t.Error(err)
	}
	if action != ActionTypeBlockedSiteBan {
		t.Errorf("Wrong action, exptected 'ActionTypeBlockedSiteBan' got '%s'", action.String())
	}
}

func TestNameRuleBlock(t *testing.T) {
	action, err := testsys.RulesEngine().Evaluate("*.rules.test", "192.168.1.17")
	if err != nil {
		t.Error(err)
	}
	if action != ActionTypeBlockedDevice {
		t.Errorf("Wrong action, exptected 'ActionTypeBlockedDevice' got '%s'", action.String())
	}
}

// Note: This requires router initialization
func TestNameToIP(t *testing.T) {
	addr, err := testsys.DeviceCache().NameToIP("nagini")
	if err != nil {
		t.Error(err)
	}

	if addr.String() != "192.168.1.8" {
		t.Error(fmt.Errorf("Expected '192.168.1.8' for 'nagini' - got: %s\n", addr.String()))
	}
}

// Note: This requires router initialization
func TestIPToName(t *testing.T) {
	name, err := testsys.DeviceCache().IPToName("192.168.1.8")
	if err != nil {
		t.Error(err)
	}

	if strings.ToLower(name) != "nagini" {
		t.Error(fmt.Errorf("Expected 'nagini' for 192.168.1.8' - got: %s\n", name))
	}
}

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	setup()
	os.Exit(m.Run())
}
