const WebSocket = require('ws');
const http = require('http');

// Create HTTP server
const server = http.createServer((req, res) => {
    res.writeHead(200, { 'Content-Type': 'text/plain' });
    res.end('WebSocket test server\n');
});

// Create WebSocket server attached to the HTTP server
const wss = new WebSocket.Server({ server });

wss.on('connection', ws => {
    ws.on('message', message => {
        console.log('received:', message);
        ws.send(`echo: ${message}`);  // Echo the message back to the client
    });

    ws.send('WebSocket connection established');
});

// Start server
const port = process.env.PORT || 3000;
server.listen(port, () => {
    console.log(`Server is listening on port ${port}`);
});
