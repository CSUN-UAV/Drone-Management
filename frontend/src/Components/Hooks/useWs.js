import { useEffect, useState } from 'react'

export function useWs() {
    const [ rs, setRs ] = useState(0);
    const [ ws, setWs ] = useState(null);
    const [ logHistory, setLogHistory ] = useState([]);

    const heartbeat = async (ws) => {
        setTimeout(
            function() {
                if(rs !== ws.readyState){
                    setRs(ws.readyState);
                }
                heartbeat(ws);
            }
            .bind(this),
            1000
        );
    }

    
    const configureWebsockets = async() => {
        ws.onopen = function (open_event) {
            ws.onmessage = function(event) {
                let message = JSON.parse(event.data);
                console.log(message)
                // switch(message['Type']){
                //     default:
                //         break;
                // }
            }
        }
        ws.onclose = function(close_event) {
            console.log(close_event);
        }
        ws.onerror = function(error_event) {
            console.log(error_event)
        }
    }

    useEffect(() => {
        if(ws === null) { setWs(new WebSocket("ws://csunuav.me:1200/ws")); }
        if(ws !== null && rs === 0) { configureWebsockets(); heartbeat(ws); }
    }, [ws,rs])

    return {
        rs,
        ws,
        logHistory,
        setRs,
        setWs,
        setLogHistory
    }
}