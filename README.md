# README #

This is a (very) rudimentary example on how to remote control the X2 web device. In this example we are controlling the brightness only.

### Building for Linux/Arm32 ###
GOOS=linux GOARCH=arm go build -o dimmingsvc x2webdimming.go

### Notes on persistence ###
Dimming value is now persistent from the service but hardcoded to the file /etc/beijer/misc.conf.

### Copy files to the device from a PC
Login over ssh and remount filesystem as read/write:
mount -o remount,rw /

Use scp from PC to copy files. Ex:
scp source-filename root@{x2web ip address}:/path/dest-filename

### Install systemd service
Remount fs to read/write and store the file x2dimming.service in the /etc/systemd/system/ folder.
This is the new preferred way, the init.d folder was removed in 2022.74.

To copy the file, use for instance, $ scp x2dimming.service root@{ipaddress}:/etc/systemd/system/x2dimming.service

### Install init.d daemon
[Deprecated since 2022.74]
Copy the initd.script to the /etc/init.d/ folder and chmod +x.
scp initd.script /etc/init.d/x2webdimming

From ssh terminal, add the service to run at bootup:
update-rc.d x2webdimming defaults

