/**
 * authCheck will call the server to check user credentials. If everything is ok, then it'll respond with OK status. If not unauthorized.
 * @param {Function} setAuth - To set state of auth
 */

export const authCheck = async (setAuth) => {};

/**
 * signUpAPI will Send the user data to Server. If If everything is ok, then it'll respond with OK status and set the jwt cookie
 * @param {Object} data An object containing user credentials
 * @param {Function} setAuth To set state of auth
 */
export const signUpAPI = async (data, setAuth) => {};

/**
 * @param {Object} An object containing user credentials
 * @param {Function} setAuth To set state of auth
 */
export const loginAPI = async (data, setAuth) => {};

/**
 * @param {Function} setAuth To set state of auth
 */
export const logout = async (setAuth) => {};

/**
 *
 * @param {Function} setUserData Function to set user data state on User Detail component
 */
export const getUserData = async (setUserData) => {};
