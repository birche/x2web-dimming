# README #

This is a (very) rudimentary example on how to remote control the X2 web device. In this example we are controlling the brightness only, and solely for runtime, i.e. it is not persistent.

### Building for Linux/Arm32 ###

GOOS=linux ARCH=arm go build -o dimmingsvc x2webdimming.go

### Notes on persistence ###
If there is a need for dimming settings to be persistent on power cycles, the file /etc/beijer/misc.conf has to be edited and saved for every new dimming value change.
