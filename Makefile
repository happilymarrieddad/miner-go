# Going to assume default distro is Ubuntu

install:
	go get -u github.com/go-vgo/robotgo \
		github.com/onsi/ginkgo/ginkgo \
		github.com/onsi/gomega/... 

install_ubuntu:
	sudo apt-get install gcc libc6-dev -y
	sudo apt-get install libx11-dev xorg-dev libxtst-dev libpng++-dev -y
	sudo apt-get install xcb libxcb-xkb-dev x11-xkb-utils libx11-xcb-dev libxkbcommon-x11-dev -y
	sudo apt-get install libxkbcommon-dev -y
	sudo apt-get install xsel xclip -y

install_fedora:
	sudo dnf install libxkbcommon-devel libXtst-devel libxkbcommon-x11-devel xorg-x11-xkb-utils-devel
	sudo dnf install libpng-devel
	sudo dnf install xsel xclip



