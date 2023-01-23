import React from "react";
import { useEffect } from "react";
import { useLocation, useParams } from "react-router-dom";
import { parse } from "url";

export default function LoginCallback() {
  const location = useLocation();
  const params = useParams();
  useEffect(() => {
    //console.log(window.location);
    const obj = parse(window.location.href, true);
    //console.log(obj.query);

    if (!obj.query?.code) return;
    fetch("http://127.0.0.1:8080/naver-login", {
      method: "POST",
      body: JSON.stringify(obj.query),
    }).then((r) => {
      console.log(
        r.json().then((r) => {
          console.log("result", r);
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
    </div>
  );
}
