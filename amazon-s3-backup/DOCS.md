# Home Assistant Add-on: S3 Backup

## How to use

1. Set the config options
2. Use an automation to run the add-on - it does not run continuously

## Automation

To automate your backup creation and syncing to Amazon S3, add these two automations in Home Assistants `configuration.yaml` and change it to your needs:
```
automation:
  # create a full backup
  - id: backup_create_full_backup
    alias: Create a full backup every day at 4am
    trigger:
      platform: time
      at: "04:00:00"
    action:
      service: hassio.backup_full
      data:
        # uses the 'now' object of the trigger to create a more user friendly name (e.g.: '202101010400_automated-backup')
        name: "{{as_timestamp(trigger.now)|timestamp_custom('%Y%m%d%H%M', true)}}_automated-backup"

  # Starts the addon 15 minutes after every hour to make sure it syncs all backups, also manual ones, as soon as possible
  - id: backup_upload_to_s3
    alias: Upload to S3
    trigger:
      platform: time_pattern
      # Matches every hour at 15 minutes past every hour
      minutes: 15
    action:
      service: hassio.addon_start
      data:
      addon: XXXXX_amazon-s3-backup
```

The automation above first creates a full backup at 4am, and then at 4:15am syncs to Amazon S3 and if configured deletes local backups according to your configuration.
