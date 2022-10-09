#SHELL = /bin/sh

update:
#	go run v2ray_domain_list.go -datapath ./data
	go run make_pac_text.go > pac.txt
	go run make_clash_text.go > good_clash.txt

