import React from "react";

import './login.css';
  
const Login = () => {
  return (
    <div className="login">
        <div className="container_center">
            <div  className="z-depth-3 y-depth-3 x-depth-3 grey green-text lighten-4 row" id="loginContainer">
            <br />
                <div className="Username">
                    <input className='validate' type="text" name='username' id='email' required />
                    <label for='email'>Username</label>
                </div>
            <div className='row'>
                <div className="password">
                    <input className='validate' type='password' name='password' id='password' required />
                    <label for='password'>Password</label>
                </div>
                <label id="forgotPasswordLabel" className="text-black">
                <a id="forgotPasswordText" href="/#">Forgot Password?</a>
                </label>
            </div>
            <br />
            <div className='row'>
                <button id="loginButton" type='submit' name='btn_login' className='col s6 btn btn-small white black-text  waves-effect z-depth-1 y-depth-1'>Login</button>
            </div>
            <br/>
            </div>
        </div>
        <br />
        <br />
    
    
        </div>
  );
};
  
export default Login;