#!/usr/bin/env python
import socket

gfile = 'data/p_site'

nfile = 'new.txt'

socket.setdefaulttimeout(5)


def try_connect(domain, port):
    for _ in range(2):
        sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        try:
            result = sock.connect_ex((domain, port))
        except socket.gaierror:
            print(f'{ domain } no dns')
            return False, 1
        if result == 0:
            return True, 0
    return False, 0

def cwho(domain):
    r443, dns = try_connect(domain, 443)
    if r443:
        return True
    if dns == 1:
        return False
    r80, dns = try_connect(domain, 80)
    if r80:
        return True
    return False

def save(domain):
    with open(nfile, 'a') as f:
        f.write(f'{ domain }\n')

def main():
    open(nfile, 'w').close()
    with open(gfile) as f:
        for l in f:
            if '#' in l: continue
            domain = l.strip()
            result = cwho(domain)
            if result:
                save(domain)
            print(domain, result)

main()


