import React, { useEffect, useState } from 'react';
import { useNavigate, } from "react-router-dom";
import './SignUp.css';

//parte funcional
const SignUp = () =>{

    const navigate = useNavigate(); //permite la navegación entre paginasd con las rutas
    const [name, setName] = useState(''); //se inicializan las variables vacias
    const [lastname, setLastname] = useState('');
    const [email, setEmail] = useState(''); 
    const [password, setPassword] = useState('');
    const signup = () =>{ // funbcion que te redirige a  home
      navigate("/home");
    };
    const singup1 = () =>{ // funbcion que te redirige a  login
        navigate("/login");
    };

    const handleSubmit = async (e) => { //recibe los datos del formulario
        e.preventDefault(); // para que no recarga la página
        if (nombre === '') {
            document.getElementById("inputNombreRegister").style.borderColor = 'red';
        }
        if (apellido === '') {
            document.getElementById("inputApellidoRegister").style.borderColor = 'red';
        }
        if (email === '') {
            document.getElementById("inputEmailRegister").style.borderColor = 'red';
        }
        if (password === '') {
            document.getElementById("inputPasswordRegister").style.borderColor = 'red';
        }
        else{
            try {   //envía la respuesta al back (postaman basicamente)
                const response = await fetch('http://localhost:8090/user', {
                    method: 'POST', headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({name, lastname, email, password}),
                });
                if (response.ok) {// si el usuario existe -> usuario creado en la db
                    // El usuario está en la base de datos
                    console.log('Usuario creado');
                    navigate("/home")

                } else {
                    // El usuario no (ya) está en la base de datos o hay un error en el servidor 
                    console.log('Usuario inválido');
                }
            } catch (error) {
                console.log('Error al realizar la solicitud al backend:', error);
            }
        }

        return (

    


            <div className="signup-container-1">
            
            <h1>Caminos del Viento</h1>
            
            <div className="signup-container">
            
            <h2>
                    Registro de Usuario
            </h2>
                 
            
            <form id="formRegister" onSubmit={handleSubmit}>
            
            <input id={"inputNombreRegister"} 
            type="text" 
            placeholder="Nombre"
            value={name}
            onChange={(e) => setName(e.target.value)}
            />
            
            <input id={"inputApellidoRegister"}
             type="text" 
             placeholder="Apellido"
                value={lastname}
                onChange={(e) => setLastname(e.target.value)}
             />
            
            <input id={"inputEmailRegister"}
             type="email" 
             placeholder="Email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
             />
            
            <input id={"inputPasswordRegister"}
             type="password" 
             placeholder="Contraseña"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
             />
            
            <input type="password" placeholder="Repita Contraseña"/>
            
            <button id="botonRegistrarse" type="submit">Registrate</button>
            
            </form>
            
            <p>Ya tengo Usuario</p>
            
            <button type="submit" onClick={() => window.location.href = '/login'}>Iniciar Sesion</button>
            <button onClick={() => window.location.href = '/#'}>Back</button>
            
                </div>
            </div>
              );

    };

};
    

    
export default SignUp;