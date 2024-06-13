"use client";

import { useEffect, useState } from "react";

interface Message {
  message: string;
  sentBy: string;
  sentAt: string;
}

export default function Home() {
  const [messages, setMessages] = useState<Message[]>([]);

  useEffect(() => {
    const ws = new WebSocket("http://localhost:8020");

    ws.addEventListener("open", () => {
      console.log("conectou");
    });

    ws.addEventListener("message", (event: any) => {
      const newMsg = parseMessage(event.data);
      setMessages((prevMessages) => [...prevMessages, newMsg]);
    });

    return () => {
      ws.close();
    };
  }, []);

  const parseMessage = (message: string): Message => {
    return JSON.parse(message);
  };

  return (
    <div className="mx-auto">
      <h1 className="text-lg">Gochat!</h1>
      <ul>
        {messages.map((message, index) => (
          <li key={index}>{message.message}</li>
        ))}
      </ul>
    </div>
  );
}
