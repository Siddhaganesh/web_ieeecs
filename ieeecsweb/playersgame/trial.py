import socket
import threading

class ConnectFourServer:
    def __init__(self, host = 'localhost', port = 5000):
        self.host = host
        self.port = port
        self.server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.server.bind((self.host, self.port))
        self.server.listen(2)
        self.clients = []

    def broadcast(self, message, client):
        for c in self.clients:
            if c != client:
                c.send(message)

    def handle_client(self, client):
        while True:
            try:
                message = client.recv(1024)
                self.broadcast(message, client)
            except:
                index = self.clients.index(client)
                self.clients.remove(client)
                client.close()
                nickname = self.nicknames[index]
                self.nicknames.remove(nickname)
                self.broadcast(f'{nickname} left the game!'.encode('ascii'), client)
                break

    def receive(self):
        while True:
            client, address = self.server.accept()
            print(f'Connected with {str(address)}')

            client.send('NICK'.encode('ascii'))
            nickname = client.recv(1024).decode('ascii')
            self.nicknames.append(nickname)
            self.clients.append(client)

            print(f'Nickname of the client is {nickname}!')
            self.broadcast(f'{nickname} joined the game!'.encode('ascii'), client)
            client.send('Connected to the server!'.encode('ascii'))

            thread = threading.Thread(target=self.handle_client, args=(client,))
            thread.start()

    def start(self):
        print('Server Started!')
        self.receive()

server = ConnectFourServer()
server.start()
