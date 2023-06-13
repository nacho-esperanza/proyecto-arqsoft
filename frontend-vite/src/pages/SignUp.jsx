import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './SignUp.css';

// importar hot toaster
import toast, { Toaster } from 'react-hot-toast';


const SignUp = () => {
  const navigate = useNavigate();
  const [name, setName] = useState('');
  const [lastname, setLastname] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const signup = () => {
    navigate('/');
  };

  const singup1 = () => {
    navigate('/login');
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (name === '') {
      document.getElementById('inputNombreRegister').style.borderColor = 'red';
    }
    if (lastname === '') {
      document.getElementById('inputApellidoRegister').style.borderColor = 'red';
    }
    if (email === '') {
      document.getElementById('inputEmailRegister').style.borderColor = 'red';
    }
    if (password === '') {
      document.getElementById('inputPasswordRegister').style.borderColor = 'red';
    } else {
      try {
        const response = await fetch('http://localhost:8090/user', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ name, lastname, email, password }),
        });

        if (response.ok) {
          console.log('Usuario creado');
          toast.success('Usuario creado');
          setTimeout(() => {
            navigate('/');
          }, 2000); // Redirige después de 2 segundos a la página de inicio
        } else {
          toast.error('Usuario inválido');
          console.log('Usuario inválido');
        }

      
      } catch (error) {
        console.log('Error al realizar la solicitud al backend:', error);
      }
    }
  };

  return (
    <div className="signup-container-1">
      <h1>Caminos del Viento</h1>
      <div className="signup-container">
        <h2>Registro de Usuario</h2>
        <form id="formRegister" onSubmit={handleSubmit}>
          <input
            id="inputNombreRegister"
            type="text"
            placeholder="Nombre"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />

          <input
            id="inputApellidoRegister"
            type="text"
            placeholder="Apellido"
            value={lastname}
            onChange={(e) => setLastname(e.target.value)}
          />

          <input
            id="inputEmailRegister"
            type="email"
            placeholder="Email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />

          <input
            id="inputPasswordRegister"
            type="password"
            placeholder="Contraseña"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />

          <input type="password" placeholder="Repita Contraseña" />

          <button id="botonRegistrarse" type="submit">
            Registrate
          </button>
        </form>
        <Toaster position="top-center" reverseOrder={false} />
        

        <p>Ya tengo Usuario</p>
        <button type="submit" onClick={() => window.location.href = '/login'}>
          Iniciar Sesión
        </button>
        <button onClick={() => window.location.href = '/'}>Back</button>
      </div>
    </div>
  );
};

export default SignUp;