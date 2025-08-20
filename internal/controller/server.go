package controller

import (
	"SoB/internal/engine"
	"crypto/rand"
	"encoding/hex"
	"github.com/gorilla/websocket"
	"log/slog"
	"net/http"
)

type GameServer struct {
	controller Controller
	engine     *engine.Engine
	log        *slog.Logger
}

func NewGameServer(controller Controller, engine *engine.Engine, logger *slog.Logger) *GameServer {
	return &GameServer{
		controller: controller,
		engine:     engine,
		log:        logger,
	}
}

func (s *GameServer) Start(addr string) error {
	// WebSocket endpoint
	http.HandleFunc("/ws", s.handleWebSocket)

	// Test UI endpoint
	http.HandleFunc("/", s.handleTestUI)

	s.log.Info("starting game server",
		slog.String("addr", addr),
		slog.String("websocket", addr+"/ws"),
		slog.String("test_ui", addr+"/"))

	return http.ListenAndServe(addr, nil)
}

func (s *GameServer) handleTestUI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(testHTML))
}

// Embedded test HTML
const testHTML = `<!DOCTYPE html>
<html>
<head>
    <title>Game Server Test Client</title>
    <style>
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            max-width: 900px;
            margin: 0 auto;
            padding: 20px;
            background: #1a1a1a;
            color: #e0e0e0;
        }
        h1 {
            color: #4CAF50;
            border-bottom: 2px solid #4CAF50;
            padding-bottom: 10px;
        }
        .controls {
            margin: 20px 0;
            display: flex;
            gap: 10px;
            flex-wrap: wrap;
        }
        button {
            padding: 10px 20px;
            font-size: 16px;
            border: none;
            border-radius: 5px;
            cursor: pointer;
            transition: all 0.3s;
        }
        button:hover {
            transform: translateY(-2px);
            box-shadow: 0 4px 8px rgba(0,0,0,0.3);
        }
        .connect { background: #4CAF50; color: white; }
        .action { background: #2196F3; color: white; }
        .disconnect { background: #f44336; color: white; }
        button:disabled {
            opacity: 0.5;
            cursor: not-allowed;
            transform: none;
        }
        .status {
            padding: 10px;
            border-radius: 5px;
            margin: 10px 0;
            font-weight: bold;
        }
        .connected { background: #4CAF50; color: white; }
        .disconnected { background: #f44336; color: white; }
        .connecting { background: #ff9800; color: white; }
        #messages {
            background: #2a2a2a;
            border: 1px solid #444;
            border-radius: 5px;
            padding: 15px;
            height: 400px;
            overflow-y: auto;
            font-family: 'Courier New', monospace;
            font-size: 14px;
        }
        .message {
            padding: 5px 0;
            border-bottom: 1px solid #333;
        }
        .message:last-child {
            border-bottom: none;
        }
        .sent { color: #4CAF50; }
        .received { color: #2196F3; }
        .error { color: #f44336; }
        .info { color: #ff9800; }
        .input-group {
            margin: 20px 0;
            display: flex;
            gap: 10px;
        }
        input[type="text"] {
            flex: 1;
            padding: 10px;
            font-size: 14px;
            border: 1px solid #444;
            border-radius: 5px;
            background: #2a2a2a;
            color: #e0e0e0;
        }
        .json-input {
            width: 100%;
            min-height: 100px;
            font-family: 'Courier New', monospace;
            background: #2a2a2a;
            color: #e0e0e0;
            border: 1px solid #444;
            border-radius: 5px;
            padding: 10px;
            margin: 10px 0;
        }
    </style>
</head>
<body>
    <h1>🎮 Game Server Test Client</h1>
    
    <div id="status" class="status disconnected">Disconnected</div>
    
    <div class="controls">
        <button id="connectBtn" class="connect" onclick="connect()">Connect</button>
        <button id="disconnectBtn" class="disconnect" onclick="disconnect()" disabled>Disconnect</button>
        <button class="action" onclick="sendMove()" disabled>Send Move</button>
        <button class="action" onclick="sendChat()" disabled>Send Chat</button>
        <button class="action" onclick="clearMessages()">Clear Messages</button>
    </div>

    <div class="input-group">
        <input type="text" id="chatInput" placeholder="Enter chat message..." disabled>
    </div>

    <h3>Custom Message (JSON)</h3>
    <textarea id="customMessage" class="json-input" placeholder='{"action": "move", "params": {"x": 10, "y": 20}}'></textarea>
    <button class="action" onclick="sendCustom()" disabled>Send Custom Message</button>

    <h3>Messages</h3>
    <div id="messages"></div>

    <script>
        let ws;
        let messageCount = 0;
        
        function updateStatus(status) {
            const statusEl = document.getElementById('status');
            statusEl.className = 'status ' + status;
            statusEl.textContent = status.charAt(0).toUpperCase() + status.slice(1);
            
            // Update button states
            const isConnected = status === 'connected';
            document.getElementById('connectBtn').disabled = isConnected;
            document.getElementById('disconnectBtn').disabled = !isConnected;
            document.getElementById('chatInput').disabled = !isConnected;
            
            // Update all action buttons
            document.querySelectorAll('.action').forEach(btn => {
                if (btn.textContent !== 'Clear Messages') {
                    btn.disabled = !isConnected;
                }
            });
        }
        
        function connect() {
            updateStatus('connecting');
            
            const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
            const wsUrl = protocol + '//' + window.location.host + '/ws';
            
            ws = new WebSocket(wsUrl);
            
            ws.onopen = () => {
                log('Connected to server', 'info');
                updateStatus('connected');
            };
            
            ws.onmessage = (event) => {
                log('Received: ' + event.data, 'received');
                
                // Try to parse and pretty-print JSON
                try {
                    const json = JSON.parse(event.data);
                    log('Parsed: ' + JSON.stringify(json, null, 2), 'info');
                } catch (e) {
                    // Not JSON, that's okay
                }
            };
            
            ws.onerror = (error) => {
                log('WebSocket error occurred', 'error');
                updateStatus('disconnected');
            };
            
            ws.onclose = () => {
                log('Disconnected from server', 'info');
                updateStatus('disconnected');
            };
        }
        
        function sendMove() {
            const msg = {
                action: "move",
                params: {
                    x: Math.random() * 100,
                    y: Math.random() * 100,
                    z: 0
                }
            };
            sendMessage(msg);
        }
        
        function sendChat() {
            const input = document.getElementById('chatInput');
            const text = input.value.trim();
            if (!text) return;
            
            const msg = {
                action: "chat",
                params: {
                    channel: "global",
                    text: text
                }
            };
            sendMessage(msg);
            input.value = '';
        }
        
        function sendCustom() {
            const textarea = document.getElementById('customMessage');
            try {
                const msg = JSON.parse(textarea.value);
                sendMessage(msg);
                textarea.value = '';
            } catch (e) {
                log('Invalid JSON: ' + e.message, 'error');
            }
        }
        
        function sendMessage(msg) {
            if (ws && ws.readyState === WebSocket.OPEN) {
                const msgStr = JSON.stringify(msg);
                ws.send(msgStr);
                log('Sent: ' + msgStr, 'sent');
            } else {
                log('Not connected', 'error');
            }
        }
        
        function disconnect() {
            if (ws) {
                ws.close();
            }
        }
        
        function clearMessages() {
            document.getElementById('messages').innerHTML = '';
            messageCount = 0;
        }
        
        function log(msg, type = 'info') {
            messageCount++;
            const timestamp = new Date().toLocaleTimeString();
            const messagesEl = document.getElementById('messages');
            
            messagesEl.innerHTML += 
                '<div class="message ' + type + '">' +
                '[' + timestamp + '] ' + 
                msg.replace(/</g, '&lt;').replace(/>/g, '&gt;') +
                '</div>';
            
            // Auto-scroll to bottom
            messagesEl.scrollTop = messagesEl.scrollHeight;
            
            // Limit messages to prevent memory issues
            if (messageCount > 100) {
                messagesEl.removeChild(messagesEl.firstChild);
                messageCount--;
            }
        }
        
        // Allow Enter key to send chat
        document.getElementById('chatInput').addEventListener('keypress', (e) => {
            if (e.key === 'Enter') {
                sendChat();
            }
        });
        
        // Pretty-print initial custom message example
        document.getElementById('customMessage').value = JSON.stringify({
            action: "move",
            params: {x: 10, y: 20}
        }, null, 2);
    </script>
</body>
</html>`

func (s *GameServer) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// Configure origin checking
			return true
		},
	}

	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.log.Info("upgrade error", "error", err)
		return
	}

	conn := &WSConnection{
		conn: wsConn,
		id:   generateID(),
		send: make(chan []byte, 256),
		log:  s.log,
	}

	// Notify controller of new connection
	if err = s.controller.OnConnect(conn); err != nil {
		if err = conn.Close(); err != nil {
			s.log.Info("error closing connection", "error", err)
		}
		s.log.Info("error on connect", "error", err)
		return
	}

	// Start goroutines for reading and writing
	go conn.readPump(s.controller)
	go conn.writePump()
}

func generateID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}
