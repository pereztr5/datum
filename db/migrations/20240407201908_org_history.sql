-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_organization_settings" table
CREATE TABLE `new_organization_settings` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `domains` json NULL, `billing_contact` text NULL, `billing_email` text NULL, `billing_phone` text NULL, `billing_address` text NULL, `tax_identifier` text NULL, `tags` json NULL, `geo_location` text NULL DEFAULT ('AMER'), `organization_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `organization_settings_organizations_setting` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "organization_settings" to new temporary table "new_organization_settings"
INSERT INTO `new_organization_settings` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `domains`, `billing_contact`, `billing_email`, `billing_phone`, `billing_address`, `tax_identifier`, `tags`, `geo_location`, `organization_id`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `domains`, `billing_contact`, `billing_email`, `billing_phone`, `billing_address`, `tax_identifier`, `tags`, `geo_location`, `organization_id` FROM `organization_settings`;
-- Drop "organization_settings" table after copying rows
DROP TABLE `organization_settings`;
-- Rename temporary table "new_organization_settings" to "organization_settings"
ALTER TABLE `new_organization_settings` RENAME TO `organization_settings`;
-- Create index "organization_settings_organization_id_key" to table: "organization_settings"
CREATE UNIQUE INDEX `organization_settings_organization_id_key` ON `organization_settings` (`organization_id`);
-- Create "organization_history" table
CREATE TABLE `organization_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `display_name` text NOT NULL DEFAULT (''), `description` text NULL, `parent_organization_id` text NULL, `personal_org` bool NULL DEFAULT (false), `avatar_remote_url` text NULL, `dedicated_db` bool NOT NULL DEFAULT (false), PRIMARY KEY (`id`));
-- Create index "organizationhistory_history_time" to table: "organization_history"
CREATE INDEX `organizationhistory_history_time` ON `organization_history` (`history_time`);
-- Create "organization_setting_history" table
CREATE TABLE `organization_setting_history` (`id` text NOT NULL, `history_time` datetime NOT NULL, `operation` text NOT NULL, `ref` text NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `domains` json NULL, `billing_contact` text NULL, `billing_email` text NULL, `billing_phone` text NULL, `billing_address` text NULL, `tax_identifier` text NULL, `tags` json NULL, `geo_location` text NULL DEFAULT ('AMER'), `organization_id` text NULL, PRIMARY KEY (`id`));
-- Create index "organizationsettinghistory_history_time" to table: "organization_setting_history"
CREATE INDEX `organizationsettinghistory_history_time` ON `organization_setting_history` (`history_time`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
