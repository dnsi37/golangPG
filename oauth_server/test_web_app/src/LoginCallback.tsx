import React, { useState } from "react";
import { useEffect } from "react";
import { useLocation, useParams } from "react-router-dom";
import { parse } from "url";

export default function LoginCallback() {
  const location = useLocation();
  const params = useParams();
  const [att, setRes] = useState("");

  const signup = async () => {
    const dto = {
      access_token: att,
      phone_number: "010-5502-7723",
      id: "testid",
      password: "testpassword$1234",
    };

    const result = await fetch("http://127.0.0.1:8080/auth/signup/naver", {
      method: "POST",
      body: JSON.stringify(dto),
      headers: {
        "Content-Type": "application/json",
        // 'Content-Type': 'application/x-www-form-urlencoded',
      },
    }).catch((e) => {
      console.log(e);
    });

    console.log(result);
  };

  useEffect(() => {
    console.log(window.location);
    const sanitizedUrl = window.location.href.replace("#", "?");
    const obj = parse(sanitizedUrl, true);
    console.log(obj.query);

    if (!obj.query?.access_token) return;
    const at = obj.query?.access_token as string;
    const dto = {
      access_token: at,
    };
    setRes(at);

    console.log(dto);
    fetch("http://127.0.0.1:8080/auth/social-login/naver", {
      method: "POST",
      body: JSON.stringify(dto),
      headers: {
        "Content-Type": "application/json",
        // 'Content-Type': 'application/x-www-form-urlencoded',
      },
    }).then((r) => {
      console.log(r.status);
      console.log(
        r.json().then((r) => {
          console.log("result", r);
          //setRes(r);
        })
      );
    });
    // 쿼리 valdiation
    // send to server by using POST request
    // receive response (Token Value) and  Save it to DB
    // redirect to Home
  }, []);
  return (
    <div>
      <div>Hello</div>
      <button
        onClick={() => {
          console.log("hg");
          console.log(att);
        }}
      >
        버튼
      </button>
      <button
        onClick={() => {
          signup();
        }}
      >
        회원가입
      </button>
    </div>
  );
}
