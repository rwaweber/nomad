package config

import (
	"reflect"
	"testing"
	"time"
)

func TestVaultConfig_Merge(t *testing.T) {
	trueValue, falseValue := true, false
	c1 := &VaultConfig{
		Enabled:              &falseValue,
		Token:                "1",
		Role:                 "1",
		AllowUnauthenticated: &trueValue,
		TaskTokenTTL:         "1",
		Addr:                 "1",
		ConnectionRetryIntv:  time.Nanosecond * 5,
		TLSCaFile:            "1",
		TLSCaPath:            "1",
		TLSCertFile:          "1",
		TLSKeyFile:           "1",
		TLSSkipVerify:        &trueValue,
		TLSServerName:        "1",
	}

	c2 := &VaultConfig{
		Enabled:              &trueValue,
		Token:                "2",
		Role:                 "2",
		AllowUnauthenticated: &falseValue,
		TaskTokenTTL:         "2",
		Addr:                 "2",
		ConnectionRetryIntv:  time.Nanosecond * 5,
		TLSCaFile:            "2",
		TLSCaPath:            "2",
		TLSCertFile:          "2",
		TLSKeyFile:           "2",
		TLSSkipVerify:        nil,
		TLSServerName:        "2",
	}

	e := &VaultConfig{
		Enabled:              &trueValue,
		Token:                "2",
		Role:                 "2",
		AllowUnauthenticated: &falseValue,
		TaskTokenTTL:         "2",
		Addr:                 "2",
		ConnectionRetryIntv:  time.Nanosecond * 5,
		TLSCaFile:            "2",
		TLSCaPath:            "2",
		TLSCertFile:          "2",
		TLSKeyFile:           "2",
		TLSSkipVerify:        &trueValue,
		TLSServerName:        "2",
	}

	result := c1.Merge(c2)
	if !reflect.DeepEqual(result, e) {
		t.Fatalf("bad:\n%#v\n%#v", result, e)
	}

	// merging the other way should not yield the same results
	result2 := c2.Merge(c1)
	if reflect.DeepEqual(result2, e) {
		t.Fatalf("bad:\n%#v\n%#v", result2, e)
	}

}

func TestVaultConfig_IsEnabledIsAuthenticated(t *testing.T) {
	trueValue, falseValue := true, false
	cases := []*VaultConfig{
		&VaultConfig{
			Enabled:              &trueValue,
			Token:                "1",
			Role:                 "1",
			AllowUnauthenticated: &trueValue,
			TaskTokenTTL:         "1",
			Addr:                 "1",
			ConnectionRetryIntv:  time.Nanosecond * 5,
			TLSCaFile:            "1",
			TLSCaPath:            "1",
			TLSCertFile:          "1",
			TLSKeyFile:           "1",
			TLSSkipVerify:        &trueValue,
			TLSServerName:        "1",
		},
		&VaultConfig{
			Enabled:              &falseValue,
			Token:                "1",
			Role:                 "1",
			AllowUnauthenticated: &falseValue,
			TaskTokenTTL:         "1",
			Addr:                 "1",
			ConnectionRetryIntv:  time.Nanosecond * 5,
			TLSCaFile:            "1",
			TLSCaPath:            "1",
			TLSCertFile:          "1",
			TLSKeyFile:           "1",
			TLSSkipVerify:        &trueValue,
			TLSServerName:        "1",
		},
	}

	if !cases[0].AllowsUnauthenticated() {
		t.Fatalf("Should allow unauthenticated")
	}

	if !cases[0].IsEnabled() {
		t.Fatalf("Should be enabled")
	}

	if cases[1].AllowsUnauthenticated() {
		t.Fatalf("Should not allow authenticated")
	}

	if cases[1].IsEnabled() {
		t.Fatalf("Should not be enabled")
	}

}

func TestVaultConfig_copy(t *testing.T) {
	trueValue, falseValue := true, false
	c1 := &VaultConfig{
		Enabled:              &falseValue,
		Token:                "1",
		Role:                 "1",
		AllowUnauthenticated: &trueValue,
		TaskTokenTTL:         "1",
		Addr:                 "1",
		ConnectionRetryIntv:  time.Nanosecond * 5,
		TLSCaFile:            "1",
		TLSCaPath:            "1",
		TLSCertFile:          "1",
		TLSKeyFile:           "1",
		TLSSkipVerify:        &trueValue,
		TLSServerName:        "1",
	}

	realcopy := c1.Copy()

	if !reflect.DeepEqual(realcopy, c1) {
		t.Fatalf("bad:\n%#v\n%#v", realcopy, c1)
	}

}

func TestVaultConfig_DefaultVaultConfig(t *testing.T) {
	result := DefaultVaultConfig()

	expected := &VaultConfig{
		Addr:                "https://vault.service.consul:8200",
		ConnectionRetryIntv: DefaultVaultConnectRetryIntv,
		AllowUnauthenticated: func(b bool) *bool {
			return &b
		}(true),
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("bad:\n%#v\n%#v", result, expected)
	}

}
