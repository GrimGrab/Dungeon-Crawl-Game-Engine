package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow connections from any origin
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}
	defer conn.Close()

	log.Printf("Client connected: %s", conn.RemoteAddr())

	// Send "hello world" message immediately upon connection
	err = conn.WriteMessage(websocket.TextMessage, []byte("hello world"))
	if err != nil {
		log.Printf("Failed to send hello message: %v", err)
		return
	}

	// Handle incoming messages and echo "hello world"
	for {
		// Read message from client
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Connection closed: %v", err)
			break
		}

		log.Printf("Received message: %s", message)

		// Always respond with "hello world"
		err = conn.WriteMessage(messageType, []byte("hello world"))
		if err != nil {
			log.Printf("Failed to send message: %v", err)
			break
		}
	}
}

func StartServer(port string) error {
	http.HandleFunc("/ws", handleWebSocket)

	// Serve a simple HTML page for testing
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
<!DOCTYPE html>
<html>
<head>
    <title>WebSocket Test</title>
</head>
<body>
    <h1>WebSocket Test</h1>
    <div id="messages"></div>
    <input type="text" id="messageInput" placeholder="Type a message">
    <button onclick="sendMessage()">Send</button>
    <button onclick="connect()">Connect</button>
    <button onclick="disconnect()">Disconnect</button>

    <script>
        let ws;
        const messages = document.getElementById('messages');

        function connect() {
            ws = new WebSocket('ws://localhost:` + port + `/ws');
            
            ws.onopen = function(event) {
                addMessage('Connected to server');
            };

            ws.onmessage = function(event) {
                addMessage('Server: ' + event.data);
            };

            ws.onclose = function(event) {
                addMessage('Disconnected from server');
            };

            ws.onerror = function(error) {
                addMessage('Error: ' + error);
            };
        }

        function disconnect() {
            if (ws) {
                ws.close();
            }
        }

        function sendMessage() {
            const input = document.getElementById('messageInput');
            if (ws && input.value) {
                ws.send(input.value);
                addMessage('You: ' + input.value);
                input.value = '';
            }
        }

        function addMessage(message) {
            const div = document.createElement('div');
            div.textContent = new Date().toLocaleTimeString() + ' - ' + message;
            messages.appendChild(div);
            messages.scrollTop = messages.scrollHeight;
        }

        document.getElementById('messageInput').addEventListener('keypress', function(e) {
            if (e.key === 'Enter') {
                sendMessage();
            }
        });
    </script>
</body>
</html>
		`))
	})

	log.Printf("WebSocket server starting on port %s", port)
	log.Printf("Visit http://localhost:%s to test the WebSocket connection", port)

	return http.ListenAndServe(":"+port, nil)
}
