import React from "react";

import "./App.css";
import NaverLogin from "./NaverLogin";

const NAVER_CLIENT_ID = "dIaEluz0xoWHm2tyDFjQ"; // 발급 받은 Client ID 입력
const NAVER_CALLBACK_URL = "http://localhost:3000/naver-login"; // 작성했던 Callback URL 입력

let naver_api_url =
  "https://nid.naver.com/oauth2.0/authorize?response_type=code&client_id=" +
  NAVER_CLIENT_ID +
  "&redirect_uri=" +
  encodeURI(NAVER_CALLBACK_URL) +
  "&state=" +
  Math.random().toString(36).substr(3, 14);

function App() {
  return (
    <div className="App">
      <head></head>
      <NaverLogin />
      <a href={naver_api_url} onLoad={(e) => {}}>
        <img
          height="50"
          src="http://static.nid.naver.com/oauth/small_g_in.PNG"
        />
      </a>
      <button
        onClick={() => {
          console.log("hg");
        }}
      >
        버튼
      </button>
    </div>
  );
}

export default App;
