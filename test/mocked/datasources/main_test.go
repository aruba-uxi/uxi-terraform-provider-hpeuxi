/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package data_source_test

import (
	"os"
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/util"
)

func TestMain(m *testing.M) {
	os.Setenv("HPEUXI_HOST_OVERRIDE", util.MockDomain)

	exitCode := m.Run()

	os.Unsetenv("HPEUXI_HOST_OVERRIDE")

	os.Exit(exitCode)
}
