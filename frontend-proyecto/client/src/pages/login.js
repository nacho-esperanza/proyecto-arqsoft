import React from "react";

import './login.css';
  
const Login = () => {
  return (
    <div className="login">
       
           <h1>Caminos del Viento</h1> 

        
       
        <div className="container_center">
            <div  className="z-depth-3 y-depth-3 x-depth-3 grey green-text lighten-4 row" id="loginContainer">
            <br />

            <h2>Iniciar Sesion</h2>

                <div className="Username">
                    <input  className='validate' type="text" name='username' placeholder="Nombre de Usuario" id='email' required />
                  
                </div>
            <div className='row'>
                <div className="password">
                    <input className='validate' type="password" name='password' placeholder="Contraseña" id='password' required />
                   <script>
                        
                   </script>
                </div>
                <label id="forgotPasswordLabel" className="text-black">
                <a id="forgotPasswordText" href="/#">Restablecer Contraseña</a>
                </label>
            </div>
            <br />
            <div className='row'>
                <button class="botom_login" id="loginButton" type='submit' name='btn_login'>Iniciar</button>
                <p>No tienes Usuario registrate:</p>
                <button class="botom_SingUp" id="SingUPButton" type='submit' name='SingUp_login' onClick={() => window.location.href = '/signup'}>Registrarse</button>
                <p></p>
                <button class="botom_SingUp" name='SingUp_login' onClick={() => window.location.href = '/#'}>Back</button>
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