import React, { useEffect, useState } from 'react';
import { useNavigate, } from "react-router-dom";

import './Login.css';

// Import react hot toast
import toast, { Toaster } from 'react-hot-toast';


const Login = () => {
    const navigate = useNavigate(); //permite la navegación entre paginasd con las rutas
    const [email, setEmail] = useState(''); //se inicializan las variables vacias
    const [password, setPassword] = useState('');
    const register = () =>{ // funbcion que te redirige a  register
      navigate("/signup");
    };
    const handleSubmit = async (e) => { //recibe los datos del formulario
        e.preventDefault(); // para que no recarga la página
        if (email === '') {
            document.getElementById("inputEmailLogin").style.borderColor = 'red';
        }
        if (password === '') {
            document.getElementById("inputPasswordLogin").style.borderColor = 'red';
        }
        else{
            try {   //envía la respuesta al back (postaman basicamente)
                const response = await fetch('http://localhost:8090/user/login', {
                    method: 'POST', headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({email, password}),
                }).then(response => {
                        if(response.ok){
                            toast.success("Usuario Valido")
                            return response.json();
                        }else{
                            toast.error("Usuario Invalido")
                        }
                });
                if(response.id_user) {// si el usuario existe
                    // El usuario está en la base de datos
                    console.log('Usuario válido');
                    
                    // Guardamos el token en el localStorage
                    localStorage.setItem("token", response.token);

                    // Guardamos el id del usuario en el localStorage
                    localStorage.setItem("user_id", response.id_user);

                    await fetch (`http://localhost:8090/user/${response.id_user}`, {
                        headers: {
                            Authorization: `Bearer ${localStorage.getItem('token')}`, // Agrega el token de autenticación en los headers
                        },
                    }).then(response => response.json())
                    .then(data => {
                        console.log(data)
                        localStorage.setItem("user_name", data.name);
                        localStorage.setItem("user_email", data.email);
                    });

                    

                    // Verificamos si el usuario es admin
                    if (localStorage.getItem("user_email") === "admin@gmail.com") {
                        // El usuario es admin
                        toast('Bienvenido Admin!', {
                            icon: '🙏',
                          });
                        localStorage.setItem("isAdmin", true);
                    } else {
                        // El usuario no es admin
                        localStorage.setItem("isAdmin", false);
                    }
                    console.log(response)

                    setTimeout(() => {
                        navigate('/');
                      }, 2000); // Redirige después de 2 segundos a la página de inicio
                }
            } catch (error) {
                console.log('Error al realizar la solicitud al backend:', error);
            }
        }

    };
    return (
      <div className="login">
         
             <h1>Caminos del Viento</h1> 
  
          
         
          <div className="container_center">
              <div  className="z-depth-3 y-depth-3 x-depth-3 grey green-text lighten-4 row" id="loginContainer">
              <br />
  
              <h2>Iniciar Sesion</h2>

              <form id="formLogin" onSubmit={handleSubmit} >
  
                  <div className="Username">
                    
                      <input id={"inputEmailLogin"}  
                      type="email" 
                      name='email' 
                      placeholder="Nombre de Usuario" 
                      value={email}
                        onChange={(e) => setEmail(e.target.value)}
                       required />
                    
                  </div>
              <div className='row'>
                  <div className="password">
                      <input id={"inputPasswordLogin"}
                       type="password" 
                       name='password' 
                       placeholder="Contraseña" 
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                       required />
                     <script>
                          
                     </script>
                  </div>
              </div>
              <br />
              <div className='row'>
                  <button className="botom_login" id="loginButton" type='submit' name='btn_login'>Iniciar</button>
                  <p>No tienes Usuario registrate:</p>
                  <button className="botom_SingUp" id="SingUPButton" type='submit' name='SignUp_login' onClick={() => navigate("/signup") /*window.location.href = '/signup'*/}>Registrarse</button>
                  <p></p>
                  <button className="botom_SingUp" name='SignUp_login' onClick={() => window.location.href = '/'}>Back</button>
              </div>
              </form>
              <Toaster position="top-center" reverseOrder={false} />
              <br/>
              </div>
              
          </div>
          <br />
          <br />
      
      
          </div>
    );
  };
    
  export default Login;