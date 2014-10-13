nano ~/.ssh/authorized_keys
mkdir .ssh
exit
raspi-config
sudo raspi-config
sudo shutdown -r 
sudo shutdown -r now
ls
sudo apt-get install libusb-dev  libdbus-1-dev 
sudo apt-get install libglib2.0-dev --fix-missing
sudo apt-get install libudev-dev libical-dev libreadline-dev
sudo apt-get update
sudo apt-get upda         
sudo apt-get upgrade
sudo mkdir bluez
cd bluez/
sudo wget www.kernel.org/pub/linux/bluetooth/bluez-5.18.tar.gz
sudo gunzip bluez-5.18.tar.gz
sudo tar xvf bluez-5.18.tar
cd bluez-5.18
sudo ./configure --disable-systemd
sudo make
./bt.sh
sudo btle/hcitool -i hci0 cmd 0x08 0x0008 1E 02 01 1A 1A FF 4C 00 02 15 D9 B9 EC 1F 39 25 43 D0 80 A9 1E 39 D4 CE A9 5C 02 03 10 04 C8 00
nano bt.sh 
nano bt_down.sh
chmod +x bt_down.sh 
./bt_down.sh
ls
./yizla 
./bt.sh
./yizla 
sudo btle/hcitool -i hci0 cmd 0x08 0x0008 1E 02 01 1A 1A FF 4C 00 02 15 D9 B9 EC 1F 39 25 43 D0 80 A9 1E 39 D4 CE A9 5C 02 03 10 04 C8 00
sudo btle/hcitool -i hci0 cmd 0x08 0x0008 1E 02 01 1A 1A FF 4C 00 02 15 E2 0A 39 F4 73 F5 4B C4 A1 2F 17 D1 AD 07 A9 61 00 00 00 00 C8 00
sudo btle/hcitool -i hci0 cmd 0x08 0x0008 1E 02 01 1A 1A FF 4C 00 02 15 D9 B9 EC 1F 39 25 43 D0 80 A9 1E 39 D4 CE A9 5C 02 03 10 04 C8 00
reboot
sudo reboot
./bt.sh
sudo btle/hcitool -i hci0 cmd 0x08 0x0008 1E 02 01 1A 1A FF 4C 00 02 15 D9 B9 EC 1F 39 25 43 D0 80 A9 1E 39 D4 CE A9 5C 02 03 10 04 C8 00
sudo reboot
./bt.sh
./yizla 
sudo btle/hcitool -i hci0 cmd 0x08 0x0008 1E 02 01 1A 1A FF 4C 00 02 15 D9 B9 EC 1F 39 25 43 D0 80 A9 1E 39 D4 CE A9 5C 02 03 10 04 C8 00
sudo btle/hcitool -i hci0 cmd 0x08 0x0008 1E 02 01 1A 1A FF 4C 00 02 15 E2 0A 39 F4 73 F5 4B C4 A1 2F 17 D1 AD 07 A9 61 00 00 00 00 C8 00
sudo reboot
./bt.sh 
./yizla
sudo btle/hcitool -i hci0 cmd 0x08 0x0008 1E 02 01 1A 1A FF 4C 00 02 15 D9 B9 EC 1F 39 25 43 D0 80 A9 1E 39 D4 CE A9 5C 02 03 10 04 C8 00
./yizla
ls
rm -rf Desktop/
rm -rf python_games/
rm ocr_pi.png 
ls
sudo apt-get install
sudo apt-get update
sudo apt-get upgrade
ls
sudo apt-get install portaudio
sudo apt-get install portaudio19-dev
ls
./wakeup 
ls
./
./bt.sh
./wakeup 
ls
nano start
chmod +x start
./start 
crontab -e
mkdir logs
crontab -e
sudo reboot
nano logs/wakeup.log 
./start &
./start > logs/wakeup.log &
./start &
./start > logs/wakeup.log &
crontab -e
sudo reboot
