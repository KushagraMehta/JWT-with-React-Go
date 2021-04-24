import { useState } from "react";
import { signUpAPI, loginAPI } from "./API";

function Login({ setAuth }) {
  const [userData, setUserData] = useState({
    username: "",
    email: "",
    password: "",
  });
  const [slideUp, setSlideUp] = useState(true);

  const updateFormData = (event) => {
    if (event.target.placeholder === "Name") {
      setUserData({ ...userData, username: event.target.value });
    } else if (event.target.placeholder === "Email") {
      setUserData({ ...userData, email: event.target.value });
    } else if (event.target.placeholder === "Password") {
      setUserData({ ...userData, password: event.target.value });
    }
  };
  return (
    <>
      <div className={`signup ${slideUp ? "slide-up" : ""}`}>
        <h2 className="form-title" id="signup" onClick={slideUp ? () => setSlideUp(!slideUp) : () => {}}>
          <span>or</span>Sign up
        </h2>
        <div className="form-holder">
          <input type="text" className="input" placeholder="Name" value={userData.username} onChange={updateFormData} />
          <input type="email" className="input" placeholder="Email" value={userData.email} onChange={updateFormData} />
          <input type="password" className="input" placeholder="Password" value={userData.password} onChange={updateFormData} />
        </div>
        <button className="submit-btn" onClick={() => signUpAPI(userData, setAuth)}>
          Sign up
        </button>
      </div>
      <div className={`login  ${slideUp ? "" : "slide-up"}`}>
        <div className="center">
          <h2 className="form-title" id="login" onClick={slideUp ? () => {} : () => setSlideUp(!slideUp)}>
            <span>or</span>Log in
          </h2>
          <div className="form-holder">
            <input type="email" className="input" value={userData.email} placeholder="Email" onChange={updateFormData} />
            <input type="password" className="input" value={userData.password} placeholder="Password" onChange={updateFormData} />
          </div>
          <button className="submit-btn" onClick={() => loginAPI(userData, setAuth)}>
            Log in
          </button>
        </div>
      </div>
    </>
  );
}

export default Login;
