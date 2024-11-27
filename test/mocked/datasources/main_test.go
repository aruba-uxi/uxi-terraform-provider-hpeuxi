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
	os.Setenv("UXI_HOST", util.MockDomain)

	exitCode := m.Run()

	os.Unsetenv("UXI_HOST")

	os.Exit(exitCode)
}
