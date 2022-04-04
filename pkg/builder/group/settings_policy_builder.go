package group

import (
	"fmt"

	"github.com/turbonomic/turbo-go-sdk/pkg/builder"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

type SettingPolicybuilder struct {
	name        string
	displayName string
	settings    []*proto.GroupDTO_Setting
	ec          *builder.ErrorCollector
}

func NewSettingPolicyBuilder() *SettingPolicybuilder {
	return &SettingPolicybuilder{}
}

func (spb *SettingPolicybuilder) WithName(name string) *SettingPolicybuilder {
	spb.name = name
	return spb
}

func (spb *SettingPolicybuilder) WithDisplayName(displayName string) *SettingPolicybuilder {
	spb.displayName = displayName
	return spb
}

func (spb *SettingPolicybuilder) WithSettings(settings []*proto.GroupDTO_Setting) *SettingPolicybuilder {
	spb.settings = settings
	return spb
}

func (spb *SettingPolicybuilder) Build() (*proto.GroupDTO_SettingPolicy_, error) {
	settingsDto := &proto.GroupDTO_SettingPolicy_{
		SettingPolicy: &proto.GroupDTO_SettingPolicy{},
	}

	if spb.name == "" {
		spb.ec.Collect(fmt.Errorf("Name is required"))
	}
	if spb.displayName == "" {
		spb.ec.Collect(fmt.Errorf("DisplayName is required"))
	}

	if spb.ec.Count() > 0 {
		return nil, fmt.Errorf("error building settings policy: %s", spb.ec.Error())
	}
	settingsDto.SettingPolicy.Name = &spb.name
	settingsDto.SettingPolicy.DisplayName = &spb.displayName
	settingsDto.SettingPolicy.Settings = spb.settings

	return settingsDto, nil
}

type SettingsBuilder struct {
	Settings []*proto.GroupDTO_Setting
}

func NewSettingsBuilder() *SettingsBuilder {
	return &SettingsBuilder{}
}

func (sb *SettingsBuilder) AddSetting(setting *proto.GroupDTO_Setting) *SettingsBuilder {
	sb.Settings = append(sb.Settings, setting)
	return sb
}

func (sb *SettingsBuilder) Build() []*proto.GroupDTO_Setting {
	return sb.Settings
}

func NewResizeAutomationPolicySetting(actionCapability string) *proto.GroupDTO_Setting {
	return &proto.GroupDTO_Setting{
		Type: proto.GroupDTO_Setting_RESIZE_AUTOMATION_MODE.Enum(),
		SettingValueType: &proto.GroupDTO_Setting_StringSettingValueType_{
			StringSettingValueType: &proto.GroupDTO_Setting_StringSettingValueType{
				Value: &actionCapability,
			},
		},
	}
}
