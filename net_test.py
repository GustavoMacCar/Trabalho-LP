import socket
import struct
import numpy as np
with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    s.connect(('127.0.0.1', 8080))
    a = np.array([10,90,5555,913098,169834]).astype(np.uint32)
    if not s.recv(1)[0]:
        print("Peer is occupied")
        exit(0)
    print(len(a))
    print(struct.pack(">I",len(a)))
    s.sendall(struct.pack(">I",len(a)))
    for i in a:
        s.sendall(struct.pack(">I",i))