import { useEffect } from "react";

const NaverLogin = (props: any) => {
  const { naver } = window as any;
  console.log(naver);
  const NAVER_CLIENT_ID = "dIaEluz0xoWHm2tyDFjQ"; // 발급 받은 Client ID 입력
  const NAVER_CALLBACK_URL = "http://localhost:3000/naver-login"; // 작성했던 Callback URL 입력

  const initializeNaverLogin = () => {
    const naverLogin = new naver.LoginWithNaverId({
      clientId: NAVER_CLIENT_ID,
      callbackUrl: NAVER_CALLBACK_URL,
      // 팝업창으로 로그인을 진행할 것인지?
      isPopup: true,
      // 버튼 타입 ( 색상, 타입, 크기 변경 가능 )
      loginButton: { color: "green", type: 3, height: 58 },
      callbackHandle: true,
    });
    const result = naverLogin.init();

    console.log(result);
  };

  // 화면 첫 렌더링이후 바로 실행하기 위해 useEffect 를 사용하였다.
  useEffect(() => {
    console.log("part rendered", window);
    initializeNaverLogin();
  }, []);

  return (
    <div>
      <div id="naverIdLogin" />
    </div>
  );
};

export default NaverLogin;

