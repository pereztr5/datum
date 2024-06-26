// Code generated by enthistory, DO NOT EDIT.

// Code generated by ent, DO NOT EDIT.

package generated

import "github.com/datumforge/enthistory"

// withHistory adds the history hooks to the appropriate schemas - generated by enthistory
func (c *Client) WithHistory() {
	for _, hook := range enthistory.HistoryHooks[*OrganizationMutation]() {
		c.Organization.Use(hook)
	}
	for _, hook := range enthistory.HistoryHooks[*OrganizationSettingMutation]() {
		c.OrganizationSetting.Use(hook)
	}
}
