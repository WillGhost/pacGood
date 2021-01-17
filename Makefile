#SHELL = /bin/sh

update:
	go run make_pac_text.go > pac.txt
	go run v2ray_domain_list.go -datapath ./data


