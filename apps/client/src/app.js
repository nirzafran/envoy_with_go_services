// Import the client and the message definition
const {
	TimeServiceClient,
} = require('./generated/grpc/jsclient/time_service_grpc_web_pb')
const {
	GetCurrentTimeRequest,
} = require('./generated/grpc/jsclient/time_service_pb')

// Connect to the gprc-web server
const client = new TimeServiceClient('http://localhost:8080', null, null)
// This is a neat chrome extension that allows you to spy on grpc-web traffic just like you would on normal traffic.
// You can find it here: https://chrome.google.com/webstore/detail/grpc-web-developer-tools/ddamlpimmiapbcopeoifjfmoabdbfbjj?hl=en
const enableDevTools = window.__GRPCWEB_DEVTOOLS__ || (() => {})
enableDevTools([client])

// Send getCurrentTime request
client.getCurrentTime(new GetCurrentTimeRequest(), {}, (err, response) => {
	// handle the response
	this.lastTimeResponse = response.getCurrentTime()
})
