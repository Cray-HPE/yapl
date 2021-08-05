const ws = new WebSocket("ws://localhost:8080/ws");
export const connect = () => {
    console.log("Attempting Connection...");

    ws.onopen = () => {
      console.log("Successfully Connected");
    };
  
    ws.onmessage = msg => {
      console.log(msg);
    };
  
    ws.onclose = event => {
      console.log("Socket Closed Connection: ", event);
    };
  
    ws.onerror = error => {
      console.log("Socket Error: ", error);
    };
}