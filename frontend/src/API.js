/**
 * authCheck will call the server to check user credentials. If everything is ok, then it'll respond with OK status. If not unauthorized.
 * @param {Function} setAuth - To set state of auth
 */

export const authCheck = async (setAuth) => {
  const response = await fetch("/auth/", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });

  if (response.status === 200) {
    setAuth(true);
  } else {
    setAuth(false);
  }
};

/**
 * signUpAPI will Send the user data to Server. If If everything is ok, then it'll respond with OK status and set the jwt cookie
 * @param {Object} data An object containing user credentials
 * @param {Function} setAuth To set state of auth
 */
export const signUpAPI = async (data, setAuth) => {
  const response = await fetch("/signup", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });
  if (!response.ok) {
    const message = `An error has occured: ${response.status}`;
    alert(message);
    return;
  }
  setAuth(true);
};

/**
 * @param {Object} An object containing user credentials
 * @param {Function} setAuth To set state of auth
 */
export const loginAPI = async (data, setAuth) => {
  const response = await fetch("/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });
  console.log(response);
  if (!response.ok) {
    const message = `An error has occured: ${response.status}`;
    alert(message);
    return;
  }
  setAuth(true);
};

/**
 * @param {Function} setAuth To set state of auth
 */
export const logout = async (setAuth) => {
  console.log("Logging Out");
  const response = await fetch("/logout", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });
  if (response.status === 200) {
    setAuth(false);
  }
};

/**
 *
 * @param {Function} setUserData Function to set user data state on User Detail component
 */
export const getUserData = async (setUserData) => {
  const data = await fetch("/auth/user", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  }).then((response) => response.json());

  setUserData({ email: data.email, username: data.username });
};
