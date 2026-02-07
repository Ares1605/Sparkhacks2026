import { useEffect, useRef, useState } from "react";
import "./ChatBox.css";
import { marked } from "marked";
import sendIcon from "../assets/send.svg";

export default function ChatBox() {
  const [sessionUUID, setSessionUUID] = useState();
  useEffect(() => {
    let isMounted = true;
    (async () => {
      const response = await fetch("http://localhost:8080/session/create");
      if (!response.ok) {
        throw new Error(`Response status: ${response.status}`);
      }

      const result = await response.json();
      console.log(result);
      if (isMounted) {
        setSessionUUID(result.session_uuid);
      }
    })();

    return () => {
      isMounted = false;
    };
  }, []);


  const [chatHistory, setChatHistory] = useState([{
    "role": "user",
    "message": "Helloworld",
  }, {
      "role": "server",
      "message": "hii!!!",
    }]);
  const [input, setInput] = useState("");
  const historyRef = useRef(null);
  const handleInputChange = (e) => {
    setInput(e.target.value);
  };

  useEffect(() => {
    const historyEl = historyRef.current;
    if (!historyEl) return;
    requestAnimationFrame(() => {
      historyEl.scrollTop = historyEl.scrollHeight;
    });
  }, [chatHistory]);
  async function handleSubmit(e) {
    e.preventDefault();
    if (!input.trim()) return;
    if (!sessionUUID)
      throw new Error(`Session UUID set to ${sessionUUID} while trying to send chat history message.`)

    setChatHistory(h => [
      ...h,
      { role: "user", message: input },
    ]);

    setInput("");

    const response = await fetch(`http://localhost:8080/session/ask?session=${encodeURI(sessionUUID)}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ "message": input })
    })
    if (!response.ok) {
      throw new Error(`Response status: ${response.status}`);
    }

    const result = await response.json();
    console.log(result);
    const serverResponse = marked.parse(result.response);
    setChatHistory(h => [
      ...h,
      { role: "server", message: serverResponse },
    ]);

  }

  return (
    <div className="chatbox">
      <div className="history" ref={historyRef}>
        {chatHistory.map((message, i) => (
          <div key={i} className={`message-wrapper ${message.role}`}>
            <div
              className={`message`}
            dangerouslySetInnerHTML={{ __html: message.message }}></div>
          </div>
        ))}
      </div>
      <form className="input-wrapper" onSubmit={handleSubmit}>
        <input
          value={input}
          onChange={handleInputChange}
        />
        <div className="options">
          <button className="send"><img src={sendIcon} alt="Send Message Icon" /></button>
        </div>
      </form>

    </div>
  );
}
