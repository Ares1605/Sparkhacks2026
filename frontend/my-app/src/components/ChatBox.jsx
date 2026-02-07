import { useState, useRef, useEffect } from "react";
import "./ChatBox.css";

export default function ChatBox() {
  const [input, setInput] = useState("");
  const [messages, setMessages] = useState([]);
  const [open, setOpen] = useState(false);

  const inputRef = useRef(null);
  const chatboxRef = useRef(null);

  const sendMessage = (e) => {
    if (e.key === "Enter" && input.trim()) {
      setMessages(prev => [
        ...prev,
        { text: input, sender: "user" }
      ]);
      setInput("");
      setOpen(true);
    }
  };

  useEffect(() => {
    const handleClickOutside = (e) => {
      if (
        chatboxRef.current &&
        !chatboxRef.current.contains(e.target)
      ) {
        setOpen(false);
      }
    };

    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  return (
    <div
      ref={chatboxRef}
      className="chatbox"
      onClick={() => inputRef.current.focus()}
    >
      <div className={`chatbox-messages ${open ? "open" : ""}`}>
        {messages.map((msg, i) => (
          <div key={i} className={`chatbox-message ${msg.sender}`}>
            {msg.text}
          </div>
        ))}
      </div>

      <div className="chatbox-input-wrapper">
        <input
          ref={inputRef}
          className="chatbox-input"
          type="text"
          placeholder="Ask me something"
          value={input}
          onChange={(e) => setInput(e.target.value)}
          onKeyDown={sendMessage}
          onFocus={() => setOpen(true)}
        />
      </div>
    </div>
  );
}