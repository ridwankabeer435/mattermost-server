// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"reflect"
	"strconv"
)

type FeatureFlags struct {
	// Exists only for unit and manual testing.
	// When set to a value, will be returned by the ping endpoint.
	TestFeature string
	// Exists only for testing bool functionality. Boolean feature flags interpret "on" or "true" as true and
	// all other values as false.
	TestBoolFeature bool

	// Enable the remote cluster service for shared channels.
	EnableRemoteClusterService bool

	// AppsEnabled toggles the Apps framework functionalities both in server and client side
	AppsEnabled bool

	// Feature flags to control plugin versions
	PluginPlaybooks  string `plugin_id:"playbooks"`
	PluginApps       string `plugin_id:"com.mattermost.apps"`
	PluginFocalboard string `plugin_id:"focalboard"`
	PluginCalls      string `plugin_id:"com.mattermost.calls"`

	PermalinkPreviews bool

	// CallsEnabled controls whether or not the Calls plugin should be enabled
	CallsEnabled bool

	// Enable DataRetention for Boards
	BoardsDataRetention bool

	NormalizeLdapDNs bool

	// Enable GraphQL feature
	GraphQL bool

	PostPriority bool

	// Enable WYSIWYG text editor
	WysiwygEditor bool

	PeopleProduct bool

	OnboardingTourTips bool

	DeprecateCloudFree bool

	CloudReverseTrial bool

	EnableExportDirectDownload bool

	DataRetentionConcurrencyEnabled bool
}

func (f *FeatureFlags) SetDefaults() {
	f.TestFeature = "off"
	f.TestBoolFeature = false
	f.EnableRemoteClusterService = false
	f.AppsEnabled = true
	f.PluginApps = ""
	f.PluginFocalboard = ""
	f.BoardsDataRetention = false
	f.NormalizeLdapDNs = false
	f.GraphQL = false
	f.CallsEnabled = true
	f.PeopleProduct = false
	f.DeprecateCloudFree = false
	f.WysiwygEditor = false
	f.OnboardingTourTips = true
	f.CloudReverseTrial = false
	f.EnableExportDirectDownload = false
	f.DataRetentionConcurrencyEnabled = true
}

func (f *FeatureFlags) Plugins() map[string]string {
	rFFVal := reflect.ValueOf(f).Elem()
	rFFType := reflect.TypeOf(f).Elem()

	pluginVersions := make(map[string]string)
	for i := 0; i < rFFVal.NumField(); i++ {
		rFieldVal := rFFVal.Field(i)
		rFieldType := rFFType.Field(i)

		pluginId, hasPluginId := rFieldType.Tag.Lookup("plugin_id")
		if !hasPluginId {
			continue
		}

		pluginVersions[pluginId] = rFieldVal.String()
	}

	return pluginVersions
}

// ToMap returns the feature flags as a map[string]string
// Supports boolean and string feature flags.
func (f *FeatureFlags) ToMap() map[string]string {
	refStructVal := reflect.ValueOf(*f)
	refStructType := reflect.TypeOf(*f)
	ret := make(map[string]string)
	for i := 0; i < refStructVal.NumField(); i++ {
		refFieldVal := refStructVal.Field(i)
		if !refFieldVal.IsValid() {
			continue
		}
		refFieldType := refStructType.Field(i)
		switch refFieldType.Type.Kind() {
		case reflect.Bool:
			ret[refFieldType.Name] = strconv.FormatBool(refFieldVal.Bool())
		default:
			ret[refFieldType.Name] = refFieldVal.String()
		}
	}

	return ret
}
