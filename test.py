# eth3r-networks 2023-2024

import socket
from time import sleep

sc = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
sc.connect(("localhost", 2142))

# sc.send("\x05\x31\xb0\x0b\x00\x01".encode())

my_bytes = bytearray()
my_bytes.append(0x05)
my_bytes.append(0x31)
my_bytes.append(0xb0)
my_bytes.append(0x0b)
my_bytes.append(0x00)
my_bytes.append(0x01)

sc.send(my_bytes)

print("Sent first packet")

sleep(1)

# sc.send("\x0e\x1f\x00\x02\xaa\xaa".encode())

my_bytes = bytearray()
my_bytes.append(0x0e)
my_bytes.append(0x1f)
my_bytes.append(0x00)
my_bytes.append(0x02)
my_bytes.append(0xaa)
my_bytes.append(0xaa)

sc.send(my_bytes)

sleep(100)

sc.close()
