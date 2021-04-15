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

### Install daemon
Copy the initd.script to the /etc/init.d/ folder and chmod +x.
scp initd.script /etc/init.d/x2webdimming

From ssh terminal, add the service to run at bootup:
update-rc.d x2webdimming defaults