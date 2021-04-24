import React, { useState, useEffect } from "react";
import { getUserData, logout } from "./API";

export default function UserDetail({ setAuth }) {
  const [userData, setUserData] = useState({
    username: "",
    email: "",
  });

  useEffect(() => {
    getUserData(setUserData);
  }, []);
  return (
    <div className="success">
      <div className="center">
        <h2 className="form-title">Hello User👋🏻</h2>
        <span>
          <h3>Username: </h3> {userData.email}
        </span>
        <span>
          <h3>Email: </h3> {userData.username}
        </span>
        <button className="submit-btn" onClick={() => logout(setAuth)}>
          Log Out
        </button>
      </div>
    </div>
  );
}
