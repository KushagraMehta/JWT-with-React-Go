import { useEffect, useState } from "react";
import "./App.scss";
import Login from "./Login";
import UserDetail from "./UserDetail";
import { authCheck } from "./API";

function App() {
  const [auth, setAuth] = useState(false);

  // Checking if the user already authenticated or not. If yes, then set auth to true else provide with login form
  useEffect(() => {
    authCheck(setAuth);
  }, []);

  return <div className="form-structor">{auth ? <UserDetail setAuth={setAuth} /> : <Login setAuth={setAuth} />}</div>;
}

export default App;
